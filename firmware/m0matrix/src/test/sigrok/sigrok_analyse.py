#!/usr/bin/env python2
#sigrok-cli -O hex:width=4096 -C 0 --config samplerate=16M --driver=fx2lafw --time=100s |sed 's:^..::'|tr -d ' '
#sigrok-cli -O hex:width=4096 -C 0 --config samplerate=16M --driver=fx2lafw --time=100s |./sigrok_analyse.py


import sys

high=False
t=0

abs_t=0
started = False
h = False
t_hi=0

freq = 16e6

t0h = 0.4e-6
t1h = 0.8e-6
t0l = 0.85e-6
t1l = 0.45e-6
err = 150e-9
timeout = 50e-6

t0h_min, t0h_max = int( (t0h-err)*freq ), int( (t0h+err)*freq )
t1h_min, t1h_max = int( (t1h-err)*freq ), int( (t1h+err)*freq )
t0l_min, t0l_max = int( (t0l-err)*freq ), int( (t0l+err)*freq )
t1l_min, t1l_max = int( (t1l-err)*freq ), int( (t1l+err)*freq )

timeout = int(timeout*freq)

bit=0x80
byte = 0
f=[]

def process(t, state):
	global started, abs_t, t_hi, h, bit, byte,f

	if started:
		if state == True:
			t_hi = t
		else:
			if t0h_min <= t_hi <= t0h_max and t0l_min <= t:
				if t0l_max < t < timeout:
					print (t_hi, t)
					sys.exit(1)
				bit >>= 1
			elif t1h_min <= t_hi <= t1h_max and t1l_min <= t:
				if t1l_max < t < timeout:
					print (t_hi, t)
					sys.exit(1)
				byte |= bit
				bit >>= 1
			else:
				print (t_hi, t)
				sys.exit(1)

			if bit == 0:
				bit = 0x80
				f.append(chr(byte).encode('hex'))
				byte = 0
		
		if state == h:
			print (state,'=', h)
			sys.exit(1)
		h = not h

	if state == False and t > timeout:
		if started:
			print ('[' + ' '.join(f) + ']')
		started = True
		h = False
		bit = 0x80
		byte = 0
		f = []

sys.stdin.readline()
sys.stdin.readline()

for line in sys.stdin:
	for c in line[2:-1]:
		if high:
			if c == 'f':
				t+=4
			elif c == ' ':
				pass
			elif c == 'e':
				process(t+3, True)
				high = False
				t = 1
			elif c == 'c':
				process(t+2, True)
				high = False
				t = 2
			elif c == '8':
				process(t+1, True)
				high = False
				t = 3
			elif c == '0':
				process(t, True)
				high = False
				t = 4
			elif c == 'd':
				process(t+2, True)
				process(1, False)
				t = 1
			elif c == '9':
				process(t+1, True)
				process(2, False)
				t = 1
			elif c == 'b':
				process(t+1, True)
				process(1, False)
				t = 2
			elif c == 'a':
				process(t+1, True)
				process(1, False)
				process(1, True)
				high = False
				t = 1
			else:
				process(t, True)
				if c == '1':
					process(3, False)
					t = 1
				elif c == '3':
					process(2, False)
					t = 2
				elif c == '7':
					process(1, False)
					t = 3
				elif c == '4':
					process(1, False)
					process(1, True)
					t = 2
					high = False
				elif c == '6':
					process(1, False)
					process(2, True)
					t = 1
					high = False
				elif c == '2':
					process(2, False)
					process(1, True)
					t = 1
					high = False
				elif c == '5':
					process(1, False)
					process(1, True)
					process(1, False)
					t = 1
				elif c not in ' \n':
					print (c)
					sys.exit(1)
		else:
			if c == '0':
				t+=4
			elif c == ' ':
				pass
			elif c == '1':
				process(t+3, False)
				high = True
				t = 1
			elif c == '3':
				process(t+2, False)
				high = True
				t = 2
			elif c == '7':
				process(t+1, False)
				high = True
				t = 3
			elif c == 'f':
				process(t, False)
				high = True
				t = 4
			elif c == '2':
				process(t+2, False)
				process(1, True)
				t = 1
			elif c == '6':
				process(t+1, False)
				process(2, True)
				t = 1
			elif c == '4':
				process(t+1, False)
				process(1, True)
				t = 2
			elif c == '5':
				process(t+1, False)
				process(1, True)
				process(1, False)
				high = True
				t = 1
			else:
				process(t, False)
				if c == 'e':
					process(3, True)
					t = 1
				elif c == 'c':
					process(2, True)
					t = 2
				elif c == '8':
					process(1, True)
					t = 3
				elif c == 'b':
					process(1, True)
					process(1, False)
					t = 2
					high = True
				elif c == '9':
					process(1, True)
					process(2, False)
					t = 1
					high = True
				elif c == 'd':
					process(2, True)
					process(1, False)
					t = 1
					high = True
				elif c == 'a':
					process(1, True)
					process(1, False)
					process(1, True)
					t = 1
				elif c not in ' \n':
					print (c)
					sys.exit(1)
