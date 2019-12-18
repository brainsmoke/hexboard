// cycles == n * 4 - 1
.macro delay_loop_4n_m1 reg iter
mov \reg, #(\iter)
0:
sub \reg, #1
bne 0b
.endm

// cycles == n * 5 - 1
.macro delay_loop_5n_m1 reg iter
mov \reg, #(\iter)
0:
nop
sub \reg, #1
bne 0b
.endm

// cycles == n * 7 - 1
.macro delay_loop_7n_m1 reg iter
mov \reg, #(\iter)
0:
b 1f
1:
sub \reg, #1
bne 0b
.endm

.macro delay reg num
	.ifge \num-11
		delay_loop \reg, \num

	.else
		.ifge \num-3
			b 1f
			1:
			delay \reg, (\num-3)
		.else
			.ifge \num-1
				nop
				delay \reg, (\num-1)

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


