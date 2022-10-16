
package screen

import (
	"sync"
)

type Cursor interface {
	Filter
	SetCursor(x, y int)
}

type RippleCursor struct {

	filter *RippleFilter
	mutex sync.Mutex
	index int
	blinkCountdown int
	screen TextScreen
}

func NewRippleCursor(brightness float64, transform func(Vector2) Vector2, screen TextScreen) Cursor {
	r := new(RippleCursor)
	r.screen = screen
	r.filter = NewRippleFilter(brightness, transform, screen)
    return r
}

func (r *RippleCursor) SetCursor(x, y int) {

	r.mutex.Lock()
	r.index = r.screen.DigitIndex(x, y)
	r.blinkCountdown = 100
	r.mutex.Unlock()
	if r.index != -1 {
		r.filter.SetRippleOrigin(r.filter.coords[r.index*16+3])
	}
}

func (r *RippleCursor) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	r.mutex.Lock()
	index := r.index
	r.blinkCountdown = (r.blinkCountdown+199)%200
	countdown := r.blinkCountdown
	r.mutex.Unlock()

	if index != -1 && countdown < 100 {
		b := 1.
		if countdown > 90 {
			b -= float64(countdown-90)/20.
		}
		for i:=0; i<14; i++ {
			if f.frame[index*16 + i] < b {
				f.frame[index*16 + i] = b
			}
		}
	}
	if index != -1 && countdown == 0 {
		r.filter.Ripple()
	}

	r.filter.Render(f, old, tick)
}

