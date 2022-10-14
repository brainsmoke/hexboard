
package screen

type FrameBuffer struct {

	digits [][]float64
	frame []float64
}

func NewFrameBuffer(size int) *FrameBuffer {
	f := &FrameBuffer{
		digits: make([][]float64, size),
		frame: make([]float64, size*16),
	}
	for i := range f.digits {
		f.digits[i] = f.frame[i*16:(i+1)*16]
	}
	return f
}

func (f *FrameBuffer) Clear() {
	frame := f.frame
	for i := range frame {
		frame[i] = 0.0
	}
}
