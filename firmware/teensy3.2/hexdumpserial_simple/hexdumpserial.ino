/*
 * Drive 16 UART LED strip drivers chips at once
 *
 * Protocol:
 *   Frame:         [00-FD]*16384 [FE]
 *   Discard frame: [FF]
 *
 */

#include <HexSerialz.h>
#include <usb_dev.h>

#define N_DATA_BYTES_PER_STRIP 1024
#define N_BYTES_PER_STRIP (N_DATA_BYTES_PER_STRIP + 1)
#define N_DATA_BAUDS_PER_STRIP (10 * N_DATA_BYTES_PER_STRIP)
#define N_BAUDS_PER_STRIP (10 * N_BYTES_PER_STRIP)
#define N_STRIPS 16

#define END_OF_FRAME 0xfe
#define DISCARD_FRAME 0xff

uint16_t buf_a[N_BAUDS_PER_STRIP];
uint16_t buf_b[N_BAUDS_PER_STRIP];
uint16_t *cur, *next;
HexSerialz *hex;
const uint16_t strip_order[] =
{
	1<<12,
	1<<14,
	1<<15,
	1<<13,
	1<<11,
	1<<9,

	1<<6,
	1<<4,
	1<<2,
	1<<0,
	1<<1,
	1<<3,
	1<<5,
	1<<7,

	1<<8,
	1<<10,

	0,
};

void init_buf(uint16_t *buf)
{
	uint32_t i;

	for (i=0; i<8; i++)
		if ( END_OF_FRAME & (1<<i) )
			buf[N_DATA_BAUDS_PER_STRIP+1+i] = 0xffff;
		else
			buf[N_DATA_BAUDS_PER_STRIP+1+i] = 0x0000;

	for (i=0; i<N_BAUDS_PER_STRIP; i+=10)
	{
		buf[i] = 0x0000;
		buf[i+9] = 0xffff;
	}
}

static usb_packet_t *rx_packet=NULL;
static int rx_i=0, rx_len=0;

static int usb_getchar(void)
{
	if (rx_len <= rx_i)
	{
		if (rx_packet)
			usb_free(rx_packet);

		while ( !(rx_packet = usb_rx(CDC_RX_ENDPOINT)) || \
		         (rx_packet->index >= rx_packet->len)  );

		rx_i   = rx_packet->index;
		rx_len = rx_packet->len;
	}
	return (uint8_t)rx_packet->buf[rx_i++];
}

static int read_next_frame(void)
{
	int baud_ix=0, c, strip = 0;
	uint16_t strip_bit = strip_order[0];

	for(;;)
	{
		c = usb_getchar();

		if (c >= END_OF_FRAME)
			return c == END_OF_FRAME;

		if ( strip_bit )
		{
			uint16_t *p = &next[baud_ix];

			if (c&0x01) p[1] |=  strip_bit;
			else        p[1] &=~ strip_bit;

			if (c&0x02) p[2] |=  strip_bit;
			else        p[2] &=~ strip_bit;

			if (c&0x04) p[3] |=  strip_bit;
			else        p[3] &=~ strip_bit;

			if (c&0x08) p[4] |=  strip_bit;
			else        p[4] &=~ strip_bit;

			if (c&0x10) p[5] |=  strip_bit;
			else        p[5] &=~ strip_bit;

			if (c&0x20) p[6] |=  strip_bit;
			else        p[6] &=~ strip_bit;

			if (c&0x40) p[7] |=  strip_bit;
			else        p[7] &=~ strip_bit;

			if (c&0x80) p[8] |=  strip_bit;
			else        p[8] &=~ strip_bit;

			baud_ix += 10;

			if (baud_ix >= N_DATA_BAUDS_PER_STRIP)
			{
				baud_ix = 0;
				if (strip < N_STRIPS)
					strip++;
				strip_bit = strip_order[strip];
			}
		}
	}
}

void setup(void)
{
	usb_init();
}

void loop(void)
{
	memset(buf_a, 0, sizeof(buf_a));
	memset(buf_b, 0, sizeof(buf_b));
	init_buf(buf_a);
	init_buf(buf_b);
	cur = buf_a;
	next = buf_b;
	uint16_t *tmp;

    hex = new HexSerialz(N_BAUDS_PER_STRIP*2, 1000000);
    hex->begin();

    for (;;)
    {
		while(!read_next_frame());
        tmp=cur;
		cur=next;
		next=tmp;
        hex->show(cur);
		memcpy(next, cur, sizeof(buf_a));
    }
}
