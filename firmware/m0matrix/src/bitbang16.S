
#include <bitbang16.h>

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
	.syntax unified
	.thumb

#define rBUF   r0
#define rGPIO  r1
#define rCLK   r2
#define rOUT   r3
bitbang64_clk_stm32_start:

/*    0 + 2 */ ldm rBUF!, {rOUT}
/*    2 + 1 */ movs rCLK, #(MASK_CLK)
.rept 15
/* +  0 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  2 + 1 */ lsrs rOUT, #8
/* +  3 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* +  5 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  7 + 1 */ lsrs rOUT, #8
/* +  8 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 10 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 12 + 1 */ lsrs rOUT, #8
/* + 13 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 15 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 17 + 2 */ ldm rBUF!, {rOUT}
/* + 19 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
.endr
/*  318 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  320 + 1 */ lsrs rOUT, #8
/*  321 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  323 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  325 + 1 */ lsrs rOUT, #8
/*  326 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  328 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  330 + 1 */ lsrs rOUT, #8
/*  331 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  333 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
#undef rBUF
#define rTMP   r0
/*  335 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
/*  337 + 1 */ movs rTMP, #(BIT_LATCH|BIT_NOT_OUTPUT_ENABLE)
/*  338 + 1 */ orrs rOUT, rTMP
/*  339 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  341 + 1 */ movs rTMP, #(MASK_DATA|BIT_LATCH|BIT_NOT_OUTPUT_ENABLE)
/*  342 + 1 */ bics rOUT, rTMP
/*  343 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]

bitbang64_clk_stm32_end:
/*  345 + 3 */ bx lr
/*  348     */

#undef rTMP
#undef rCLK
#undef rGPIO
#undef rOUT
	.size	bitbang64_clk_stm32, .-bitbang64_clk_stm32




	.global	bitbang64_clk_no_enable_stm32
	.code	16
	.thumb_func
	.type	bitbang64_clk_no_enable_stm32, %function
bitbang64_clk_no_enable_stm32:
	.syntax unified
	.thumb

#define rBUF   r0
#define rGPIO  r1
#define rCLK   r2
#define rOUT   r3
bitbang64_clk_no_enable_stm32_start:

/*    0 + 2 */ ldm rBUF!, {rOUT}
/*    2 + 1 */ movs rCLK, #(MASK_CLK)
.rept 15
/* +  0 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  2 + 1 */ lsrs rOUT, #8
/* +  3 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* +  5 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* +  7 + 1 */ lsrs rOUT, #8
/* +  8 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 10 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 12 + 1 */ lsrs rOUT, #8
/* + 13 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/* + 15 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/* + 17 + 2 */ ldm rBUF!, {rOUT}
/* + 19 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
.endr
/*  318 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  320 + 1 */ lsrs rOUT, #8
/*  321 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  323 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  325 + 1 */ lsrs rOUT, #8
/*  326 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  328 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  330 + 1 */ lsrs rOUT, #8
/*  331 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]

/*  333 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
#undef rBUF
#define rTMP   r0
/*  335 + 2 */ str rCLK, [rGPIO, #GPIO_BSRR_OFFSET]
/*  337 + 1 */ movs rTMP, #(BIT_LATCH|BIT_NOT_OUTPUT_ENABLE)
/*  338 + 1 */ orrs rOUT, rTMP
/*  339 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]
/*  341 + 1 */ movs rTMP, #(MASK_DATA|BIT_LATCH)
/*  342 + 1 */ bics rOUT, rTMP
/*  343 + 2 */ strb rOUT, [rGPIO, #GPIO_ODR_OFFSET]

bitbang64_clk_no_enable_stm32_end:
/*  345 + 3 */ bx lr
/*  348     */

#undef rTMP
#undef rCLK
#undef rGPIO
#undef rOUT
	.size	bitbang64_clk_no_enable_stm32, .-bitbang64_clk_no_enable_stm32

	.global	precomp64
	.code	16
	.thumb_func
	.type	precomp64, %function
precomp64:
	.syntax unified
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
/* +  3 + 1 */ lsrs rA, rBIT
/* +  4 + 1 */ ands rA, rMASK
#if PIN_DATA_NORTH != 0
/* +  5 + 1 */ lsls rA, #(PIN_DATA_NORTH)
#endif

/* +  6 + 1 */ lsrs rB, rBIT
/* +  7 + 1 */ ands rB, rMASK
#if PIN_DATA_SOUTH != 0
/*          */ lsls rB, #(PIN_DATA_SOUTH)
#endif
/* +  8 + 1 */ orrs rA, rB
/* +  9 + 1 */ orrs rA, rOTHERPINS

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

//	.global	dithcomp128
//	.code	16
//	.thumb_func
//	.type	dithcomp128, %function
//dithcomp128:
//	.syntax unified
//	.thumb
//
//#define rRESIDUAL r0
//#define rFRAME    r1
//#define rMASK     r2
//#define rRESA     r3
//#define rRESB     r4
//#define rA        r5
//#define rB        r6
//dithcomp128_start:
///*    0 + 4 */ push {r4-r6}
///*    4 + 2 */ ldr rMASK,    =0x7f7f7f7f
//
//.rept 16
///* +  0 + 3 */ ldm rRESIDUAL!, {rRESA, rRESB}
///* +  3 + 3 */ ldm rFRAME!, {rA, rB}
///* +  6 + 1 */ subs rRESIDUAL, #8
///* +  7 + 1 */ ands rRESA, rMASK
///* +  8 + 1 */ ands rRESB, rMASK
///* +  9 + 1 */ ands rA, rMASK
///* + 10 + 1 */ ands rB, rMASK
///* + 11 + 1 */ adds rRESA, rA
///* + 12 + 1 */ adds rRESB, rB
///* + 13 + 3 */ stm rRESIDUAL!, {rRESA, rRESB}
//.endr
//
///*  262 + 4 */ pop {r4-r6}
//dithcomp128_end:
///*  266 + 3 */ bx lr
///*  269 */
//
//	.size	dithcomp128, .-dithcomp128
//#undef rRESIDUAL
//#undef rFRAME
//#undef rMASK
//#undef rRESA
//#undef rRESB
//#undef rA
//#undef rB

	.global	write_wait_write
	.code	16
	.thumb_func
	.type	write_wait_write, %function
write_wait_write:
	.syntax unified
	.thumb

cmp r3, #11
bge 11f
cmp r3, #9
bgt 10f
beq 9f
cmp r3, #7
bgt 8f
beq 7f
cmp r3, #5
bgt 6f
beq 5f
cmp r3, #3
bgt 4f
beq 3f
2:
str r1, [r0]
str r2, [r0]
bx lr
3:
str r1, [r0]
delay r3, 1
str r2, [r0]
bx lr
4:
str r1, [r0]
delay r3, 2
str r2, [r0]
bx lr
5:
str r1, [r0]
delay r3, 3
str r2, [r0]
bx lr
6:
str r1, [r0]
delay r3, 4
str r2, [r0]
bx lr
7:
str r1, [r0]
delay r3, 5
str r2, [r0]
bx lr
8:
str r1, [r0]
delay r3, 6
str r2, [r0]
bx lr
9:
str r1, [r0]
delay r3, 7
str r2, [r0]
bx lr
10:
str r1, [r0]
delay r3, 8
str r2, [r0]
bx lr
11:
str r1, [r0]
delay_reg r3, r1, 2
str r2, [r0]
bx lr

	.size	write_wait_write, .-write_wait_write

