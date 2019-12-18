
#include <bitbang.h>

.include "m0delay.s"

	.cpu cortex-m0
	.fpu softvfp
	.section	.bss
	.thumb
	.section	.text.bitbang,"ax",%progbits
	.align	2
	.global	bitbang64_clk_stm32
	.code	16
	.thumb_func
	.type	bitbang64_clk_stm32, %function
bitbang64_clk_stm32:
	.syntax divided
	.thumb

#define rBUF   r0
#define rGPIO  r1
#define rCLK   r2
#define rOUT   r3
bitbang64_clk_stm32_start:

/*    0 + 2 */ ldm rBUF!, {rOUT}
/*    2 + 1 */ mov rCLK, #(MASK_CLK)
.rept 15
/* +  0 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  2 + 1 */ lsr  rOUT, #8
/* +  3 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* +  5 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  7 + 1 */ lsr  rOUT, #8
/* +  8 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 10 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 12 + 1 */ lsr  rOUT, #8
/* + 13 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 15 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 17 + 2 */ ldm rBUF!, {rOUT}
/* + 19 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
.endr
/*  318 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  320 + 1 */ lsr  rOUT, #8
/*  321 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  323 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  325 + 1 */ lsr  rOUT, #8
/*  326 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  328 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  330 + 1 */ lsr  rOUT, #8
/*  331 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  333 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
#undef rBUF
#define rTMP   r0
/*  335 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
/*  337 + 1 */ mov rTMP, #(BIT_LATCH|BIT_NOT_OUTPUT_ENABLE)
/*  338 + 1 */ orr rOUT, rTMP
/*  339 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  341 + 1 */ mov rTMP, #(MASK_DATA|BIT_LATCH|BIT_NOT_OUTPUT_ENABLE)
/*  342 + 1 */ bic rOUT, rTMP
/*  343 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]

bitbang64_clk_stm32_end:
/*  345 + 3 */ bx lr
/*  348     */

#undef rTMP
#undef rCLK
#undef rGPIO
#undef rOUT
	.size	bitbang64_clk_stm32, .-bitbang64_clk_stm32

	.global	precomp64
	.code	16
	.thumb_func
	.type	precomp64, %function
precomp64:
	.syntax divided
	.thumb

#define rBUF     r0
#define rFRAME   r1
#define rBIT     r2
#define rOTHERPINS r3
#define rA       r4
#define rB       r5
#define rMASK    r6
precomp64_start:
/*    0 + 4 */ push {r4-r6}
/*    4 + 2 */ ldr rMASK,    =0x01010101

.rept 16
/* +  0 + 3 */ ldm rFRAME!, {rA, rB}
/* +  3 + 1 */ lsr  rA, rBIT
/* +  4 + 1 */ and rA, rMASK
#if PIN_DATA_NORTH != 0
/* +  5 + 1 */ lsl  rA, #(PIN_DATA_NORTH)
#endif

/* +  6 + 1 */ lsr  rB, rBIT
/* +  7 + 1 */ and rB, rMASK
#if PIN_DATA_SOUTH != 0
/*          */ lsl  rB, #(PIN_DATA_SOUTH)
#endif
/* +  8 + 1 */ orr rA, rB
/* +  9 + 1 */ orr rA, rOTHERPINS

/* + 10 + 2 */ stm rBUF!, {rA}
.endr

/*  198 + 4 */ pop {r4-r6}
precomp64_end:
/*  202 + 3 */ bx lr
/*  205 */

	.size	precomp64, .-precomp64
#undef rBUF
#undef rFRAME
#undef rBIT
#undef rA
#undef rB
#undef rMASK
#undef rOTHERPINS

	.global	dithcomp128
	.code	16
	.thumb_func
	.type	dithcomp128, %function
dithcomp128:
	.syntax divided
	.thumb

#define rRESIDUAL r0
#define rFRAME    r1
#define rMASK     r2
#define rRESA     r3
#define rRESB     r4
#define rA        r5
#define rB        r6
dithcomp128_start:
/*    0 + 4 */ push {r4-r6}
/*    4 + 2 */ ldr rMASK,    =0x7f7f7f7f

.rept 16
/* +  0 + 3 */ ldm rRESIDUAL!, {rRESA, rRESB}
/* +  3 + 3 */ ldm rFRAME!, {rA, rB}
/* +  6 + 1 */ sub rRESIDUAL, #8
/* +  7 + 1 */ and rRESA, rMASK
/* +  8 + 1 */ and rRESB, rMASK
/* +  9 + 1 */ add rRESA, rA
/* + 10 + 1 */ add rRESB, rB
/* + 11 + 3 */ stm rRESIDUAL!, {rRESA, rRESB}
.endr

/*  230 + 4 */ pop {r4-r6}
dithcomp128_end:
/*  234 + 3 */ bx lr
/*  237 */

	.size	dithcomp128, .-dithcomp128
#undef rRESIDUAL
#undef rFRAME
#undef rMASK
#undef rRESA
#undef rRESB
#undef rA
#undef rB

	.global	wait_a_bit
	.code	16
	.thumb_func
	.type	wait_a_bit, %function
wait_a_bit:
	.syntax divided
	.thumb
wait_a_bit_start:
delay r0, 100
wait_a_bit_end:
bx lr

	.size	wait_a_bit, .-wait_a_bit
