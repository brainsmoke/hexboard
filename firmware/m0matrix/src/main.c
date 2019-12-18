
/* Joining the FPGA cult next project, I promise */

#include <stdlib.h>
#include "stm32f0xx.h"

#include "bitbang.h"
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
 *                  [  0  1 ] } 2x 16 bytes (A B C D E F G1 G2 H J K L M N Dp nc)
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

uint8_t residual[N_LEDS];

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

/*
max_in=0xfd
max_val=0xffff
gamma=2.5

def cutoff(x, cutoff=32):
	if x < cutoff:
		x = 0
	return x

factor = max_val / (max_in**gamma)
gamma_map = [ int( (x**gamma * factor) + .5 ) for x in range(max_in+1) ]
gamma_map = [ cutoff(x) for x in gamma_map ]
gamma_map = [ (x&~0x3ff) | int((x&0x3ff)/8 + .5) for x in gamma_map ]

print ("""
const uint16_t gamma_map[] =
{
    """+','.join(str(x) for x in gamma_map)+""",
};""")

*/

static const uint16_t gamma_map[] =
{
	0,0,0,0,0,0,0,0,0,0,0,0,4,5,6,7,8,10,11,13,14,16,18,20,23,25,28,31,33,37,40,43,47,50,54,58,63,67,72,76,81,87,92,98,103,109,116,122,1024,1031,1038,1046,1053,1061,1068,1077,1085,1093,1102,1111,1120,1130,1140,1150,2056,2066,2077,2088,2099,2110,2122,2134,2146,2158,2171,3080,3093,3107,3120,3134,3149,3163,3178,3193,4104,4120,4136,4152,4169,4185,4202,4220,5133,5151,5169,5188,5207,5226,5245,6161,6181,6201,6222,6242,6264,7181,7203,7225,7247,7270,7293,8213,8236,8260,8285,8309,9230,9255,9281,9307,9333,10256,10283,10310,10338,10366,11290,11319,11348,11377,12302,12332,12363,12393,13320,13352,13384,13416,14344,14377,14410,14444,15373,15408,15442,15477,16408,16444,16480,17413,17449,17486,17524,18458,18496,18535,19470,19509,19549,20485,20526,20566,20608,21545,21587,21630,22569,22612,22656,23595,23640,24581,24626,24671,25613,25660,25707,26650,26697,26745,27690,27738,28684,28733,28783,29730,29780,30728,30779,30831,31780,31833,32782,32836,32890,33840,33895,34847,34903,35855,35912,36865,36922,36980,37935,37993,38949,39008,39964,40025,40982,41043,42001,42064,43022,43086,44045,44109,45070,45135,46096,46162,47124,47191,48154,48222,49186,49255,50220,50289,51255,51326,52292,53260,53331,54300,54372,55342,55415,56385,57356,57431,58402,58478,59451,60424,60501,61475,61553,62528,63503,63583,64559,64640,
};

enum
{
	B15, B14, B13, B12, BX, BY,
	P15, P14, P13, P12, P11, P10, PD, PR,
	ZX, /* send zeroes */
	ZZZ,
};


static const uint8_t dtable[68] = /* SysTick dispatch table for BCM bitbang & precomputation */
{
/*	 15    .    .    .    .    .    .    .   14    .    .    .   13    .   12  11/10 dith */

	B15, ZZZ, P14, P13, P12, ZZZ, P11,  PD, B14, ZZZ,  PR, P15, B13, ZZZ, B12,  BX,  BY,
	B15, ZZZ, P14, P13, P12, ZZZ, P10,  PD, B14, ZZZ,  PR, P15, B13, ZZZ, B12,  BX,  BY,
	B15, ZZZ, P14, P13, P12, ZZZ, P11,  PD, B14, ZZZ,  PR, P15, B13, ZZZ, B12,  BX,  BY,
	B15, ZZZ, P14, P13, P12, ZZZ,  ZX,  PD, B14, ZZZ,  PR, P15, B13, ZZZ, B12,  BX,  BY,
};

const uint32_t select_row[] =
{
	CLEAR ( BIT_SELECT_A ) | CLEAR ( BIT_SELECT_B ),
	SET   ( BIT_SELECT_A ) |   SET ( BIT_SELECT_B ),
	CLEAR ( BIT_SELECT_A ) |   SET ( BIT_SELECT_B ),
	SET   ( BIT_SELECT_A ) | CLEAR ( BIT_SELECT_B ),
};

static void bitbang_15(void)       { bitbang64_clk_stm32(buf15, (void *)GPIOA);
                                     GPIOA->BSRR = select_row[cur_row];
	                                 GPIOA->BSRR = SET(BIT_ENABLE_HIGH); }
static void bitbang_14(void)       { bitbang64_clk_stm32(buf14, (void *)GPIOA); }
static void bitbang_13(void)       { bitbang64_clk_stm32(buf13, (void *)GPIOA); }
static void bitbang_12(void)       { bitbang64_clk_stm32(buf12, (void *)GPIOA); }
static void bitbang_x(void)        { bitbang64_clk_stm32(bufx,  (void *)GPIOA); }
static void bitbang_y_switch(void) { bitbang64_clk_stm32(bufy,  (void *)GPIOA);
                                     cur_row = (cur_row+1)&(N_ROWS-1);
                                     if (cur_row) iter -= 17; }

static void prepare_15(void) { draw_frame = cur_frame;
                               cur_pos += N_BITS_PER_ROW; if (cur_pos>=N_LEDS) cur_pos = 0;
                               precomp64(buf15, &draw_frame->high_bits[cur_pos], 7, X4(BIT_NOT_OUTPUT_ENABLE));}
static void prepare_14(void) { precomp64(buf14, &draw_frame->high_bits[cur_pos], 6, X4(BIT_ENABLE_HIGH)); }
static void prepare_13(void) { precomp64(buf13, &draw_frame->high_bits[cur_pos], 5, X4(BIT_ENABLE_HIGH)); }
static void prepare_12(void) { precomp64(buf12, &draw_frame->high_bits[cur_pos], 4, X4(BIT_ENABLE_HIGH)); }
static void prepare_11(void) { precomp64(bufx,  &draw_frame->high_bits[cur_pos], 3, X4(BIT_ENABLE_HIGH)); }
static void prepare_10(void) { precomp64(bufx,  &draw_frame->high_bits[cur_pos], 2, X4(BIT_ENABLE_HIGH)); }

static void prepare_dith(void) { dithcomp128(&residual[cur_pos], &draw_frame->low_bits[cur_pos]); }
static void prepare_res(void)  { precomp64(bufy, &residual[cur_pos], 7, X4(BIT_ENABLE_HIGH)); }

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
	[BY] = bitbang_y_switch,
	[P15] = prepare_15,
	[P14] = prepare_14,
	[P13] = prepare_13,
	[P12] = prepare_12,
	[P11] = prepare_11,
	[P10] = prepare_10,
	[ZX] = zero_x,
	[PD] = prepare_dith,
	[PR] = prepare_res,
	[ZZZ]= ret,
};

void SysTick_Handler(void)
{
	dispatch[dtable[iter]]();

	if (iter < 67)
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
	usart1_rx_pa3_dma3_enable(recv_buf, RECV_BUF_SZ, 48e6/1e6);

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

    enable_sys_tick(F_SYS_TICK_CLK/27200/4);
}

static int dma_getchar(void)
{
	static uint32_t last = 0;
	if (last == 0)
		last = RECV_BUF_SZ;
	while (last == DMA1_Channel3->CNDTR);
	last--;
	return recv_buf[RECV_BUF_SZ-1-last];
}

static int read_next_frame(void)
{
	int i, c, v, digit_location, chip_delta, seg;

	if (is_2nd_half)
		for (i=0; i<N_LEDS; i++)
			if(dma_getchar() >= END_OF_FRAME)
				return 0;

	/* do the mapping to an easy-to-use internal representation once */
#define NEXT_CHIP (N_PINS_PER_CHIP*N_CHANNELS)
#define NEXT_ROW (N_BITS_PER_ROW)
#define SEGMENT_MASK (N_PINS_PER_CHIP-1)
#define CHIP_MASK (N_PINS_PER_CHIP*N_ROWS-1)

	chip_delta = NEXT_CHIP - N_ROWS*NEXT_ROW; /* pre-undo first iteration */
	digit_location = - chip_delta - NEXT_ROW;
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
		c = dma_getchar();
		if (c>=END_OF_FRAME)
			break;

		v = gamma_map[c];
		next_frame->high_bits[digit_location+segment_pin_interleaved[seg]] = v>>8;
		next_frame->low_bits [digit_location+segment_pin_interleaved[seg]] = v;
	}

	while (c<END_OF_FRAME)
		c = dma_getchar();

	return c != DISCARD_FRAME;
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

