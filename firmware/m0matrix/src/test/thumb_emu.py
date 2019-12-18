

def get_state(pc, end_pc, code, mem, buf, gpio):

	return {
		'r0': buf,
		'r1': gpio,
		'r2': "BAD",
		'r3': "BAD",
		'r4': "BAD",
		'r5': "BAD",
		'r6': "BAD",
		'r7': "BAD",
		'r8': "BAD",
		'r9': "BAD",
		'r10': "BAD",
		'fp': "BAD",
#		'r12': "BAD",
		'ip': "BAD",
		'sl': "BAD",
		't': 0,
		'c': "BAD",
		'z': "BAD",
		'out' : [],
		'mem' : mem,
		'code' : code,
		'pc' : pc,
		'end_pc' : end_pc,
	}

def from_le(arr):
	n=0
	for i,b in enumerate(arr):
		n+= b<<(8*i)
	return n

def to_le(n, size):
	arr = [0]*size
	for i in xrange(size):
		arr[i] = (n>>(i*8))&0xff
	return arr

def from_le_array(arr, elem_size):
	bytearr = [0]*(len(arr)*elem_size)
	for i,e in enumerate(arr):
		for j in xrange(elem_size):
			bytearr[i*elem_size+j] = (e>>(j*8))&0xff
	return bytearr

def to_le_array(bytearr, elem_size):
	arr = [0]*(len(bytearr)/elem_size)
	for i,e in enumerate(bytearr):
		j, shift = int(i/elem_size), (i%elem_size)*8
		arr[j] |= (e&0xff)<<shift
	return arr

def read_mem(mem, addr, size):
	arr = [0]*size
	for i in range(size):
		if (addr+i) not in mem:
			print mem
			print addr+i
			raise "meh."
		arr[i] = mem[addr+i]
	return arr

def write_mem(mem, addr, arr, ro=False):
	if not ro and not (0x20000000 <= addr <= addr+len(arr) <= 0x20002000):
		raise "meh."
	for i in range(len(arr)):
		mem[addr+i] = arr[i]

def write_word(mem, addr, word, ro=False):
	if addr & 3:
		raise "meh."
	write_mem(mem, addr, to_le(word, 4), ro)

def read_word(mem, addr):
	if addr & 3:
		raise "meh."
	return from_le(read_mem(mem, addr, 4))


def write_short(mem, addr, short, ro=False):
	if addr & 1:
		raise "meh."
	write_mem(mem, addr, to_le(short, 2), ro)

def read_short(mem, addr):
	if addr & 1:
		raise "meh."
	return from_le(read_mem(mem, addr, 2))

def write_byte(mem, addr, b, ro=False):
	write_mem(mem, addr, to_le(b, 1), ro)

def read_byte(mem, addr):
	return from_le(read_mem(mem, addr, 1))

def parse(line):
	pc = int(line.split(':',1)[0].strip(' '), 16)
	line = line[line.find('   \t')+4:]
	comment = None
	if line.find(';') != -1:
		line, comment = line.split(';',1)
	
	op, rest = (line.split('\t', 1) + [None])[:2]
	if op == 'nop':
		return pc, (op, None, None, None)
	if op in ['b.n', 'bne.n', 'beq.n']:
		dest = int(rest.split(' ', 1)[0].strip(' ][\t#\n'), 16)
		return pc, (op, dest, None, None)

	dest, s1, s2 = (rest.split(',', 2)  + [None])[:3]
	s1 = s1.strip(' ][\t#\n')
	if s2:
		s2 = s2.strip(' ][\t#\n')
	if op == 'ldr' and s1 == 'pc':
		s2 = int(comment[2:].split(' ')[0],16)
	return pc, (op, dest, s1, s2)


def nop(state, dest, s1, s2):
	pass

def bne_n(state, dest, s1, s2):
	if not state['z']:
		state['pc'] = dest-2 # -2 hackmeh.
		state['t'] += 2

def beq_n(state, dest, s1, s2):
	if state['z']:
		state['pc'] = dest-2 # -2 hackmeh.
		state['t'] += 2

def b_n(state, dest, s1, s2):
	if not ( -4 <= (dest-2-state['pc']) <= 2 ): # empirical
		state['t'] += 3
	state['pc'] = dest-2 # -2 hackmeh.


def mov(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	state[dest] = state[s1]

def uxtb(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	state[dest] = state[s1]&0xff
	state['c'] = state['z'] = "bad" # ???

def ldr(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	if s1 == 'pc':
		state[dest] = read_word(state['mem'], int(s2))
	elif s2 in state:
		state[dest] = read_word(state['mem'], state[s1]+state[s2])
	else:
		state[dest] = read_word(state['mem'], state[s1]+int(s2))

def ldrh(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	if s1 == 'pc':
		state[dest] = read_short(state['mem'], int(s2))
	elif s2 in state:
		state[dest] = read_short(state['mem'], state[s1]+state[s2])
	else:
		state[dest] = read_short(state['mem'], state[s1]+int(s2))


def ldrb(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	if s1 == 'pc':
		state[dest] = read_byte(state['mem'], int(s2))
	elif s2 in state:
		state[dest] = read_byte(state['mem'], state[s1]+state[s2])
	else:
		state[dest] = read_byte(state['mem'], state[s1]+int(s2))

def movs(state, dest, s1, s2):
	if dest not in state:
		raise "meh"
	state[dest] = int(s1,0)
	state['c'] = 0
	state['z'] = int(int(s1,0) == 0)


def add(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	if s2 == None:
		state[dest] += state[s1]
	elif s2 not in state:
		state[dest] = state[s1] + int(s2)
	else:
		state[dest] = state[s1] + state[s2]
	state[dest] = state[dest] & 0xffffffff

def adds(state, dest, s1, s2):
	if dest not in state:
		raise "meh"
 	if s1 not in state:
		state[dest] += int(s1,0)
	elif s2 not in state:
		state[dest] = state[s1] + int(s2,0)
	else:
		state[dest] = state[s1] + state[s2]
	state['c'] = (state[dest] >> 32) & 1
	state[dest] = state[dest] & 0xffffffff
	state['z'] = int(state[dest] == 0)

def subs(state, dest, s1, s2):
	if dest not in state:
		raise "meh"
 	if s1 not in state:
		state[dest] -= int(s1,0)
	elif s2 not in state:
		state[dest] = state[s1] - int(s2,0)
	else:
		state[dest] = state[s1] - state[s2]
	state['c'] = (state[dest] >> 32) & 1
	state[dest] = state[dest] & 0xffffffff
	state['z'] = int(state[dest] == 0)


def cmp_(state, dest, s1, s2):
	x=0
	if dest not in state:
		raise "meh"
 	if s1 not in state:
		x = state[dest] - int(s1,0)
	else:
		x = state[dest] - state[s1]
	state['c'] = (x >> 32) & 1
	x = x & 0xffffffff
	state['z'] = int(x == 0)

def adcs(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	state[dest] = state[dest] + state[s1] + state['c']
	state['c'] = (state[dest] >> 32) & 1
	state[dest] = state[dest] & 0xffffffff
	state['z'] = int(state[dest] == 0)

def lsls(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	n = int(s2,0)
	state[dest] = state[s1] << n
	state['c'] = (state[dest] >> 32) & 1
	state[dest] = state[dest] & 0xffffffff
	state['z'] = int(state[dest] == 0)


def lsrs(state, dest, s1, s2):
	if dest not in state or s1 not in state:
		raise "meh"
	n = int(s2,0)
	state['c'] = (state[dest] >> (n-1)) & 1
	state[dest] = state[s1] >> n
	state[dest] = state[dest] & 0xffffffff
	state['z'] = int(state[dest] == 0)

def strh(state, dest, s1, s2):
	if dest not in state:
		raise "meh"
	addr = state[s1]+int(s2)
	if 0x40000000 <= addr <= 0x50000000:
		state['out'] += [ (state['t'], addr, state[dest]&0xffff, state['pc']) ]
	else:
		write_short(state['mem'], addr, state[dest])

def strb(state, dest, s1, s2):
	if dest not in state:
		raise "meh"
	addr = state[s1]+int(s2)
	if 0x40000000 <= addr <= 0x50000000:
		state['out'] += [ (state['t'], addr, state[dest]&0xff, state['pc']) ]
	else:
		write_byte(state['mem'], addr, state[dest])

op_func = {

	'mov': mov,
	'movs': movs,
	'cmp': cmp_,
	'subs': subs, 
	'adds': adds, 
	'add': add, 
	'adcs': adcs, 
	'lsls': lsls, 
	'lsrs': lsrs, 
	'uxtb': uxtb, 
	'strh': strh, 
	'strb': strb, 
	'ldr':ldr,
	'ldrh':ldrh,
	'ldrb':ldrb,
	'bne.n':bne_n,
	'beq.n':beq_n,
	'b.n':b_n,
	'nop':nop,
}

t_op = {

	'mov': 1,
	'movs': 1,
	'cmp': 1, 
	'subs': 1, 
	'adds': 1, 
	'add': 1, 
	'adcs': 1, 
	'lsls': 1, 
	'lsrs': 1, 
	'uxtb': 1, 
	'strh': 2, 
	'strb': 2, 
	'ldr': 2,
	'ldrh': 2,
	'ldrb': 2,
	'bne.n':1, #*
	'beq.n':1, #*
	'b.n':3,
	'nop':1,
}

def run(state):
	while state['pc'] != state['end_pc']:
		op, dest, s1, s2 = state['code'][state['pc']]
		op_func[op](state, dest, s1, s2)
		state['t'] += t_op[op]
		state['pc'] += 2


def read_code(objdump_out_file, start_label, end_label):
	mem = {}
	code = {}
	start_pc = None
	end_pc = None

	output = False
	for l in objdump_out_file:
		if l.find('.word') != -1:
			addr = int(l.split(':',1)[0].strip(' \t'),16)
			write_word(mem, addr, int(l.split('.word',1)[1].strip(' \t'),16), ro=True)
	
		if l.find('.short') != -1:
			addr = int(l.split(':',1)[0].strip(' \t'),16)
			write_short(mem, addr, int(l.split('.short',1)[1].strip(' \t'),16), ro=True)
	
		if l.find('.byte') != -1:
			addr = int(l.split(':',1)[0].strip(' \t'),16)
			write_byte(mem, addr, int(l.split('.byte',1)[1].strip(' \t'),16), ro=True)

		if l.endswith('<'+start_label+'>:\n'):
			output = True
			start_pc = int(l.split(' ')[0], 16)

		if l.endswith('<'+end_label+'>:\n'):
			output = False
			end_pc = int(l.split(' ')[0], 16)

		if not output or l[0] != ' ':
			continue

		pc, inst = parse(l)
		code[pc] = inst

	return start_pc, end_pc, code, mem

