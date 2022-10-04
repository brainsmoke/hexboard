
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
}

func NewRippleCursor(brightness float64) Cursor {
	r := new(RippleCursor)
	r.filter = NewRippleFilter(brightness, nil)
    return r
}

func (r *RippleCursor) SetCursor(x, y int) {

	r.mutex.Lock()
	r.index = screenInfo.GetIndex(x, y)
	r.blinkCountdown = 100
	r.mutex.Unlock()
	if r.index != -1 {
		r.filter.SetRippleOrigin(r.filter.coords[r.index])
	}
}

func (r *RippleCursor) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	r.mutex.Lock()
	index := r.index
	r.blinkCountdown = (r.blinkCountdown+199)%200
	countdown := r.blinkCountdown
	r.mutex.Unlock()

	if index != -1 && countdown < 100 {
		for i:=0; i<14; i++ {
			if f.frame[index*16 + i] < .5 {
				f.frame[index*16 + i] = .5
			}
		}
	}
	if index != -1 && countdown == 0 {
		r.filter.Ripple()
	}

	r.filter.Render(f, old, tick)
}

