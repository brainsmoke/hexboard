
#include <Arduino.h>

#include "scan_frame.h"

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

#define STATE_COUNT   10

#define GOOD_RETURN  11
#define BAD_RETURN   12

#define IN_00     0
#define IN_01_EF  1
#define IN_F0     2
#define IN_F1_FE  3
#define IN_FF     4

const uint8_t fsm[STATE_COUNT][8] =
{
//	[GOOD]        = { [IN_00]=GOOD_00, [IN_01_EF]=GOOD_01_FE, [IN_F0]=GOOD_01_FE , [IN_F1_FE]=GOOD_01_FE, [IN_FF]=GOOD_FF     },
//	[GOOD_00]     = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=GOOD        },
//	[GOOD_01_FE]  = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=BAD_FF      },
//	[GOOD_FF]     = { [IN_00]=GOOD   , [IN_01_EF]=GOOD      , [IN_F0]=GOOD       , [IN_F1_FE]=GOOD      , [IN_FF]=GOOD_FFFF   },
//	[GOOD_FFFF]   = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=GOOD_FFFFFF },
//	[GOOD_FFFFFF] = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=GOOD_RETURN, [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },
//	[BAD]         = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FF      },
//	[BAD_FF]      = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFF    },
//	[BAD_FFFF]    = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD        , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },
//	[BAD_FFFFFF]  = { [IN_00]=BAD    , [IN_01_EF]=BAD       , [IN_F0]=BAD_RETURN , [IN_F1_FE]=BAD       , [IN_FF]=BAD_FFFFFF  },
//
	{ GOOD_00, GOOD_01_FE, GOOD_01_FE , GOOD_01_FE, GOOD_FF     },
	{ GOOD   , GOOD      , GOOD       , GOOD      , GOOD        },
	{ GOOD   , GOOD      , GOOD       , GOOD      , BAD_FF      },
	{ GOOD   , GOOD      , GOOD       , GOOD      , GOOD_FFFF   },
	{ BAD    , BAD       , BAD        , BAD       , GOOD_FFFFFF },
	{ BAD    , BAD       , GOOD_RETURN, BAD       , BAD_FFFFFF  },
	{ BAD    , BAD       , BAD        , BAD       , BAD_FF      },
	{ BAD    , BAD       , BAD        , BAD       , BAD_FFFF    },
	{ BAD    , BAD       , BAD        , BAD       , BAD_FFFFFF  },
	{ BAD    , BAD       , BAD_RETURN , BAD       , BAD_FFFFFF  },
};

const uint8_t input_lookup[] = {
	/* 0         1         2         3         4         5         6         7         8         9         A         B         C         D         E         F   */
	IN_00,    IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 0 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 1 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 2 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 3 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 4 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 5 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 6 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 7 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 8 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* 9 */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* A */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* B */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* C */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* D */
	IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, IN_01_EF, /* E */
	IN_F0,    IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_F1_FE, IN_FF   , /* F */	
};

static int state = GOOD;

int scan_for_marker(uint8_t *buf, int len)
{
	int i;
	for (i=0; i<len; i++)
	{
		state = fsm[state][input_lookup[buf[i]]];
		if ( state == GOOD_RETURN || state == BAD_RETURN )
		{
			state = GOOD;
			return i+1;
		}
	}

	if (state >= GOOD_FFFF)
	{
		eat_frame();
		return len;
	}

	return -1;
}

int eat_frame(void)
{
	uint8_t buf[1];

	while (state != GOOD_RETURN && state != BAD_RETURN)
	{
		while ( !usb_serial_read(buf, 1) );
		state = fsm[state][input_lookup[buf[0]]];
	}
	int ret = (state == GOOD_RETURN);
	state = GOOD;
	return ret;
}

