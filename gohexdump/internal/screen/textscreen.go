
package screen

import (
	"sync"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/util/clip"
	"errors"
)

type Vector2 struct {

	X, Y float64
}

var digitLocation = Vector2{ 6.333, 14.062 }
var digitSize = Vector2{ 2.54*5, 2.54*11 }

var segmentLocations = []Vector2{

	{  6.8266127,  4.399   }, // A
	{  10.028027,  7.44098 }, // B
	{  9.4245537, 13.92735 }, // C
	{  5.6392029, 16.88621 }, // D
	{  2.4377948, 13.84316 }, // E
	{  3.0412692,  7.35735 }, // F
	{  4.4950327, 10.65115 }, // G1
	{  8.0169715, 10.64515 }, // G2
	{  4.7753463,  7.54048 }, // H
	{  6.4952691,  7.87499 }, // J
	{  8.2592078,  7.57657 }, // K
	{  7.6904820, 13.74251 }, // L
	{  5.9705644, 13.40853 }, // M
	{  4.2065789, 13.706   }, // N
	{ 10.8210000, 16.600   }, // Dp
	{  6.3330000, 14.062   }, // unused (center location)
}

type ScreenInfo interface {

	DigitCount() int
	SegmentCount() int

	DigitCoord(ix int) Vector2
	SegmentCoord(ix int) Vector2

	Coords() []Vector2

	Dimensions() Vector2
}

type TextScreen interface {

	Screen
	ScreenInfo

	DigitIndex(column, row int) int
	DigitPosition(ix int) (int, int)

	Size() (int, int)
	Rows() int
	Columns() int

	/* locking model assumes one updater, one renderer */
	Hold()
	Update()

	SetFont(font *font.Font)
	Font() *font.Font

	SetStyle(s Style)
	SetStyleAt(s Style, column, row int)

	WriteRawAt(g []font.Glyph, column, row int) (int, int, error)
	WriteAt(s string, column, row int) (int, int, error)

	First() (int, int)
	Last() (int, int)

	Next(column, row int) (int, int, error)
	Previous(column, row int) (int, int, error)

	NextWrap(column, row int) (int, int)
	PreviousWrap(column, row int) (int, int)

	Up(column, row int) (int, int, error)
	Down(column, row int) (int, int, error)
	Left(column, row int) (int, int, error)
	Right(column, row int) (int, int, error)

	UpWrap(column, row int) (int, int)
	DownWrap(column, row int) (int, int)
	LeftWrap(column, row int) (int, int)
	RightWrap(column, row int) (int, int)

	Scroll(right, down int) /* only makes sense for rectangular screens */
	Clear()
}

type screenPos struct { column, row int }

type digit struct {
	glyph font.Glyph
	style Style
}

type textScreen struct {

	columns, rows int // Size != Columns * Rows, since there are gaps
	positions []screenPos
	indices []int
	digits, staging []digit

	style Style
	font *font.Font
	held bool
	mutex sync.Mutex
}

type Panel struct {

	digitPositions []screenPos
}

var HorizontalPanel = &Panel{
	digitPositions : []screenPos{
		{ 0, 0}, { 1, 0}, { 2, 0}, { 3, 0}, { 4, 0}, { 5, 0}, { 6, 0}, { 7, 0},
		{ 8, 0}, { 9, 0}, {10, 0}, {11, 0}, {12, 0}, {13, 0}, {14, 0}, {15, 0},
		{16, 0}, {17, 0}, {18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0},
		{24, 0}, {25, 0}, {26, 0}, {27, 0}, {28, 0}, {29, 0}, {30, 0}, {31, 0},
	},
}

var VerticalPanel = &Panel{
	digitPositions : []screenPos{
		{ 0, 0}, { 1, 0},
		{ 0, 1}, { 1, 1},
		{ 0, 2}, { 1, 2},
		{ 0, 3}, { 1, 3},
		{ 0, 4}, { 1, 4},
		{ 0, 5}, { 1, 5},
		{ 0, 6}, { 1, 6},
		{ 0, 7}, { 1, 7},
		{ 0, 8}, { 1, 8},
		{ 0, 9}, { 1, 9},
		{ 0,10}, { 1,10},
		{ 0,11}, { 1,11},
		{ 0,12}, { 1,12},
		{ 0,13}, { 1,13},
		{ 0,14}, { 1,14},
		{ 0,15}, { 1,15},
	},
}

type PanelPosition struct {

	Column, Row int
	Type *Panel
}

type Configuration []PanelPosition


func (s *textScreen) init(conf Configuration) {

	columns := 0
	rows := 0
	size := 0

	for _, panel := range conf {
		for _, pos := range panel.Type.digitPositions {
			x, y := panel.Column + pos.column, panel.Row + pos.row
			if x < 0 {
				panic("row pos < 0")
			}
			if y < 0 {
				panic("column pos < 0")
			}
			if x+1 > columns {
				columns = x+1
			}
			if y+1 > rows {
				rows = y+1
			}
			size += 1
		}
	}

	positions := make([]screenPos, size)
	indices := make([]int, rows*columns)

	digits := make([]digit, size)
	staging := make([]digit, size)

	for i := range indices {
		indices[i] = -1
	}

	ix := 0
	for _, panel := range conf {
		for _, pos := range panel.Type.digitPositions {
			x, y := panel.Column + pos.column, panel.Row + pos.row

			positions[ix].column = x
			positions[ix].row    = y
			if indices[y*columns + x] != -1 {
				panic("panel overlap")
			}

			indices[y*columns + x] = ix
			ix += 1
		}
	}

	*s = textScreen{
		columns: columns,
		rows: rows,
		positions: positions,
		indices: indices,
		digits: digits,
		staging: staging,
		font: font.GetFont(),
		style: defaultStyle,
	}
}

func NewTextScreen(conf Configuration) TextScreen {
	s := new(textScreen)
	s.init(conf)
	return s
}

func (s *textScreen) DigitCount() int {
	return len(s.positions)
}

func (s *textScreen) SegmentCount() int {
	return len(s.positions)*16
}

func (s *textScreen) Size() (int, int) {
	return s.columns, s.rows
}

func (s *textScreen) Rows() int {
	return s.rows
}

func (s *textScreen) Columns() int {
	return s.columns
}

func (s *textScreen) Dimensions() Vector2 {
	return Vector2{ float64(s.columns) * digitSize.X, float64(s.rows) * digitSize.Y }
}

func (s *textScreen) DigitIndex(column, row int) int {
	if row < 0 || row >= s.rows || column < 0 || column >= s.columns {
		return -1
	}
	return s.indices[row*s.columns + column]
}

func (s *textScreen) DigitPosition(ix int) (int, int) {
	p := s.positions[ix]
	return p.column, p.row
}

func (s *textScreen) DigitCoord(ix int) Vector2 {
	x, y := s.DigitPosition(ix)
	return Vector2 { X: float64(x) * digitSize.X + digitLocation.X,
	                 Y: float64(y) * digitSize.Y + digitLocation.Y }
}

func (s *textScreen) SegmentCoord(ix int) Vector2 {
	v := s.DigitCoord(ix>>4)
	d := segmentLocations[ix&0xf]
	return Vector2{ v.X+d.X, v.Y+d.Y }
}


func (s *textScreen) Coords() []Vector2 {
	coords := make([]Vector2, len(s.positions))
	for i := range coords {
		coords[i] = s.SegmentCoord(i)
	}

	return coords
}

func (s *textScreen) NextFrame(f, old *FrameBuffer, tick uint64) bool {

	s.mutex.Lock()
	for i := range s.digits {
		style := s.digits[i].style
		if style == nil {
			style = defaultStyle
		}
		style.Render(f.digits[i], s.digits[i].glyph, i, tick)
	}
	s.mutex.Unlock()

	return true
}


func (s *textScreen) Hold() {
	s.held = true
}

func (s *textScreen) Update() {
	s.held = false

	s.mutex.Lock()
	for i := range s.staging {
		style, glyph := s.staging[i].style, s.staging[i].glyph
//		if style != nil {
			s.digits[i].style = style
			s.digits[i].glyph = glyph
//		}
//		s.staging[i].style = nil
	}
	s.mutex.Unlock()
}

func (s *textScreen) tryUpdate() {
	if !s.held {
		s.Update()
	}
}

func (s *textScreen) SetFont(font *font.Font) {
	s.font = font
}

func (s *textScreen) Font() *font.Font {
	return s.font
}

func (s *textScreen) SetStyle(style Style) {
	s.style = style
}

func (s *textScreen) SetStyleAt(style Style, column, row int) {
	index := s.DigitIndex(column, row)
	if index != -1 {
		s.staging[index].style = style.Apply()
		s.tryUpdate()
	}
}

func (s *textScreen) WriteAt(str string, column, row int) (int, int, error) {
	return s.WriteRawAt(s.font.Glyphs(str), column, row)
}

func (s *textScreen) WriteRawAt(g []font.Glyph, column, row int) (int, int, error) {

	x, y, err := s.Next(column-1, row)

	for _, glyph := range g {

		if err != nil {
			break
		}
		index := s.DigitIndex(column, row)
		s.staging[index].glyph = glyph
		s.staging[index].style = s.style.Apply()
		x, y, err = s.Next(x, y)
	}
	s.tryUpdate()
	return x, y, err
}

func (s * textScreen) Scroll(right, down int) {

	if right == 0 && down == 0 {
		return
	}

	startx, starty, endx, endy, dx, dy := 0, 0, s.columns, s.rows, 1, 1

	if right < 0 {
		startx, endx, dx = endx-1, -1, -1
	}
	if down < 0 {
		starty, endy, dy = endy-1, -1, -1
	}

	var empty font.Glyph

	for y := starty; y < endy; y += dy {
		for x := startx; x < endx; x += dx {
			destIndex := s.DigitIndex(x, y)
			if destIndex == -1 {
				continue
			}
			srcIndex := s.DigitIndex(x+right, y+down)
			g, style := empty, Style(nil)
			if srcIndex != -1 {
				g = s.staging[srcIndex].glyph
				style = s.staging[srcIndex].style
			}

			s.staging[destIndex].glyph = g
			s.staging[destIndex].style = style
		}
	}

	s.tryUpdate()
}

func (s * textScreen) Clear() {

	var empty font.Glyph
	for i := range s.staging {
		s.staging[i].glyph = empty
		s.staging[i].style = s.style
	}
}

var noSuchField = errors.New("no such field")

func (s *textScreen) First() (int, int) {

	x,y,err := s.Next(-1, 0)
	if err == nil {
		return x, y
	}
	panic(err)
}

func (s *textScreen) Last() (int, int) {

	x, y, err := s.Next(-1, 0)
	if err == nil {
		return x, y
	}
	panic(err)
}

func (s *textScreen) Next(column, row int) (int, int, error) {

	xstart := clip.IntMax(column+1, 0)
	ystart := clip.IntMax(row, 0)

	for y := ystart ; y < s.rows ; y+=1 {
		for x := xstart ; x < s.columns ; x+=1 {
			if s.indices[y*s.columns + x] != -1 {
				return x, y, nil
			}
		}
		xstart = 0
	}
	return column, row, noSuchField
}

func (s *textScreen) Previous(column, row int) (int, int, error) {

	xstart := clip.IntMin(column-1, s.columns-1)
	ystart := clip.IntMin(row, s.rows-1)

	for y := ystart ; y >= 0 ; y-=1 {
		for x := xstart ; x >= 0 ; x-=1 {
			if s.indices[y*s.columns + x] != -1 {
				return x, y, nil
			}
		}
		xstart = s.columns-1
	}
	return column, row, noSuchField
}


func (s *textScreen) NextWrap(column, row int) (int, int) {
	x, y, err := s.Next(column, row)
	if err == nil {
		return x, y
	}
	return s.First()
}

func (s *textScreen) PreviousWrap(column, row int) (int, int) {
	x, y, err := s.Previous(column, row)
	if err == nil {
		return x, y
	}
	return s.Last()
}


func (s *textScreen) Up(column, row int) (int, int, error) {

	if column >= 0 && column < s.columns {
		for y := clip.IntMin(row-1, s.rows-1) ; y >= 0 ; y-=1 {
			if s.indices[y*s.columns + column] != -1 {
				return column, y, nil
			}
		}
	}
	return column, row, noSuchField
}

func (s *textScreen) Down(column, row int) (int, int, error) {
	if column >= 0 && column < s.columns {
		for y := clip.IntMax(row+1, 0) ; y < s.rows ; y+=1 {
			if s.indices[y*s.columns + column] != -1 {
				return column, y, nil
			}
		}
	}
	return column, row, noSuchField
}

func (s *textScreen) Left(column, row int) (int, int, error) {

	if row >= 0 && row < s.rows {
		for x := clip.IntMin(column-1, s.columns-1) ; x >= 0 ; x-=1 {
			if s.indices[row*s.columns + x] != -1 {
				return row, x, nil
			}
		}
	}
	return column, row, noSuchField
}

func (s *textScreen) Right(column, row int) (int, int, error) {

	if row >= 0 && row < s.rows {
		for x := clip.IntMax(column+1, 0) ; x < s.columns ; x+=1 {
			if s.indices[row*s.columns + x] != -1 {
				return row, x, nil
			}
		}
	}
	return column, row, noSuchField
}


func (s *textScreen) UpWrap(column, row int) (int, int) {

	x, y, err := s.Up(column, row)
	if err != nil {
		x, y, err = s.Up(column, s.rows)
		if err != nil {
			panic(err)
		}
	}
	return x, y
}

func (s *textScreen) DownWrap(column, row int) (int, int) {

	x, y, err := s.Down(column, row)
	if err != nil {
		x, y, err = s.Down(column, -1)
		if err != nil {
			panic(err)
		}
	}
	return x, y
}

func (s *textScreen) LeftWrap(column, row int) (int, int) {

	x, y, err := s.Left(column, row)
	if err != nil {
		x, y, err = s.Left(s.columns, row)
		if err != nil {
			panic(err)
		}
	}
	return x, y
}

func (s *textScreen) RightWrap(column, row int) (int, int) {

	x, y, err := s.Right(column, row)
	if err != nil {
		x, y, err = s.Right(-1, row)
		if err != nil {
			panic(err)
		}
	}
	return x, y
}

