
package screen

import (
	"math"
)

type Filter interface {

	Render(f *FrameBuffer, old *FrameBuffer, tick uint64)
}

type afterGlowFilter struct {
	factor float64
}

func NewAfterGlowFilter(factor float64) Filter {
	return &afterGlowFilter{factor: math.Min(math.Max(0, factor), 1)}
}

func (s *afterGlowFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for i := range old.frame {
		fade := old.frame[i] * s.factor
		if fade > f.frame[i] {
			f.frame[i] = fade
		}
	}
}
