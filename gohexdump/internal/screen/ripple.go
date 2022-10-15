
package screen

import (
	"math"
	"sync"
)

type RippleCache struct {
	d []uint32
	maxD uint32
	origin Vector2
}

type Ripple struct {
	cache *RippleCache
	count uint32
}

func RippleTransform(pos Vector2) Vector2 {
	return Vector2 { X: (1000-pos.X)*550/(pos.Y+500), Y: 1100000/(pos.Y+500) }
}

type RippleFilter struct {

	coords []Vector2

	cache *RippleCache

	ripples []*Ripple
	mutex sync.Mutex

	brightness float64
}

func NewRippleFilter(brightness float64, transform func(Vector2) Vector2, screen TextScreen) *RippleFilter {

	if transform == nil {
		transform = RippleTransform
	}

	r := new(RippleFilter)
	r.coords = make([]Vector2, screen.SegmentCount())
	r.brightness = math.Max(0, math.Min(1, brightness))
	for i := range r.coords {
		r.coords[i] = transform(screen.SegmentCoord(i))
	}
    return r
}

func (r *RippleFilter) SetRippleOrigin(origin Vector2) {
	r.rippleCache(origin)
}

func (r *RippleFilter) rippleCache(origin Vector2) *RippleCache {

	r.mutex.Lock()
	cache := r.cache
	r.mutex.Unlock()

	if cache == nil || origin != cache.origin {

		cache = new(RippleCache)
		cache.d = make([]uint32, len(r.coords))
		cache.origin = origin
		maxD := uint32(0)

		for i := range r.coords {
			p := r.coords[i]
			dx, dy := (p.X - origin.X), (p.Y - origin.Y)
			d := uint32(math.Sqrt( dx*dx + dy*dy ))
			if d > maxD {
				maxD = d
			}
			cache.d[i] = d
			cache.maxD = maxD
		}

		r.mutex.Lock()
		r.cache = cache
		r.mutex.Unlock()
	}

	return cache
}

func (r *RippleFilter) RippleAt(origin Vector2) {
	cache := r.rippleCache(origin)
	r.mutex.Lock()
	r.ripples = append(r.ripples, &Ripple{ cache: cache } )
	r.mutex.Unlock()
}

func (r *RippleFilter) Ripple() {
	r.mutex.Lock()
	cache := r.cache
	r.ripples = append(r.ripples, &Ripple{ cache: cache } )
	r.mutex.Unlock()
}

func (r *RippleFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for j := range r.ripples {
		c := r.ripples[j].count
		full := r.brightness * ( 100/float64(c+100) )
		third := full/3
		r.ripples[j].count+=1
		for i, d := range r.ripples[j].cache.d {
			if c > d && c < d+30 {
				x := third
				if c > d+2 && c < d+5 {
					x = full
				}
				f.frame[i] += x
			}
		}
	}

	l := len(r.ripples)

	r.mutex.Lock()
	for j:=l-1; j >= 0; j-- {
		if r.ripples[j].count > r.ripples[j].cache.maxD + 30 {
			r.ripples[j] = r.ripples[l-1]
			l--
		}
	}
	r.ripples = r.ripples[:l]
	r.mutex.Unlock()
}

