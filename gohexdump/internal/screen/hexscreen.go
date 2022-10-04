
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

type hexScreen struct {

	digits [960]struct {
		glyph font.Glyph
		style Style
	}
	staging[960]struct {
		glyph font.Glyph
		style Style
	}
	style Style
	font *font.Font
	held bool
	mutex sync.Mutex
}

func NewHexScreen() HexScreen {

	return &hexScreen{ style: defaultStyle, font: font.GetFont() }
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
	index := screenInfo.GetIndex(column, row)
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
	loop: for ; row < screenInfo.Rows ; row++ {
		for ; column < screenInfo.Columns ; column++ {
			if i >= len(g) {
				break loop
			}
			index := screenInfo.GetIndex(column, row)

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
	t.WriteRawAt(twodigits, usedColumns[8+2*(field%16)], 2+(field/16))
}

func (t *hexScreen) WriteRawAsciiField(g font.Glyph, field int) {
	g_a := []font.Glyph{ g }
	if 0 <= field && field < 256 {
		t.WriteRawAt(g_a,  usedColumns[40+(field%16)], 2+(field/16))
	}
}

func (t *hexScreen) WriteRawOffset(g []font.Glyph, line int) {
	offsetdigits := make([]font.Glyph, 8)
	copy(offsetdigits, g)
	t.WriteRawAt(offsetdigits, 0, line+2)
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

	row -= 1
	for index := -1; index == -1; index = screenInfo.GetIndex(column, row) {
		row -= 1
		if row < 0 {
			row = screenInfo.Rows-1
		}
	}
	return column, row
}


func (t *hexScreen) DownWrap(column, row int) (int, int) {

	row += 1
	for index := -1; index == -1; index = screenInfo.GetIndex(column, row) {
		row -= 1
		if row >= screenInfo.Rows {
			row = 0
		}
	}
	return column, row
}


func (t *hexScreen) LeftWrap(column, row int) (int, int) {

	column -= 1
	for index := -1; index == -1; index = screenInfo.GetIndex(column, row) {
		column -= 1
		if column < 0 {
			column = screenInfo.Columns-1
		}
	}
	return column, row
}


func (t *hexScreen) RightWrap(column, row int) (int, int) {

	column += 1
	for index := -1; index == -1; index = screenInfo.GetIndex(column, row) {
		column += 1
		if column >= screenInfo.Columns {
			column = 0
		}
	}
	return column, row
}

