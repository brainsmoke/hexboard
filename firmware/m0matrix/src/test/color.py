#!/usr/bin/python3
# 
import sys, math, time

n_leds = int(sys.argv[1])

#r, g, b = ( min(0xff00, max(0, int(x, 16))) for x in sys.argv[2:5] )

for i, _ in enumerate(sys.stdin):
    x=i*0x10
    sys.stderr.write(hex(x))
    sys.stderr.flush()
    r,g,b=x,x,x
    sys.stdout.buffer.write ( bytes( [g&0xff,g>>8,r&0xff,r>>8,b&0xff,b>>8]*n_leds + [ 0xff,0xff,0xff,0xf0 ]) )
    sys.stdout.flush()


