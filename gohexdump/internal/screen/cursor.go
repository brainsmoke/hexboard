
package screen

import (
	"sync"
)

type Cursor interface {
	Filter
	SetCursor(index int)
}

type RippleCursor struct {

	filter *RippleFilter
	mutex sync.Mutex
	index int
	blinkCountdown int
}

func NewRippleCursor(brightness float64) Cursor {
	r := new(RippleCursor)
	r.filter = NewRippleFilter(brightness, nil)
    return r
}

func (r *RippleCursor) SetCursor(index int) {

	r.mutex.Lock()
	r.index = index
	r.blinkCountdown = 50
	r.mutex.Unlock()
	r.filter.SetRippleOrigin(r.filter.coords[index*16+3])
}

func (r *RippleCursor) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	r.mutex.Lock()
	index := r.index
	r.blinkCountdown = (r.blinkCountdown+99)%100
	countdown := r.blinkCountdown
	r.mutex.Unlock()

	if countdown < 50 {
		for i:=0; i<14; i++ {
			if f.frame[index*16 + i] < .5 {
				f.frame[index*16 + i] = .5
			}
		}
	}
	if countdown == 0 {
		r.filter.Ripple()
	}

	r.filter.Render(f, old, tick)
}

