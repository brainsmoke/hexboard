
package screen

type FrameBuffer struct {

	digits [][]float64
	frame []float64
}

func NewFrameBuffer() *FrameBuffer {
	f := &FrameBuffer{
		digits: make([][]float64, 960),
		frame: make([]float64, 960*16),
	}
	for i := range f.digits {
		f.digits[i] = f.frame[i*16:(i+1)*16]
	}
	return f
}

