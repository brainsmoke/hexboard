
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

var usedColumns = []int{
	0, 1, 2, 3, 4, 5, 6, 7, // offset
	13, 14,   16, 17,   19, 20,   22, 23,   25, 26,   28, 29,   31, 32,   34, 35, // hex bytes
	41, 42,   44, 45,   47, 48,   50, 51,   53, 54,   56, 57,   59, 60,   62, 63, // hex bytes
	69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, // ascii representations
}

func GetScreenInfo() *ScreenInfo {

	size := 960
	columns := 85
	rows := 18
	height := float64(rows) * digitSize.Y
	width  := float64(columns) * digitSize.X
	digits   := make([]DigitInfo, size)
	segments := make([]SegmentInfo, size*16)
	indices := make([]int, rows*columns)

	for i := range indices {
		indices[i] = -1
	}

	var row, column int
	ix := 0
	for panel := 0; panel < 30; panel++ {
		for module := 0; module < 16; module++ {
			for digit := 0; digit < 2; digit++ {
				if panel < 28 {
					row = module+2
					column = usedColumns[panel*2+digit]
				} else {
					row = 0
					column = 32*(panel-28)+module*2+digit
				}
				digits[ix].Column = column
				digits[ix].Row    = row
				indices[row*columns + column] = ix
				ix++
			}
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
	return Vector2 { X: (float64(column)+.5) * digitSize.X + digitLocation.X,
	                 Y: (float64(row)+.5)    * digitSize.Y + digitLocation.Y }
}

var screenInfo *ScreenInfo

func init() {
	screenInfo = GetScreenInfo()
}
