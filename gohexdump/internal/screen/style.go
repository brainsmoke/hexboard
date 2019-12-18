
package screen

import (
	"math"
	"time"
	"post6.net/gohexdump/internal/font"
	"post6.net/gohexdump/internal/demomath/wave"
)

type Style interface {

	Apply() Style /* allow for per-use memory if needed */
	Render(dst []float64, glyph font.Glyph, frameIndex int, tick uint64)
}

type SimpleStyle struct {

	fg, bg float64
}

var defaultStyle = &SimpleStyle{ .3, 0 }

func NewBrightness( brightness float64 ) Style {
	return &SimpleStyle{ fg: brightness, bg: 0 }
}

func (s *SimpleStyle) Render(dst []float64, glyph font.Glyph, frameIndex int, tick uint64) {

	for i := range dst {
		if glyph & (1<<uint(i)) != 0 {
			dst[i] = s.fg
		} else {
			dst[i] = s.bg
		}
	}
}

func (s *SimpleStyle) Apply() Style {
	return s
}

type PeriodicStyle struct {

	wave []float64
	fgBase, fgAmp, bgBase, bgAmp float64
	fgPhase, bgPhase int
	multiplier uint64
}

func (s *PeriodicStyle) Apply() Style {
	return s
}

func (s *PeriodicStyle) Render(dst []float64, glyph font.Glyph, frameIndex int, tick uint64) {
	ix := int( (tick * s.multiplier)>>wave.Shift )
	fg := s.fgBase + s.wave[ (ix + s.fgPhase) & wave.Mask ]*s.fgAmp
	bg := s.bgBase + s.wave[ (ix + s.bgPhase) & wave.Mask ]*s.bgAmp

	for i := range dst {
		if glyph & (1<<uint(i)) != 0 {
			dst[i] = fg
		} else {
			dst[i] = bg
		}
	}
}

func NewBounce(min, max float64, period time.Duration) Style {
	min, max = math.Max(0, min), math.Min(1., max)
	m := uint64(wave.Multiplier) / uint64(period)
	return &PeriodicStyle{ wave:wave.Wave, fgBase: min, fgAmp: max-min, multiplier: m}
}

