
package screen

import (
	"math"
	"post6.net/gohexdump/internal/util/clip"
)

type gammaFilter struct {
	table [257]float64
}

func DefaultGamma() Filter {
	return NewGammaFilter(2.5, 1.)
}

func NewGammaFilter(gamma, brightness float64) Filter {

	brightness = clip.FloatBetween(brightness, 0, 1)

	g := &gammaFilter{ }
	factor := brightness / math.Pow(255, gamma)
	for i := 0; i < 256; i++ {
		g.table[i] = factor * math.Pow(float64(i), gamma)
	}
	g.table[256] = g.table[255]
	return g
}

func (g *gammaFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for i := range f.frame {
		v := clip.FloatBetween(f.frame[i] * 255, 0, 255)
        v_int := math.Floor(v)
		v_rem := v - v_int
		lo, hi := g.table[int(v_int)], g.table[int(v_int)+1]
		f.frame[i] = lo + v_rem*(hi-lo)
	}
}
