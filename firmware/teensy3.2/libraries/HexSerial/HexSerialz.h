/*  HexSerialz - an OctoWS2811 fork to send 16 UART streams in parallel
    using a Teensy 3.x
    by Erik Bosman <erik@minemu.org> Zero copy version, similar to the Fadecandy
    implementation.

    Based on:

    OctoWS2811 - High Performance WS2811 LED Display Library
    http://www.pjrc.com/teensy/td_libs_OctoWS2811.html
    Copyright (c) 2013 Paul Stoffregen, PJRC.COM, LLC

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
    THE SOFTWARE.
*/

#ifndef HexSerialz_h
#define HexSerialz_h

#include <Arduino.h>
#include "DMAChannel.h"

#if TEENSYDUINO < 121
#error "Teensyduino version 1.21 or later is required to compile this library."
#endif
#ifdef __AVR__
#error "HexSerialz does not work with Teensy 2.0 or Teensy++ 2.0."
#endif

class HexSerialz {
public:
	HexSerialz(uint32_t bufsize, int freq=2000000);
	void begin(void);
	void show(void *frameBuffer);

private:
	static uint32_t bufsize;
	static int freq;
	static DMAChannel dma;
	static void isr(void);
	static void initGPIO(void);
};

#endif
