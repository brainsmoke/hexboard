#!/usr/bin/python3
import sys, math, time

n_leds = int(sys.argv[1])

while True:
	i = sys.stdin.buffer.read(n_leds*6 + 4)
	o = bytes(i[j*2+1] for j in range(n_leds))
	sys.stdout.buffer.write ( o )
	sys.stdout.buffer.flush()


