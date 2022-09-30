/*  OctoUART - (synchronous) parallel uart tx Library by Erik Bosman,

    Largely copied from:

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

#ifndef OctoUart_h
#define OctoUart_h

#include <Arduino.h>

#ifdef __AVR__
#error "Sorry, OctoUart only works on 32 bit Teensy boards.  AVR isn't supported."
#endif

#if TEENSYDUINO < 121
#error "Teensyduino version 1.21 or later is required to compile this library."
#endif

#include "DMAChannel.h"


class OctoUart {
public:
#if defined(__IMXRT1062__)
	// Teensy 4.x can use any arbitrary group of pins!
	OctoUart(uint32_t baudrate, uint8_t numPins = 8, const uint8_t *pinList = defaultPinList);
	void begin(void);
#else
#error "Only implemented for Teensy 4.x"
#endif

	void transmit(uint8_t *buffer, uint32_t bytesPerOutput);
	int busy(void);

private:
	static uint32_t baudrate;
	static uint8_t *buffer;
	static DMAChannel dma;
	static void isr(void);
	static uint8_t defaultPinList[8];
};

#endif
