package font

import (
	"strings"
)

type Glyph uint16

type Font struct {

	page []Glyph
	highGlyphs map[uint]Glyph
}

const segments = "ABCDEFxyHJKLMNz"

func parseGlyph(s string) Glyph {

	var glyph Glyph
	s = strings.ToUpper(s)
	s = strings.Replace(s, "G1", "x", -1)
	s = strings.Replace(s, "G2", "y", -1)
	s = strings.Replace(s, "G", "xy", -1)
	s = strings.Replace(s, "DP", "z", -1)
	for _,c := range []rune(s) {
		i := strings.IndexRune(segments, c)
		if i >= 0 {
			glyph |= 1 << uint(i)
		}
	}
	return glyph
}

func ParseFont(data string) *Font {
	var f = Font { page: make([]Glyph, 128) }

	lines := strings.Split(data, "\n")
	for i := range lines {
		segments := string([]rune(lines[i])[2:])
		g := parseGlyph(segments)
		c := uint([]rune(lines[i])[0])
		if c < uint(len(f.page)) {
			f.page[c] = g
		} else {
			if f.highGlyphs == nil {
				f.highGlyphs = make(map[uint]Glyph)
			}
			f.highGlyphs[c] = g
		}
	}

	return &f
}

func (f *Font) GetGlyph(c rune) Glyph {
	i := uint(c)
	if i < uint(len(f.page)) {
		return f.page[i]
	} else {
		return f.highGlyphs[i]
	}
}

func (f *Font) Glyphs(s string) []Glyph {
	r := []rune(s)
	g := make([]Glyph, len(r))
	for i,c := range r {
		g[i] = f.GetGlyph(c)
	}
	return g
}

func GetFont() *Font {
	return ParseFont(fontData)
}

const fontData = `0:abcdefkn
1:kbc
2:abged
3:abgcd
4:fgbc
5:afgcd
6:afgcde
7:akm
8:abcdefg
9:abcdfg
@:eg1medcba
A:efabcg
B:abcdjmg2
C:adef
D:abcdjm
E:afedg1
F:afeg1
G:afedcg2
H:bcefg
I:adjm
J:bcd
K:efg1kl
L:def
M:bckhfe
N:fehlcb
O:abcdef
P:abefg
Q:abcdefl
R:efabgl
S:afg1ld
T:ajm
U:bcdef
V:knef
W:bclnef
X:hkln
Y:hkm
Z:aknd
":fb
#:bcefjmg
$:afgcdjm
%:fknc
&:afedlg1
':j
(:klbcg2
):hnefg1
*:hjknmlg
+:gjm
,:n
.:Dp
-:g
/:nk
\:hl
[:defaDp
]:abcd
_:d
{:dng1ha
}:dlg2ka
|:jm
^:kb
<:klg1
>:hng2
=:gd
?:abg2m
~:fhg2b
`+"`"+`:h
a:nkbg2c
b:feg1dl
c:ged
d:ndg2cb
e:g1end
f:efag1
g:kbg2cd
h:efgc
i:m
j:cd
k:jkml
l:bcdp
m:egcm
n:mlc
o:cdeg
p:afkg1e
q:ahbg2c
r:eg1
s:g2ld
t:gm
u:cde
v:en
w:enlc
x:gnl
y:hkn
z:g1nd
;:jn
::hn
!:bdp`

