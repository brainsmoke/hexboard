
package screen

import (
	"sync"
	"post6.net/gohexdump/internal/font"
)

type HexScreen interface {
	Screen

	/* locking model assumes one updater, one renderer */
	Hold()
	Update()

	SetFont(font *font.Font)
	Font() *font.Font

	Info() *ScreenInfo

	SetStyle(s Style)
	SetStyleAt(s Style, column, row int)

	UpWrap(column, row int) (int, int)
	DownWrap(column, row int) (int, int)
	LeftWrap(column, row int) (int, int)
	RightWrap(column, row int) (int, int)

	WriteRawAt(g []font.Glyph, column, row int)
	WriteAt(s string, column, row int)

	WriteRawTitle(g []font.Glyph, start int)
	WriteRawHexField(g []font.Glyph, field int)
	WriteRawAsciiField(g font.Glyph, field int)
	WriteRawOffset(g []font.Glyph, line int)


	WriteTitle(s string, start int)
	WriteHexField(s string, field int)
	WriteAsciiField(s string, field int)
	WriteOffset(s string, line int)

}

type digit struct {
	glyph font.Glyph
	style Style
}

type hexScreen struct {

	digits []digit
	staging[]digit

	style Style
	font *font.Font
	held bool
	mutex sync.Mutex
	info *ScreenInfo
}

var hexConfig = ScreenConfiguration{

	{  0, 2, VerticalPanel }, // offset
	{  2, 2, VerticalPanel },
	{  4, 2, VerticalPanel },
	{  6, 2, VerticalPanel },

	{ 13, 2, VerticalPanel }, // hex bytes
	{ 16, 2, VerticalPanel },
	{ 19, 2, VerticalPanel },
	{ 22, 2, VerticalPanel },
	{ 25, 2, VerticalPanel },
	{ 28, 2, VerticalPanel },
	{ 31, 2, VerticalPanel },
	{ 34, 2, VerticalPanel },

	{ 41, 2, VerticalPanel },
	{ 44, 2, VerticalPanel },
	{ 47, 2, VerticalPanel },
	{ 50, 2, VerticalPanel },
	{ 53, 2, VerticalPanel },
	{ 56, 2, VerticalPanel },
	{ 59, 2, VerticalPanel },
	{ 62, 2, VerticalPanel },

	{ 69, 2, VerticalPanel }, // ascii representations
	{ 71, 2, VerticalPanel },
	{ 73, 2, VerticalPanel },
	{ 75, 2, VerticalPanel },
	{ 77, 2, VerticalPanel },
	{ 79, 2, VerticalPanel },
	{ 81, 2, VerticalPanel },
	{ 83, 2, VerticalPanel },

	{  0, 0, HorizontalPanel },
	{ 32, 0, HorizontalPanel },

}

const hexStartRow = 2
var hexColumns = []int{13,16,19,22,25,28,31,34,41,44,47,50,53,56,59,62}
const offsetColumn = 0
const asciiColumn = 69


func NewHexScreen() HexScreen {

	info := GetScreenInfo(hexConfig)
	digits := make([]digit, info.Size)
	staging := make([]digit, info.Size)
	return &hexScreen{ style: defaultStyle, font: font.GetFont(), digits: digits, staging: staging, info: info }
}

func (t *hexScreen) NextFrame(f *FrameBuffer, old *FrameBuffer, tick uint64) bool {

	t.mutex.Lock()
	for i := range t.digits {
		s := t.digits[i].style
		if s == nil {
			s = defaultStyle
		}
		s.Render(f.digits[i], t.digits[i].glyph, i, tick)
	}
	t.mutex.Unlock()

	return true
}

func (t *hexScreen) Info() *ScreenInfo {
	return t.info
}

func (t *hexScreen) Hold() {
	t.held = true
}

func (t *hexScreen) Update() {
	t.held = false

	t.mutex.Lock()
	for i := range t.staging {
		style, glyph := t.staging[i].style, t.staging[i].glyph
		if style != nil {
			t.digits[i].style = style
			t.digits[i].glyph = glyph
		}
		t.staging[i].style = nil
	}
	t.mutex.Unlock()
}

func (t *hexScreen) tryUpdate() {
	if !t.held {
		t.Update()
	}
}

func (t *hexScreen) SetFont(font *font.Font) {
	t.font = font
}

func (t *hexScreen) Font() *font.Font {
	return t.font
}

func (t *hexScreen) SetStyle(s Style) {
	t.style = s
}

func (t *hexScreen) SetStyleAt(s Style, column, row int) {
	index := t.info.GetIndex(column, row)
	if index != -1 {
		t.staging[index].style = s.Apply()
		t.tryUpdate()
	}
}

func (t *hexScreen) WriteRawAt(g []font.Glyph, column, row int) {

	if column < 0 || row < 0 {
		return
	}

	i := 0
	loop: for ; row < t.info.Rows ; row++ {
		for ; column < t.info.Columns ; column++ {
			if i >= len(g) {
				break loop
			}
			index := t.info.GetIndex(column, row)

			if index != -1 {
				s := t.style.Apply()
				t.staging[index].glyph = g[i]
				t.staging[index].style = s
				i++
			}
		}
	}
	t.tryUpdate()
}

func (t *hexScreen) WriteRawTitle(g []font.Glyph, start int) {
	if start < 0 || start >= 64 {
		return
	}
	l := 64-start
	if l > len(g) {
		l = len(g)
	}
	t.WriteRawAt(g[:l], start, 0)
}

func (t *hexScreen) WriteRawHexField(g []font.Glyph, field int) {
	twodigits := make([]font.Glyph, 2)
	copy(twodigits, g)
	t.WriteRawAt(twodigits, hexColumns[field%16], hexStartRow+(field/16))
}

func (t *hexScreen) WriteRawAsciiField(g font.Glyph, field int) {
	g_a := []font.Glyph{ g }
	if 0 <= field && field < 256 {
		t.WriteRawAt(g_a,  asciiColumn+(field%16), hexStartRow+(field/16))
	}
}

func (t *hexScreen) WriteRawOffset(g []font.Glyph, line int) {
	offsetdigits := make([]font.Glyph, 8)
	copy(offsetdigits, g)
	if 0 <= line && line < 16 {
		t.WriteRawAt(offsetdigits, offsetColumn, line+hexStartRow)
	}
}


func (t *hexScreen) WriteAt(s string, column, row int) {
	t.WriteRawAt(t.font.Glyphs(s), column, row)
}

func (t *hexScreen) WriteTitle(s string, start int) {
	t.WriteRawTitle(t.font.Glyphs(s), start)
}

func (t *hexScreen) WriteHexField(s string, field int) {

	t.WriteRawHexField(t.font.Glyphs(s), field)
}

func (t *hexScreen) WriteAsciiField(s string, field int) {
	t.WriteRawAsciiField(t.font.Glyphs(s)[0], field)
}

func (t *hexScreen) WriteOffset(s string, line int) {
	t.WriteRawOffset(t.font.Glyphs(s), line)
}


func (t *hexScreen) UpWrap(column, row int) (int, int) {

	for index := -1; index == -1; index = t.info.GetIndex(column, row) {
		row -= 1
		if row < 0 {
			row = t.info.Rows-1
		}
	}
	return column, row
}


func (t *hexScreen) DownWrap(column, row int) (int, int) {

	for index := -1; index == -1; index = t.info.GetIndex(column, row) {
		row += 1
		if row >= t.info.Rows {
			row = 0
		}
	}
	return column, row
}


func (t *hexScreen) LeftWrap(column, row int) (int, int) {

	for index := -1; index == -1; index = t.info.GetIndex(column, row) {
		column -= 1
		if column < 0 {
			column = t.info.Columns-1
		}
	}
	return column, row
}


func (t *hexScreen) RightWrap(column, row int) (int, int) {

	for index := -1; index == -1; index = t.info.GetIndex(column, row) {
		column += 1
		if column >= t.info.Columns {
			column = 0
		}
	}
	return column, row
}

