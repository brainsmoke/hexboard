
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
	SetStyleAtIndex(s Style, index int)

	SetDigit(g font.Glyph, s Style, index int)

	WriteRawAt(g []font.Glyph, column, row int)
	WriteRawAtIndex(g []font.Glyph, index int)

	WriteAt(s string, column, row int)
	WriteAtIndex(s string, index int)

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
		t.SetStyleAtIndex(s, index)
	}
}

func (t *hexScreen) SetStyleAtIndex(s Style, index int) {
	if index >= 0 && index < 960 {
		t.staging[index].style = s.Apply()
	}
	t.tryUpdate()
}

func (t *hexScreen) SetDigit(g font.Glyph, s Style, index int) {
	if index >= 0 && index < 960 {
		t.staging[index].glyph = g
		t.staging[index].style = s.Apply()
	}
	t.tryUpdate()
}

func (t *hexScreen) WriteRawAt(g []font.Glyph, column, row int) {
	index := screenInfo.GetIndex(column, row)
	if index != -1 {
		t.WriteRawAtIndex(g, index)
	}
}

func (t *hexScreen) WriteRawAtIndex(g []font.Glyph, index int) {
	s := t.style.Apply()
	for i :=range(g) {
		if i+index >= 0 && i+index < 960 {
			t.staging[i+index].glyph = g[i]
			t.staging[i+index].style = s
		}
	}
	t.tryUpdate()
}


func (t *hexScreen) WriteRawTitle(g []font.Glyph, start int) {
	l := 64-start
	if l > len(g) {
		l = len(g)
	}
	t.WriteRawAtIndex(g[:l], start)
}

func (t *hexScreen) WriteRawHexField(g []font.Glyph, field int) {
	twodigits := make([]font.Glyph, 2)
	copy(twodigits, g)
	index := 64 + 8 + 2*(field%16) + (field/16)*56
	t.WriteRawAtIndex(twodigits, index)
}

func (t *hexScreen) WriteRawAsciiField(g font.Glyph, field int) {
	index := 64 + 8 + 32 + field%16 + (field/16)*56
	t.SetDigit(g, t.style, index)
}

func (t *hexScreen) WriteRawOffset(g []font.Glyph, line int) {
	offsetdigits := make([]font.Glyph, 8)
	copy(offsetdigits, g)
	index := 64 + (line)*56
	t.WriteRawAtIndex(offsetdigits, index)
}


func (t *hexScreen) WriteAt(s string, column, row int) {
	index := screenInfo.GetIndex(column, row)
	if index != -1 {
		t.WriteAtIndex(s, index)
	}
}

func (t *hexScreen) WriteAtIndex(s string, index int) {
	t.WriteRawAtIndex(t.font.Glyphs(s), index)
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

