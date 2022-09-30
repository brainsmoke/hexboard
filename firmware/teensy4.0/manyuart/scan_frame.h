#ifndef SCAN_FRAME_H
#define SCAN_FRAME_H

/* Helper functions for a protocol which transfers frames with 16 bit values
 * over (usb) serial and allows the receiver to recover from a desynchonized
 * state.
 *
 * A frame consists of N 16 bit little endian values in the range [ 0x0000 ... 0xff00 ]
 * inclusive followed by an end of frame marker consistsing of 4 bytes: ff ff ff f0
 */

#include "stdint.h"

/* scan for an end of frame marker in a buffer,
 * In normal circumstances no marker should be present
 * before the end of the frame.
 *
 * - return -1 if no marker is found
 *
 * - if some data in the buffer belongs to the next frame.
 *   return the offset in the buffer where the next frame
 *   starts.
 * - if the end of frame marker is at the very end of the frame
 *   return the buffer length.
 * - if a desychronization problem is found, read from usb serial
 *   until the state is synchonized again, and return the buffer
 *   length.
 */
int scan_for_marker(uint8_t *buf, int len);


int eat_frame(void);

#endif
