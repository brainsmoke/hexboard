
package screen

import (
	"math"
	"sync"
)

type Filter interface {

	Render(f *FrameBuffer, old *FrameBuffer, tick uint64)
}

type Cursor interface {
	Filter
	SetCursor(index int)
}

type afterGlowFilter struct {
	factor float64
}

func NewAfterGlowFilter(factor float64) Filter {
	return &afterGlowFilter{factor: math.Min(math.Max(0, factor), 1)}
}

func (s *afterGlowFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for i := range f.frame {
		fade := old.frame[i] * s.factor
		if fade > f.frame[i] {
			f.frame[i] = fade
		}
	}
}

type RippleCache struct {
	d [960*16]uint32
	maxD uint32
}

type Ripple struct {
	cache *RippleCache
	count uint32
}

func transform(pos Vector2) (Vector2) {
	return Vector2 { X: (1000-pos.X)*550/(pos.Y+500), Y: 1100000/(pos.Y+500) }
}

type RippleCursor struct {

	rippleCache *RippleCache
	ripples []*Ripple
	coords [960*16]Vector2
	mutex sync.Mutex
	index int
	blinkCountdown int
}

func NewRippleCursor() Cursor {
	r := new(RippleCursor)
	for i := range screenInfo.Segments {
		r.coords[i] = transform(screenInfo.Segments[i].Position)
	}
    return r
}

func (r *RippleCursor) SetCursor(index int) {

	cache := new(RippleCache)
	pos := r.coords[index*16+3]
	maxD := uint32(0)

	for i := range r.coords {
		p := r.coords[i]
		dx, dy := (p.X - pos.X), (p.Y - pos.Y)
		d := uint32(math.Sqrt( dx*dx + dy*dy ))
		if d > maxD {
			maxD = d
		}
		cache.d[i] = d
		cache.maxD = maxD
	}

	r.mutex.Lock()
	r.rippleCache = cache
	r.index = index
	r.blinkCountdown = 50
	r.mutex.Unlock()
}

func (r *RippleCursor) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	r.mutex.Lock()
	cache := r.rippleCache
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
		r.ripples = append(r.ripples, &Ripple{ cache: cache } )
	}

	for j := range r.ripples {
		c := r.ripples[j].count
		r.ripples[j].count+=2
		for i, d := range r.ripples[j].cache.d {
			if c > d && c < d+30 {
				x := float64(.05)
//				x := float64(.05)
				if c > d+2 && c < d+5 {
					x = float64(.15)
//					x = float64(.15)
				}
	//			if x > f.frame[i] {
	//				f.frame[i] = x
					f.frame[i] += x
	//			}
			}
		}
	}

	l := len(r.ripples)

	for j:=l-1; j >= 0; j-- {
		if r.ripples[j].count > r.ripples[j].cache.maxD + 30 {
			r.ripples[j] = r.ripples[l-1]
			l--
		}
	}
	r.ripples = r.ripples[:l]
}


