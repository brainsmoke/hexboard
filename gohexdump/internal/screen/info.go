
package screen

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


type DigitInfo struct {

	Column, Row int
	Position Vector2
	Segments []SegmentInfo
}

type SegmentInfo struct {

	Position Vector2
	Digit int
	Segment int
}


type ScreenInfo struct {

	Size, Columns, Rows int // Size != Columns * Rows, since there are gaps
	Height, Width float64   // in mm
	Digits []DigitInfo
	Segments []SegmentInfo
	indices []int
}

type screenPos struct { column, row int }

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

type ScreenConfiguration []PanelPosition


func GetScreenInfo(conf ScreenConfiguration) *ScreenInfo {

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

	height := float64(rows) * digitSize.Y
	width  := float64(columns) * digitSize.X
	digits   := make([]DigitInfo, size)
	segments := make([]SegmentInfo, size*16)
	indices := make([]int, rows*columns)

	for i := range indices {
		indices[i] = -1
	}

	ix := 0
	for _, panel := range conf {
		for _, pos := range panel.Type.digitPositions {
			x, y := panel.Column + pos.column, panel.Row + pos.row

			digits[ix].Column = x
			digits[ix].Row    = y
			if indices[y*columns + x] != -1 {
				panic("panel overlap")
			}

			indices[y*columns + x] = ix
			ix += 1
		}
	}

	for i := range digits {

		digits[i].Position = Vector2{ X:float64(digits[i].Column) * digitSize.X + digitLocation.X,
		                              Y:float64(digits[i].Row)    * digitSize.Y + digitLocation.Y }

		s := segments[i*16:i*16+16]

		for j := range s {
			s[j].Position = Vector2{ X:float64(digits[i].Column) * digitSize.X + segmentLocations[j].X,
			                         Y:float64(digits[i].Row)    * digitSize.Y + segmentLocations[j].Y }
			s[j].Digit = i
			s[j].Segment = j
		}

		digits[i].Segments = s
	}

	return &ScreenInfo{
		Size: size,
		Columns: columns, Rows: rows,
		Width: width, Height: height,
		Digits: digits,
		Segments: segments,
		indices: indices,
	}
}


func (s *ScreenInfo) GetIndex(column, row int) int {
	if row < 0 || row >= s.Rows || column < 0 || column >= s.Columns {
		return -1
	}
	return s.indices[row*s.Columns + column]
}

func (s *ScreenInfo) GetCoord(column, row int) Vector2 {
	return Vector2 { X: float64(column) * digitSize.X + digitLocation.X,
	                 Y: float64(row)    * digitSize.Y + digitLocation.Y }
}

