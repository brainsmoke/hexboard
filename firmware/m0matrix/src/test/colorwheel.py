#!/usr/bin/python3
# 
#  Copyright (c) 2017 Erik Bosman <erik@minemu.org>
# 
#  Permission  is  hereby  granted,  free  of  charge,  to  any  person
#  obtaining  a copy  of  this  software  and  associated documentation
#  files (the "Software"),  to deal in the Software without restriction,
#  including  without  limitation  the  rights  to  use,  copy,  modify,
#  merge, publish, distribute, sublicense, and/or sell copies of the
#  Software,  and to permit persons to whom the Software is furnished to
#  do so, subject to the following conditions:
# 
#  The  above  copyright  notice  and this  permission  notice  shall be
#  included  in  all  copies  or  substantial portions  of the Software.
# 
#  THE SOFTWARE  IS  PROVIDED  "AS IS", WITHOUT WARRANTY  OF ANY KIND,
#  EXPRESS OR IMPLIED,  INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
#  MERCHANTABILITY,  FITNESS  FOR  A  PARTICULAR  PURPOSE  AND
#  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
#  BE LIABLE FOR ANY CLAIM,  DAMAGES OR OTHER LIABILITY, WHETHER IN AN
#  ACTION OF CONTRACT,  TORT OR OTHERWISE,  ARISING FROM, OUT OF OR IN
#  CONNECTION  WITH THE SOFTWARE  OR THE USE  OR OTHER DEALINGS IN THE
#  SOFTWARE.
# 
#  (http://opensource.org/licenses/mit-license.html)
#

import sys, math, time

n_leds = int(sys.argv[1])

def gamma_map():

    max_val=0xff00
    cut_off = 0x18
    gamma=2.5

    factor = max_val / (255.**gamma)
    gamma16 = [ int(x**gamma * factor) for x in range(256) ]
    m = [0]*256

    for i,v in enumerate(gamma16):
        lo, hi = v&0xff, v>>8
        if lo <= cut_off/2:
            lo = 0
        elif lo < cut_off:
            lo = cut_off
        elif lo > 256-cut_off/2:
            lo, hi = 0, min(hi+1, max_val>>8)
        elif lo > 256-cut_off:
            lo = 256-cut_off
        m[i] = lo + hi*256
    return m

m = gamma_map()

wave = [int(-math.cos(x*math.pi*2/256)*127.5+127.5) for x in range(256)]
w = [ m[x] for x in wave ]
w = w+w

r, g, b = 0,85,170

while True:
    sys.stdout.buffer.write ( bytes( [(w[c+i*256//n_leds]>>sh)&0xff for i in range(n_leds) for c in (r,g,b) for sh in (0,8) ] + [ 0xff,0xff,0xff,0xf0 ]) )
    sys.stdout.flush()
    r, g, b = (r+1)%256, (g+1)%256, (b+1)%256
    #time.sleep(.0025)


