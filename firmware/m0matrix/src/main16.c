
/* Joining the FPGA cult next project, I promise */

#include <stdlib.h>
#include "stm32f0xx.h"

#include "bitbang16.h"
#include "util.h"

#define CHANNEL_OFFSET (4)
/* with interleaving for speedy computation of bitbang buffers */
#define I(n) ( (n&~3) + n )
static const uint8_t segment_pin_interleaved[] =
/*    A     B     C     D     E     F    G1    G2     H     J     K     L     M     N    Dp    nc -> outX */
 { I(12),I(13), I(1), I(6), I(0),I(15),I(11), I(7),I(14),I(10), I(9), I(4), I(3), I(2), I(5), I(8) };
#undef I

/*
 *         ._  ___________  _.
 *         | \ \____A____/ / |
 *         | |._    _    _.| |
 *         | || \  | |  / || |
 *         |F| \H\ |J| /K/ |B|
 *         | |  \_\|_|/_/  | |
 *         |_| ____   ____ |_|
 *          _ <_G1_> <_G2_> _
 *         | |   ._._._.   | |
 *         | |  / /| |\ \  | |
 *         |E| /N/ |M| \L\ |C|
 *         | ||_/  |_|  \_|| |
 *         | | ___________ | |   __
 *         |/ /_____D_____\ \|  (Dp)
 *
 *
 * Frame sent over serial:
 *
 *                  [  0  1 ] } 2x 16x little-endian uint16 in range: [0x0000..0xff00] (A B C D E F G1 G2 H J K L M N Dp nc)
 *                  [  2  3 ]   same...
 *                  [  4  5 ]   ...
 *                  [  6  7 ]
 *                  [  8  9 ]
 *                  [  A  B ]
 *                  [  C  D ]
 *                  [  E  F ]
 *                  [ 10 11 ]
 *                  [ 12 13 ]
 *                  [ 14 15 ]
 *                  [ 16 17 ]
 *                  [ 18 19 ]
 *                  [ 1A 1B ]
 *                  [ 1C 1D ]
 *                  [ 1E 1F ]
 *                            + END OF FRAME MARKER "\xff\xff\xff\xf0"
 *
 *                Frame internal mapping:
 * 
 *                 /[ R0 R1 ] \_N[0..15] segment bits re-ordered to match driver pins
 *                | [ R2 R3 ] /
 *                | [ R0 R1 ] \_N[16..31]
 *  north channel-| [ R2 R3 ] /
 *                | [ R0 R1 ] \_N[32..47]
 *                | [ R2 R3 ] /
 *                | [ R0 R1 ] \_N[48..63]
 *                 \[ R2 R3 ] /
 *                 /[ R0 R1 ] \_S[48..63]
 *                | [ R2 R3 ] /
 *                | [ R0 R1 ] \_S[32..47]
 *  south channel-| [ R2 R3 ] /
 *                | [ R0 R1 ] \_S[16..31]
 *                | [ R2 R3 ] /
 *                | [ R0 R1 ] \_S[0..15]
 *                 \[ R2 R3 ] /
 *
 * R0-R3: 4 rows, (selected using 2 I/O pins, similar to rows in a normal matrix display)
 *
 * 1 byte/segment, the 2 channels are interleaved every 4 bytes
 * FRAME: [ ROW_0 ROW_1 ROW_2 ROW_3 ]
 * ROW_X: [ N0 N1 N2 N3 S0 S1 S2 S3 N4 N5 N6 N7 S4 ... S59 N60 N61 N62 N63 S60 S61 S62 S63 ]
 */

static uint8_t buf15[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));
static uint8_t buf14[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));
static uint8_t buf13[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));
static uint8_t buf12[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));
static uint8_t bufx[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));
static uint8_t bufy[N_BITS_PER_CHANNEL] __attribute__((aligned(4)));

typedef struct
{
	uint8_t high_bits[N_LEDS];
	uint8_t low_bits[N_LEDS];
} frame_t;

static frame_t frame_a __attribute__((aligned(4)));
static frame_t frame_b __attribute__((aligned(4)));

static uint8_t zeroes[N_BITS_PER_ROW] __attribute__((aligned(4)));

static frame_t * volatile cur_frame;
static frame_t * volatile next_frame;
static frame_t * volatile draw_frame;

static int cur_row;
static int cur_pos;
static uint16_t iter;
static int is_2nd_half = 0;

#define RECV_BUF_SZ (128)
static volatile uint8_t recv_buf[RECV_BUF_SZ];

enum
{
	B15, B14, B13, B12, BX, BY,
	P15, P14, P13, P12, P11, P10, P9, P8,
	P7, P6, P5, P4, P3, P2, P1, P0,
	ZX, /* send zeroes */
//	OFF,
	E0, E1, E2, E3, E4, E5, E6, E7, E8,
	ZZZ,
};


static const uint8_t dtable[288] = /* SysTick dispatch table for BCM bitbang & precomputation */
{
/*	 15    .    .    .    .    .    .    .   14    .    .    .   13    .   12  11/10 dith */

	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P8, P15, B13, ZZZ, B12,  BX,  BY, E8,
	B15, ZZZ, P14, P13, P12, ZZZ, P10, ZZZ, B14, ZZZ,  P4, P15, B13, ZZZ, B12,  BX,  BY, E4,
	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P6, P15, B13, ZZZ, B12,  BX,  BY, E6,
	B15, ZZZ, P14, P13, P12, ZZZ,  P9, ZZZ, B14, ZZZ,  P2, P15, B13, ZZZ, B12,  BX,  BY, E2,

	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P7, P15, B13, ZZZ, B12,  BX,  BY, E7,
	B15, ZZZ, P14, P13, P12, ZZZ, P10, ZZZ, B14, ZZZ,  P3, P15, B13, ZZZ, B12,  BX,  BY, E3,
	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P5, P15, B13, ZZZ, B12,  BX,  BY, E5,
	B15, ZZZ, P14, P13, P12, ZZZ,  ZX, ZZZ, B14, ZZZ,  P1, P15, B13, ZZZ, B12,  BX,  BY, E2,


	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P8, P15, B13, ZZZ, B12,  BX,  BY, E8,
	B15, ZZZ, P14, P13, P12, ZZZ, P10, ZZZ, B14, ZZZ,  P4, P15, B13, ZZZ, B12,  BX,  BY, E4,
	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P6, P15, B13, ZZZ, B12,  BX,  BY, E6,
	B15, ZZZ, P14, P13, P12, ZZZ,  P9, ZZZ, B14, ZZZ,  P2, P15, B13, ZZZ, B12,  BX,  BY, E2,

	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P7, P15, B13, ZZZ, B12,  BX,  BY, E7,
	B15, ZZZ, P14, P13, P12, ZZZ, P10, ZZZ, B14, ZZZ,  P3, P15, B13, ZZZ, B12,  BX,  BY, E3,
	B15, ZZZ, P14, P13, P12, ZZZ, P11, ZZZ, B14, ZZZ,  P5, P15, B13, ZZZ, B12,  BX,  BY, E5,
	B15, ZZZ, P14, P13, P12, ZZZ,  ZX, ZZZ, B14, ZZZ,  P0, P15, B13, ZZZ, B12,  BX,  BY, E1,

};

const uint32_t select_row[] =
{
	CLEAR ( BIT_SELECT_A ) | CLEAR ( BIT_SELECT_B ),
	SET   ( BIT_SELECT_A ) |   SET ( BIT_SELECT_B ),
	CLEAR ( BIT_SELECT_A ) |   SET ( BIT_SELECT_B ),
	SET   ( BIT_SELECT_A ) | CLEAR ( BIT_SELECT_B ),
};

//static void off(void)              { GPIOA->BSRR = CLEAR(BIT_ENABLE_HIGH) | SET(BIT_NOT_OUTPUT_ENABLE); }

static void bitbang_15(void)       { bitbang64_clk_stm32(buf15, (void *)GPIOA);
                                     GPIOA->BSRR = select_row[cur_row];
	                                 GPIOA->BSRR = SET(BIT_ENABLE_HIGH); }
static void bitbang_14(void)       { bitbang64_clk_stm32(buf14, (void *)GPIOA); }
static void bitbang_13(void)       { bitbang64_clk_stm32(buf13, (void *)GPIOA); }
static void bitbang_12(void)       { bitbang64_clk_stm32(buf12, (void *)GPIOA); }
static void bitbang_x(void)        { bitbang64_clk_stm32(bufx,  (void *)GPIOA); }
static void bitbang_y_no_enable(void) { bitbang64_clk_no_enable_stm32(bufy,  (void *)GPIOA); }

static void prepare_15(void) { draw_frame = cur_frame;
                               cur_pos += N_BITS_PER_ROW; if (cur_pos>=N_LEDS) cur_pos = 0;
                               precomp64(buf15, &draw_frame->high_bits[cur_pos], 7, X4(BIT_NOT_OUTPUT_ENABLE));}
static void prepare_14(void) { precomp64(buf14, &draw_frame->high_bits[cur_pos], 6, X4(BIT_ENABLE_HIGH)); }
static void prepare_13(void) { precomp64(buf13, &draw_frame->high_bits[cur_pos], 5, X4(BIT_ENABLE_HIGH)); }
static void prepare_12(void) { precomp64(buf12, &draw_frame->high_bits[cur_pos], 4, X4(BIT_ENABLE_HIGH)); }
static void prepare_11(void) { precomp64(bufx,  &draw_frame->high_bits[cur_pos], 3, X4(BIT_ENABLE_HIGH)); }
static void prepare_10(void) { precomp64(bufx,  &draw_frame->high_bits[cur_pos], 2, X4(BIT_ENABLE_HIGH)); }
static void prepare_9(void)  { precomp64(bufx,  &draw_frame->high_bits[cur_pos], 1, X4(BIT_ENABLE_HIGH)); }
static void prepare_8(void)  { precomp64(bufy,  &draw_frame->high_bits[cur_pos], 0, X4(BIT_ENABLE_HIGH)); }
static void prepare_7(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 7, X4(BIT_ENABLE_HIGH)); }
static void prepare_6(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 6, X4(BIT_ENABLE_HIGH)); }
static void prepare_5(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 5, X4(BIT_ENABLE_HIGH)); }
static void prepare_4(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 4, X4(BIT_ENABLE_HIGH)); }
static void prepare_3(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 3, X4(BIT_ENABLE_HIGH)); }
static void prepare_2(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 2, X4(BIT_ENABLE_HIGH)); }
static void prepare_1(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 1, X4(BIT_ENABLE_HIGH)); }
static void prepare_0(void)  { precomp64(bufy,  &draw_frame->low_bits[cur_pos], 0, X4(BIT_ENABLE_HIGH)); }

#define FLIP_OFF (SET(BIT_NOT_OUTPUT_ENABLE)|CLEAR(BIT_ENABLE_HIGH))
#define FLIP_ON  (SET(BIT_ENABLE_HIGH)|CLEAR(BIT_NOT_OUTPUT_ENABLE))
#define SYSTICK_PERIOD ((uint32_t)(8*( F_SYS_TICK_CLK/28800/4 )))

static void enable_8(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>1); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_7(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>2); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_6(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>3); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_5(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>4); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_4(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>5); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_3(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>6); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_2(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>7); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_1(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>8); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }
static void enable_0(void) { write_wait_write(&GPIOA->BSRR, FLIP_ON, FLIP_OFF, SYSTICK_PERIOD>>9); cur_row = (cur_row+1)&(N_ROWS-1); if (cur_row) iter -= 18; }

static void zero_x(void) { precomp64(bufx, zeroes, 0, X4(BIT_ENABLE_HIGH)); }

static void ret(void) { }

typedef void (*func_t)(void);

const func_t dispatch[] =
{
	[B15] = bitbang_15,
	[B14] = bitbang_14,
	[B13] = bitbang_13,
	[B12] = bitbang_12,
	[BX] = bitbang_x,
	[BY] = bitbang_y_no_enable,
	[P15] = prepare_15,
	[P14] = prepare_14,
	[P13] = prepare_13,
	[P12] = prepare_12,
	[P11] = prepare_11,
	[P10] = prepare_10,
	[P9] = prepare_9,
	[P8] = prepare_8,
	[P7] = prepare_7,
	[P6] = prepare_6,
	[P5] = prepare_5,
	[P4] = prepare_4,
	[P3] = prepare_3,
	[P2] = prepare_2,
	[P1] = prepare_1,
	[P0] = prepare_0,

	[E8] = enable_8,
	[E7] = enable_7,
	[E6] = enable_6,
	[E5] = enable_5,
	[E4] = enable_4,
	[E3] = enable_3,
	[E2] = enable_2,
	[E1] = enable_1,
	[E0] = enable_0,

	[ZX] = zero_x,
//	[OFF] = off,
	[ZZZ]= ret,
};

void SysTick_Handler(void)
{
	dispatch[dtable[iter]]();

	if (iter < 287)
		iter = iter+1;
	else
		iter = 0;
}

static void init(void)
{
    RCC->AHBENR |= RCC_AHBENR_GPIOAEN | RCC_AHBENR_GPIOBEN; /* enable clock on GPIO A & B */
	GPIOA->ODR = BIT_NOT_OUTPUT_ENABLE;
    GPIOA->MODER = SWD|O(0)|O(1)|O(2)|O(4)|O(5)|O(6)|O(7)|O(9)|O(10);

	GPIOA->OSPEEDR = OSPEED_SWD                  |
	                 OSPEED_HIGH(PIN_LATCH)      |
	                 OSPEED_HIGH(PIN_DATA_SOUTH) |
	                 OSPEED_HIGH(PIN_DATA_NORTH) |
	                 OSPEED_HIGH(PIN_CLK_SOUTH)  |
	                 OSPEED_HIGH(PIN_CLK_NORTH)  ;

    GPIOB->ODR = 0*BIT_ENABLE_LOW;
    GPIOB->MODER = O(1);

	/* There aren't enough lines from the teensy to drive 30 boards in
     * parallel, so the 16 uarts each send two frames. swdio pulled to ground
	 * on the backplane tells the board it should display the first frame.
	 * (A floating swdio is internally pulled up by default)
	 *
	 * NOTE: this makes this code work differently with a debugger attached on startup
	 *
	 * - attached: swdio pulled to ground on startup, display first frames
	 * - detached: swdio pulled up internally on startup, display second frames
	 *
	 */
	is_2nd_half = !!( GPIOA->IDR & 1<<13 );

	clock48mhz();
	usart1_rx_pa3_dma3_enable(recv_buf, RECV_BUF_SZ, 48e6/15e5);

	cur_frame = &frame_a;
	next_frame = &frame_b;
	cur_row = 0;
	iter = 0;

	int i;
	for (i=0; i<N_LEDS; i++)
	{
		frame_a. low_bits[i] =
		frame_b. low_bits[i] = 0;
		frame_a.high_bits[i] =
		frame_b.high_bits[i] = 0;
	}

	for (i=0; i<N_BITS_PER_ROW; i++)
		zeroes[i] = 0;

	cur_pos = -N_BITS_PER_ROW; /* pre-compensate for the first invocation of prepare_10() */
	prepare_15();

    enable_sys_tick(F_SYS_TICK_CLK/28800/4);
}

static inline int dma_getchar(void)
{
	static uint32_t last = 0;
	if (last == 0)
		last = RECV_BUF_SZ;
	while (last == DMA1_Channel3->CNDTR);
	last--;
	return recv_buf[RECV_BUF_SZ-1-last];
}

static inline int dma_get_u16(void)
{
	int c;
	c  =   dma_getchar();
	c |= ( dma_getchar() << 8 );
	return c;
}

#define GOOD          0
#define GOOD_00       1
#define GOOD_01_FE    2
#define GOOD_FF       3
#define GOOD_FFFF     4
#define GOOD_FFFFFF   5
#define BAD           6
#define BAD_FF        7
#define BAD_FFFF      8
#define BAD_FFFFFF    9

#define STATE_COUNT  10

#define GOOD_RETURN  11
#define BAD_RETURN   12

#define IN_00     0
#define IN_01_EF  1
#define IN_F0     2
#define IN_F1_FE  3
#define IN_FF     4

const uint8_t fsm[STATE_COUNT][8] =
{
	[GOOD]        = { [IN_00]=GOOD_00, [IN_01_EF]=GOOD_01_FE, [IN_F0]=GOOD_01_FE , [IN_F1_FE]=GOOD_01_FE, [IN_FF]=GOOD_FF     },
	[GOOD_00]     = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=GOOD        },
	[GOOD_01_FE]  = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=BAD_FF      },
	[GOOD_FF]     = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=GOOD_FFFF   },
	[GOOD_FFFF]   = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=GOOD_FFFFFF },
	[GOOD_FFFFFF] = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=GOOD_RETURN, [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },
	[BAD]         = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FF      },
	[BAD_FF]      = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFF    },
	[BAD_FFFF]    = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },
	[BAD_FFFFFF]  = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD_RETURN , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },

};


static int read_next_frame(void)
{
	int i, c=0, v, digit_location, chip_delta, seg;

	if (is_2nd_half)
		for (i=0; i<N_LEDS; i++)
		{
			c = dma_get_u16();
			if (c > 0xff00)
				break;
		}

	/* do the mapping to an easy-to-use internal representation once */
#define NEXT_CHIP (N_PINS_PER_CHIP*N_CHANNELS)
#define NEXT_ROW (N_BITS_PER_ROW)
#define SEGMENT_MASK (N_PINS_PER_CHIP-1)
#define CHIP_MASK (N_PINS_PER_CHIP*N_ROWS-1)

	chip_delta = NEXT_CHIP - N_ROWS*NEXT_ROW; /* pre-undo first iteration */
	digit_location = - chip_delta - NEXT_ROW;

	if (c <= 0xff00)
	for (i=0; i<N_LEDS; i++)
	{
		seg = SEGMENT_MASK & i;
		if (seg == 0)
		{
			digit_location += NEXT_ROW;
			if ( (i & CHIP_MASK) == 0 )
			{
				digit_location += chip_delta;
				if (i == N_LEDS/N_CHANNELS)
				{
					digit_location = (N_CHIPS_PER_CHANNEL-1)*NEXT_CHIP+CHANNEL_OFFSET;
					chip_delta = - NEXT_CHIP - N_ROWS*NEXT_ROW;
				}
			}
		}
		c = dma_get_u16();
		if (c > 0xff00)
			break;

	//	v = c + (c>>8);
		v = c;
		next_frame->high_bits[digit_location+segment_pin_interleaved[seg]] = v>>8;
		next_frame->low_bits [digit_location+segment_pin_interleaved[seg]] = v;
	}

	int s=GOOD;

	if (c == 0xffff)
		s = BAD_FFFF;
	else if (c > 0xff00)
		s = BAD_FF;

	for(;;)
	{
		c = dma_getchar();

		if (c == 0)
			i = IN_00;
		else if (c < 0xf0)
			i = IN_01_EF;
		else if (c == 0xf0)
			i = IN_F0;
		else if (c == 0xff)
			i = IN_FF;
		else
			i = IN_F1_FE;

		s = fsm[s][i];

		if (s == GOOD_RETURN)
			return 1;
		else if (s == BAD_RETURN)
			return 0;
	}
}

int main(void)
{
	init();

	for(;;)
	{
		while (! read_next_frame() );
		frame_t *tmp = cur_frame;
		cur_frame = next_frame;
		next_frame = tmp;
		while (draw_frame == next_frame);
	}

	return 0;
}

