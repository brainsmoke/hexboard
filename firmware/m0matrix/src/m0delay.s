// cycles == n * 4 - 1
.macro delay_loop_4n_m1 reg iter
movs \reg, #(\iter)
0:
subs \reg, #1
bne 0b
.endm

// cycles == n * 5 - 1
.macro delay_loop_5n_m1 reg iter
movs \reg, #(\iter)
0:
nop
subs \reg, #1
bne 0b
.endm

// cycles == n * 7 - 1
.macro delay_loop_7n_m1 reg iter
movs \reg, #(\iter)
0:
b 1f
1:
subs \reg, #1
bne 0b
.endm

.macro delay reg num
	.ifge \num-11
		delay_loop \reg, \num

	.else
		.ifge \num-4
			b 1f
			1:
			delay \reg, (\num-3)
		.else
			.ifge \num-1
				delay \reg, (\num-1)
				movs \reg, #0

			.endif
		.endif
;	.endif
.endm

.macro delay_loop reg num
	.ifeq (\num+1)%4
		delay_loop_4n_m1 \reg, (\num+1)/4

	.else
		.ifeq (\num+1)%5
			delay_loop_5n_m1 \reg, (\num+1)/5

		.else
			.ifeq (\num+1)%7
				delay_loop_7n_m1 \reg, (\num+1)/7

			.else
				delay \reg, (\num+1)%4
				delay_loop_4n_m1 \reg, (\num+1)/4
			.endif
		.endif
	.endif
.endm

.macro addsub reg num
	.ifge \num
	adds \reg, #(\num)
	.else
	subs \reg, #(-\num)
	.endif
.endm

.macro delay_big reg num
	delay \reg (\num%4)
	ldr \reg, =(\num>>2)   /*     0 + 2       */
1:	subs \reg, #1          /*     2 + reg     */
	bne 1b                 /* 2+reg + 3*reg-2 */
	                       /* 4*reg           */
.endm



.macro delay_reg_div8 reg scratch sub cycleadj


                                              //  [reg&3==0] [reg&3==1] [reg&3==2] [reg&3==3]

   mov \scratch, \reg                         //       1          1          1          1
   addsub \scratch, (\cycleadj-(\sub+5)*8)    //       1          1          1          1
   lsrs \scratch, \scratch, #4                //       1          1          1          1
   bcs 1f                                     //       1          3          1          3
   nop                                        //       1                     1
1: lsrs \scratch, \scratch, #1                //       1          1          1          1
   bcs 2f                                     //       1          1          3          3
2: subs \scratch, #1                          //                    reg>>2
   bne 2b                                     //                 3*(reg>>2)-2

.endm


.macro delay_reg reg scratch sub


                                              //  [reg&3==0] [reg&3==1] [reg&3==2] [reg&3==3]

   mov \scratch, \reg                         //       1          1          1          1
   addsub \scratch, (-(\sub+5))               //       1          1          1          1
   lsrs \scratch, \scratch, #1                //       1          1          1          1
   bcs 1f                                     //       1          3          1          3
   nop                                        //       1                     1
1: lsrs \scratch, \scratch, #1                //       1          1          1          1
   bcs 2f                                     //       1          1          3          3
2: subs \scratch, #1                          //                    reg>>2
   bne 2b                                     //                 3*(reg>>2)-2

.endm




