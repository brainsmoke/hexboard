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

#include <string.h>
#include "HexSerialz.h"

uint32_t HexSerialz::bufsize;
int HexSerialz::freq;
DMAChannel HexSerialz::dma;

static volatile uint8_t update_in_progress = 0;

HexSerialz::HexSerialz(uint32_t bufsize, int freq)
{
	this->bufsize = bufsize;
	this->freq = freq;
}

void HexSerialz::initGPIO(void)
{
	GPIOC_PDOR = 0xff;
	GPIOD_PDOR = 0xff;
}

void HexSerialz::begin(void)
{
	initGPIO();
	// configure the 16 output pins

	pinMode(15, OUTPUT);	// strip #1
	pinMode(22, OUTPUT);	// strip #2
	pinMode(23, OUTPUT);	// strip #3
	pinMode(9, OUTPUT);	// strip #4
	pinMode(10, OUTPUT);	// strip #5
	pinMode(13, OUTPUT);	// strip #6
	pinMode(11, OUTPUT);	// strip #7
	pinMode(12, OUTPUT);	// strip #8

	pinMode(2, OUTPUT);	// strip #9
	pinMode(14, OUTPUT);	// strip #10
	pinMode(7, OUTPUT);	// strip #11
	pinMode(8, OUTPUT);	// strip #12
	pinMode(6, OUTPUT);	// strip #13
	pinMode(20, OUTPUT);	// strip #14
	pinMode(21, OUTPUT);	// strip #15
	pinMode(5, OUTPUT);	// strip #16

	analogWriteResolution(8);
	analogWriteFrequency(3, this->freq);
	analogWrite(3, 128);

	// pin 16 triggers DMA(port B) on rising edge (configure for pin 3's waveform)
	CORE_PIN16_CONFIG = PORT_PCR_IRQC(1)|PORT_PCR_MUX(3)|PORT_PCR_PE;
	pinMode(3, INPUT_PULLUP); // pin 3 no longer needed

//	dma.TCD->SADDR = frameBuffer;
	dma.TCD->SADDR = NULL;
	dma.TCD->SOFF = 2;
	dma.TCD->ATTR_SRC = DMA_TCD_ATTR_SIZE_16BIT;
	dma.TCD->SLAST = -bufsize;

	/* Send data to both PORT C and D in the same minor loop (executed after the same trigger( */
	#define PORT_DELTA ( (uint32_t)&GPIOD_PDOR - (uint32_t)&GPIOC_PDOR )
    dma.TCD->DADDR = &GPIOC_PDOR;
	dma.TCD->DOFF = PORT_DELTA;
                        /* loop GPIOC_PDOR, GPIOD_PDOR and back */
	dma.TCD->ATTR_DST = ((31 - __builtin_clz(PORT_DELTA*2)) << 3) | DMA_TCD_ATTR_SIZE_8BIT;
	dma.TCD->DLASTSGA = 0;

	dma.TCD->NBYTES = 2;
	dma.TCD->BITER = bufsize / 2;
	dma.TCD->CITER = bufsize / 2;

	dma.disableOnCompletion();
	dma.interruptAtCompletion();

#ifdef __MK20DX256__
	MCM_CR = MCM_CR_SRAMLAP(1) | MCM_CR_SRAMUAP(0);
	AXBS_PRS0 = 0x1032; /* not sure how this maps to DMA channels */
#endif

	dma.triggerAtHardwareEvent(DMAMUX_SOURCE_PORTB);
	dma.attachInterrupt(isr);
}

void HexSerialz::isr(void)
{
	dma.clearInterrupt();
	initGPIO();
	update_in_progress = 0;
}

void HexSerialz::show(void *frameBuffer)
{
	while (update_in_progress) ; 
	update_in_progress = 1;

	dma.TCD->SADDR = frameBuffer;

	uint32_t sc = FTM1_SC;
	FTM1_SC = sc & 0xE7;	// stop FTM1 timer

	FTM1_CNTIN = FTM1_MOD-1;
	FTM1_CNT = 0xbeef;      // FTM1_CNT == FTM1_MOD
	FTM1_CNTIN = 0;

	PORTB_ISFR = (1<<0);    // clear any prior rising edge
	dma.enable();		// enable all 3 DMA channels
	FTM1_SC = sc;		// restart FTM1 timer
}

