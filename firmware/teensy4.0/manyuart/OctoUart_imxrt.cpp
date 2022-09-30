/*  OctoUART - (synchronous) parallel uart tx Library by Erik Bosman,

    Largely copied from, (mostly deleted stuff :-D):

    OctoWS2811 - High Performance WS2811 LED Display Library
    http://www.pjrc.com/teensy/td_libs_OctoWS2811.html
    Copyright (c) 2020 Paul Stoffregen, PJRC.COM, LLC

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

#include <Arduino.h>
#include "OctoUart.h"

#if defined(__IMXRT1062__)

#define PERF_DEBUG_PIN (5)

// Ordinary RGB data is converted to GPIO bitmasks on-the-fly using
// a transmit buffer sized for 2 DMA transfers.  The larger this setting,
// the more interrupt latency OctoUart can tolerate, but the transmit
// buffer grows in size.  For good performance, the buffer should be kept
// smaller than the half the Cortex-M7 data cache.
#define BYTES_PER_DMA	40
#define BAUDS_PER_BYTE  10
#define DMA_ITERATIONS(count) (count*BAUDS_PER_BYTE)
#define DMA_CHUNK_INTERATIONS (DMA_ITERATIONS(BYTES_PER_DMA))

#define NUM_GPIO_REGS    4
#define N_CHUNKS         2
#define DMA_CHUNK_WRITES (DMA_CHUNK_INTERATIONS*NUM_GPIO_REGS)
#define DMA_CHUNK_SIZE (DMA_CHUNK_SIZE*sizeof(uint32_t))

uint32_t OctoUart::baudrate;
uint8_t OctoUart::defaultPinList[8] = {2, 14, 7, 8, 6, 20, 21, 5};
uint8_t *OctoUart::buffer;
DMAChannel OctoUart::dma;
static DMASetting dma_next;
static uint32_t numbytes;

static uint8_t numpins;
static uint8_t pinlist[NUM_DIGITAL_PINS]; // = {2, 14, 7, 8, 6, 20, 21, 5};
static uint8_t pin_bitnum[NUM_DIGITAL_PINS];
static uint8_t pin_offset[NUM_DIGITAL_PINS];

DMAMEM static uint32_t bitmask[4] __attribute__ ((used, aligned(32)));
DMAMEM static uint32_t bitdata[DMA_CHUNK_WRITES*N_CHUNKS] __attribute__ ((used, aligned(32)));
volatile uint32_t framebuffer_index = 0;
volatile uint8_t *framebuffer;
volatile bool dma_first;
volatile bool dma_done;

OctoUart::OctoUart(uint32_t baudrate, uint8_t numPins, const uint8_t *pinList)
{
	this->baudrate = baudrate;
	if (numPins > NUM_DIGITAL_PINS) numPins = NUM_DIGITAL_PINS;
	numpins = numPins;
	memcpy(pinlist, pinList, numpins);
}


extern "C" void xbar_connect(unsigned int input, unsigned int output); // in pwm.c
static volatile uint32_t *standard_gpio_addr(volatile uint32_t *fastgpio) {
	return (volatile uint32_t *)((uint32_t)fastgpio - 0x01E48000);
}

void OctoUart::begin(void)
{
#ifdef PERF_DEBUG_PIN
	digitalWrite(PERF_DEBUG_PIN, LOW);
	pinMode(PERF_DEBUG_PIN, OUTPUT);
	digitalWrite(PERF_DEBUG_PIN, LOW);
#endif

	// configure which pins to use
	memset(bitmask, 0, sizeof(bitmask));
	for (uint32_t i=0; i < numpins; i++) {
		uint8_t pin = pinlist[i];
		if (pin >= NUM_DIGITAL_PINS) continue; // ignore illegal pins
		uint8_t bit = digitalPinToBit(pin);
		uint8_t offset = ((uint32_t)portOutputRegister(pin) - (uint32_t)&GPIO6_DR) >> 14;
		if (offset > 3) continue; // ignore unknown pins
		pin_bitnum[i] = bit;
		pin_offset[i] = offset;
		uint32_t mask = 1 << bit;
		bitmask[offset] |= mask;
		*(&IOMUXC_GPR_GPR26 + offset) &= ~mask;
		*standard_gpio_addr(portSetRegister(pin)) |= mask;
		*standard_gpio_addr(portModeRegister(pin)) |= mask;
		*standard_gpio_addr(portSetRegister(pin)) |= mask;
	}
	arm_dcache_flush_delete(bitmask, sizeof(bitmask));

	TMR4_ENBL &= ~1;
	TMR4_SCTRL0 = TMR_SCTRL_OEN | TMR_SCTRL_FORCE | TMR_SCTRL_MSTR;
	TMR4_CSCTRL0 = TMR_CSCTRL_CL1(1) | TMR_CSCTRL_TCF1EN;
	TMR4_CNTR0 = 0;
	TMR4_LOAD0 = 0;
	TMR4_COMP10 = (uint16_t)(F_BUS_ACTUAL/baudrate);
	TMR4_CMPLD10 = (uint16_t)(F_BUS_ACTUAL/baudrate);
	TMR4_CTRL0 = TMR_CTRL_CM(1) | TMR_CTRL_PCS(8) | TMR_CTRL_LENGTH | TMR_CTRL_OUTMODE(3);

	// route the timer outputs through XBAR to edge trigger DMA request
	CCM_CCGR2 |= CCM_CCGR2_XBAR1(CCM_CCGR_ON);
	xbar_connect(XBARA1_IN_QTIMER4_TIMER0, XBARA1_OUT_DMA_CH_MUX_REQ30);
	XBARA1_CTRL0 = XBARA_CTRL_STS0 | XBARA_CTRL_EDGE0(3) | XBARA_CTRL_DEN0;

	// configure DMA channels
	dma_next.TCD->SADDR = bitdata;
	dma_next.TCD->SOFF = 8;
	dma_next.TCD->ATTR = DMA_TCD_ATTR_SSIZE(3) | DMA_TCD_ATTR_DSIZE(2);
	dma_next.TCD->NBYTES_MLOFFYES = DMA_TCD_NBYTES_DMLOE |
		DMA_TCD_NBYTES_MLOFFYES_MLOFF(-65536) |
		DMA_TCD_NBYTES_MLOFFYES_NBYTES(16);
	dma_next.TCD->SLAST = 0;
	dma_next.TCD->DADDR = &GPIO1_DR_TOGGLE;
	dma_next.TCD->DOFF = 16384;
	dma_next.TCD->CITER_ELINKNO = DMA_CHUNK_INTERATIONS;
	dma_next.TCD->DLASTSGA = (int32_t)(dma_next.TCD);
	dma_next.TCD->BITER_ELINKNO = DMA_CHUNK_INTERATIONS;

	dma.begin();
	dma = dma_next; // copies TCD
	dma.triggerAtHardwareEvent(DMAMUX_SOURCE_XBAR1_0);
	dma.attachInterrupt(isr);
	dma_done = true;
}

static void fillbits(uint32_t *dest, const uint8_t *pixels, int n, uint32_t mask)
{
	do {
		uint8_t pix = *pixels++; 
		/* the XOR logic is because we write to DR_TOGGLE */
		uint32_t flips = pix ^ (pix<<1);

		*dest |= mask; /* START BIT == 0 */
		dest += 4;
		if ( flips & 0x01 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x02 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x04 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x08 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x10 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x20 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x40 ) *dest |= mask;
		dest += 4;
		if ( flips & 0x80 ) *dest |= mask;
		dest += 4;
		if ( !(flips & 0x100) ) *dest |= mask; /* STOP BIT == 1 */
		dest += 4;
	} while (--n > 0);
}

int OctoUart::busy(void)
{
	return !dma_done;
}

void OctoUart::transmit(uint8_t *buffer, uint32_t bytesPerOutput)
{
	while (!dma_done) ; // can't use dma.complete() because we reload

	dma_done = false;
	numbytes = bytesPerOutput;
	framebuffer = buffer;
	// wait for any prior DMA operation

	// disable timers
	uint16_t enable = TMR4_ENBL;
	TMR4_ENBL = enable & ~1;

	// force all timer outputs to logic low
	TMR4_SCTRL0 = TMR_SCTRL_OEN | TMR_SCTRL_FORCE | TMR_SCTRL_MSTR;

	// clear any prior pending DMA requests
	XBARA1_CTRL0 |= XBARA_CTRL_STS0;

	// fill the DMA transmit buffer
	memset(bitdata, 0, sizeof(bitdata));
	uint32_t count = numbytes;
	if (count > BYTES_PER_DMA*N_CHUNKS) count = BYTES_PER_DMA*N_CHUNKS;
	framebuffer_index = count;
	for (uint32_t i=0; i < numpins; i++) {
		fillbits(bitdata + pin_offset[i], (uint8_t *)framebuffer + i*numbytes,
			count, 1<<pin_bitnum[i]);
	}
	arm_dcache_flush_delete(bitdata, count * BAUDS_PER_BYTE*NUM_GPIO_REGS*sizeof(uint32_t));

    // set up DMA transfers
	dma.TCD->SADDR = bitdata;
	dma.TCD->DADDR = &GPIO1_DR_TOGGLE;
	dma.TCD->CSR = 0; // important!

	if (numbytes <= BYTES_PER_DMA*N_CHUNKS)
	{
		dma.TCD->CITER_ELINKNO = DMA_ITERATIONS(count);
		dma.TCD->CSR = DMA_TCD_CSR_DREQ | DMA_TCD_CSR_INTMAJOR;
	}
	else
	{
		dma.TCD->CSR = DMA_TCD_CSR_ESG | DMA_TCD_CSR_INTMAJOR;
		dma.TCD->CITER_ELINKNO = DMA_CHUNK_INTERATIONS;
		dma_next.TCD->SADDR = &bitdata[DMA_CHUNK_WRITES];
		dma_next.TCD->CITER_ELINKNO = DMA_CHUNK_INTERATIONS;
		if (numbytes <= BYTES_PER_DMA*(N_CHUNKS+1))
			dma_next.TCD->CSR = DMA_TCD_CSR_ESG; /* no interrupt before last loop */
		else
			dma_next.TCD->CSR = DMA_TCD_CSR_ESG | DMA_TCD_CSR_INTMAJOR; /* interrupt at end */
	}
	dma_first = true;
	dma.enable();

	// initialize timers
	TMR4_CNTR0 = 0;

	// start everything running!
	TMR4_ENBL = enable | 1;
}

void OctoUart::isr(void)
{
	// first ack the interrupt
	dma.clearInterrupt();

#ifdef PERF_DEBUG_PIN
	digitalWriteFast(PERF_DEBUG_PIN, HIGH);
#endif

	uint32_t index = framebuffer_index;
	uint32_t count = numbytes - index;

	/* We don't throw an interrupt before the last loop since
	 * all output data has been processed by then, but we enable
	 * the interrupt for /after/ the last loop.
	 * So if we're at this point and all data has been processed,
	 * this must be the final interrupt after all data has been
	 * sent out.
	 */
	if (count == 0)
	{
		dma_done = true;

		for (int i=0; i < numpins; i++) {
			/* provide some robustness since we only use toggle in the DMA loop */
			/* set the known good value */
			*standard_gpio_addr(portSetRegister(pinlist[i])) |= 1<<pin_bitnum[i];;
		}

#ifdef PERF_DEBUG_PIN
		digitalWriteFast(PERF_DEBUG_PIN, LOW);
#endif
		return; 
	}

	// fill (up to) half the transmit buffer with new data
	uint32_t *dest;
	if (dma_first) {
		dma_first = false;
		dest = &bitdata[0];
	} else {
		dma_first = true;
		dest = &bitdata[DMA_CHUNK_WRITES];
	}
	memset(dest, 0, sizeof(bitdata)/2);
	if (count > BYTES_PER_DMA) count = BYTES_PER_DMA;
	framebuffer_index = index + count;
	for (int i=0; i < numpins; i++) {
		fillbits(dest + pin_offset[i], (uint8_t *)framebuffer + index + i*numbytes,
			count, 1<<pin_bitnum[i]);
	}
	arm_dcache_flush_delete(dest, count * BAUDS_PER_BYTE*NUM_GPIO_REGS*sizeof(uint32_t));

	// queue it for the next DMA transfer
	dma_next.TCD->SADDR = dest;
	dma_next.TCD->CITER_ELINKNO = DMA_ITERATIONS(count);
	uint32_t remain = numbytes - (index + count);
	if (remain == 0) {
		dma_next.TCD->CSR = DMA_TCD_CSR_DREQ | DMA_TCD_CSR_INTMAJOR;
	} else if (remain <= BYTES_PER_DMA) {
		dma_next.TCD->CSR = DMA_TCD_CSR_ESG; /* no interrupt before last loop */
	} else {
		dma_next.TCD->CSR = DMA_TCD_CSR_ESG | DMA_TCD_CSR_INTMAJOR; /* interrupt at end */
	}

#ifdef PERF_DEBUG_PIN
	digitalWriteFast(PERF_DEBUG_PIN, LOW);
#endif
}

#endif // __IMXRT1062__
