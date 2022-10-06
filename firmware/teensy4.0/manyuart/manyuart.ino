
//#include <usb_serial.h>
#include <string.h>

#include "OctoUart.h"
#include "scan_frame.h"

#define DATA_BYTES_PER_LINE (2048)
const uint8_t marker[] = { '\xff', '\xff', '\xff', '\xf0' };

#define BYTES_PER_LINE (DATA_BYTES_PER_LINE+sizeof(marker))
#define N_PINS (15)
#define FRAMESIZE (BYTES_PER_LINE*N_PINS)

uint8_t surplus_buffer[DATA_BYTES_PER_LINE];
size_t surplus_len = 0;
uint8_t buf_a[N_PINS][BYTES_PER_LINE];
uint8_t buf_b[N_PINS][BYTES_PER_LINE];

const uint8_t pinlist[N_PINS] = {

                 /* v- desolder r1 to disable led */
	/* offsets: */ 13,16,/* hex bytes: */ 15,14,18,20,  6,8,12,9, /* ascii: */ 10,11,7,21,

	/* title bar: */ 19, /* unused: 17 */
};

OctoUart manyUart(1500000, 15, pinlist);

void setup(void)
{
	memset(surplus_buffer, 0, sizeof(surplus_buffer));
	memset(buf_a, 0, sizeof(buf_a));
	memset(buf_b, 0, sizeof(buf_b));
	int i=0;
	for (i=0; i<N_PINS; i++)
	{
		memcpy(&buf_a[i][BYTES_PER_LINE-sizeof(marker)], marker, sizeof(marker));
		memcpy(&buf_b[i][BYTES_PER_LINE-sizeof(marker)], marker, sizeof(marker));
	}

	Serial.begin(480000000);
	manyUart.begin();
}

int read_frame(uint8_t frame[N_PINS][BYTES_PER_LINE])
{
	int n = surplus_len, i, pos;

	memcpy(frame[0], surplus_buffer, surplus_len);
	surplus_len = 0;

	for (i=0; i<N_PINS; i++)
	{
		while (n < DATA_BYTES_PER_LINE)
			n += usb_serial_read(&frame[i][n], DATA_BYTES_PER_LINE-n);

		n = 0;

		/* returns offset to data belonging to the next frame */
		pos = scan_for_marker(&frame[i][0], DATA_BYTES_PER_LINE);
		if ( pos != -1 )
		{
			surplus_len = DATA_BYTES_PER_LINE-pos;
			memcpy(surplus_buffer, &frame[i][pos], surplus_len);
			return 0;
		}
	}

	return eat_frame();
}

void loop(void)
{
	uint8_t (* cur)[N_PINS][BYTES_PER_LINE] = &buf_a;
	uint8_t (* next)[N_PINS][BYTES_PER_LINE] = &buf_b;
	uint8_t (* tmp)[N_PINS][BYTES_PER_LINE];

	for (;;)
	{
		while (!read_frame(*next));
		tmp = cur;
		cur = next;
		next = tmp;
		manyUart.transmit(&(*cur)[0][0], BYTES_PER_LINE);
	};
}
