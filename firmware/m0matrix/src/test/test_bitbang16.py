
# arm-none-eabi-objdump -d main.elf | python test_bitbang.py
import sys, random

import thumb_emu

REMAINDERS=0x20000000
BUFFER=0x20000600
LEDS_SIZE=0x600
GPIO=0x48000014

def get_output_data(state):
	data = []
	t0=last=None
	for t, addr, val, pc in state['out']:

		if addr != GPIO:
			raise "meh."

		if t0 == None:
			if val == 0:
				continue
			if val != 255:
				raise "meh."
			t0 = t
			last = t-18

		ph = (t-t0)%60
		ph_diff = t-last

		if (ph,ph_diff) not in ( (0,18), (18,18), (42, 24) ):
			print t, hex(addr), val, hex(pc)
			raise "meh"

		if ph == 0 and val != 255:
			raise "meh."	

		if ph == 18 and val > 255:
			raise "meh."	

		if ph == 42 and val != 0:
			raise "meh."	

		if ph == 18:
			data += [val]

		last = t

	if (state['out'][-1][0]-t0)%60 != 42:
		raise "Meh"

	return tuple(data)

def run_code(start_pc, end_pc, code, mem, remainders, buf):

	remainders = list(remainders)
	buf = list(buf)
	mem = dict(mem)
	thumb_emu.write_mem(mem, REMAINDERS, remainders)
	thumb_emu.write_mem(mem, BUFFER, thumb_emu.from_le_array(buf, 2))
	state = thumb_emu.get_state(start_pc, end_pc, code, mem, BUFFER, GPIO)
	thumb_emu.run(state)
	data = get_output_data(state)
	remainders = thumb_emu.read_mem(mem, REMAINDERS, LEDS_SIZE)
	buf = thumb_emu.to_le_array(thumb_emu.read_mem(mem, BUFFER, LEDS_SIZE*2), 2)
	return tuple(data), tuple(remainders), tuple(buf)

def scatter_gather(m):
	n = [0]*8
	for i,val in enumerate(m):
		if not 0 <= val <= 255:
			raise "meh."
		for j in range(8):
			if val & (1<<(7-j)):
				n[j] |= (1<<(7-i))
	return tuple(n)

def run_algo(remainders, buf):
	remainders, buf = list(remainders), list(buf)
	leds_per_strip = len(buf)/8
	data = []
	for i in xrange(leds_per_strip):
		m = [0]*8

		for j in xrange(8):
			ix = i+j*leds_per_strip
			v16 = buf[ix]+remainders[ix]
			remainders[ix] = v16&0xff
			m[j] = v16>>8

		data += scatter_gather(m)

	return tuple(data), tuple(remainders), tuple(buf)


def run_test(start_pc, end_pc, code, mem, remainders, buf):
	d1, r1, b1 = run_algo(remainders, buf)
	d2, r2, b2 = run_code(start_pc, end_pc, code, mem, remainders, buf)

	if b1 != b2:
		raise "meh"

	if d1 != d2:
		for i, e1, e2 in zip(range(len(d1)), d1, d2):
			print i, bin(e1), bin(e2),
			if e1 != e2:
				print "XXXX",
			print
		raise "meh"

	if r1 != r2:
		raise "meh"

def run_tests(n):

	start_pc, end_pc, code, mem = thumb_emu.read_code(sys.stdin, 'bitbang_start', 'bitbang_end')

	for _ in xrange(n):
		remainders = [ random.randint(0, 0xff)   for x in xrange(LEDS_SIZE) ]
		buf =        [ random.randint(0, 0xff00)   for x in xrange(LEDS_SIZE) ]

		run_test(start_pc, end_pc, code, mem, remainders, buf)



run_tests(100)
