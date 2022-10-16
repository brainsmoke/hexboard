
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

/*
import math
a = [ math.sin(math.tau/128*x) / (2+(x/8)) for x in range(128) ]
a = [ x/max(a) for x in a ]
print(a)
*/
var ripw = [...]float64{
0.0, 0.12436561881943207, 0.23462987274943167, 0.3327518688729652, 0.42030011840629583, 0.49854611579437286, 0.5685323922167587, 0.63112281089507, 0.6870402824743831, 0.736895421717973, 0.7812085833126956, 0.8204269922901144, 0.8549381944439929, 0.8850807141101417, 0.9111525700664814, 0.9334181323977383, 0.9521136824832831, 0.9674519504968202, 0.9796258402716647, 0.9888115034466225, 0.995170888854884, 0.9988538659122841, 1.0, 0.9987400418676564, 0.9951971807077786, 0.9894881008949165, 0.9817238747954508, 0.972010718055588, 0.960450629004155, 0.9471419299906602, 0.9321797254100731, 0.9156562886848425, 0.8976613884592332, 0.8782825626160734, 0.8576053473772672, 0.8357134676388177, 0.8126889937727467, 0.7886124693658544, 0.7635630137297951, 0.7376184024853014, 0.7108551290769847, 0.6833484496989229, 0.6551724137931035, 0.6263998820127884, 0.5971025333129306, 0.567350862633308, 0.5372141704715581, 0.5067605454983344, 0.4760568412416416, 0.44516864775892145, 0.4141602591211279, 0.38309463745065836, 0.35203337418282055, 0.3210366491570013, 0.2901631880875728, 0.25947021891477917, 0.2290134274914611, 0.1988469130217346, 0.16902314363200938, 0.1395929124224245, 0.11060529431744648, 0.08210760400761484, 0.054145355249868965, 0.02676222177127017, 6.595908876568612e-17, -0.026101426171979417, -0.05150411840841182, -0.0761721145612811, -0.10007145676340383, -0.12317021684331562, -0.1454385189391707, -0.166848559202145, -0.18737462249301354, -0.2069930959881945, -0.22568247962366764, -0.24342339331684706, -0.26019858091773684, -0.2759929108515495, -0.2907933734254727, -0.3045890747824199, -0.3173712274944277, -0.3291331377978872, -0.339870189482006, -0.3495798244508262, -0.3582615199877583, -0.3659167627599456, -0.37254901960784315, -0.37816370517319026, -0.38276814642606866, -0.38637154415896746, -0.38898493152272573, -0.3906211296858905, -0.3912947007053965, -0.39102189770256607, -0.3898206124442124, -0.38771032043412246, -0.38471202362538565, -0.3808481908689168, -0.3761426962180997, -0.3706207552137365, -0.3643088592774381, -0.35723470834521615, -0.34942714187634694, -0.34091606837555954, -0.33173239356925954, -0.32190794737883144, -0.31147540983606564, -0.3004682360874351, -0.288920580635289, -0.27686722096505445, -0.26434348070822694, -0.2513851524912998, -0.23802842062082083, -0.22430978375449537, -0.2102659777076497, -0.19593389854346643, -0.1813505260941803, -0.16655284805889548, -0.1515777848218664, -0.13646211513295808, -0.12124240278959722, -0.10595492445683657, -0.09063559875919346, -0.07531991677468952, -0.0600428740580424, -0.044838904316215365, -0.02974181485556173, -0.014784723915596843,
}
func (r *RippleFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for j := range r.ripples {
		c := r.ripples[j].count
		full := r.brightness
		//third := full/3
		r.ripples[j].count+=1
		for i, d := range r.ripples[j].cache.d {
			w := int(c)-int(d)
			if w >= 0 && w < int(len(ripw)) {
				x := full * ripw[w]
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

