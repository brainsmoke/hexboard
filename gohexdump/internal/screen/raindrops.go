
package screen

import (
	"math"
	"math/rand"
)


type raindrop struct {
	column, endRow int
	yPos, ySpeed float64
	brightness float64
	ripple *RippleFilter
	info *ScreenInfo
}

func newRaindrop(column int, ripple *RippleFilter, info *ScreenInfo) *raindrop {
	drop := &raindrop{ column: column, ripple: ripple, info: info }
	drop.reset()
	return drop
}

var columns = [...]int{ 0, 3, 6,    13, 16, 19, 22, 25, 28, 31, 34,    41, 44, 47, 50, 53, 56, 59, 62,     69, 72, 75, 77, 80, 83 };

func (d *raindrop) reset() {

	d.endRow = 6+rand.Intn(d.info.Rows*3)
	d.yPos = float64(-rand.Intn(60))
	d.ySpeed = float64(20+d.endRow)/float64(256)
	d.brightness = .1 + .1*float64(rand.Intn(5))
}

func (d *raindrop) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {
	oldPos := d.yPos
	newPos := oldPos+d.ySpeed
	oldRow := int(math.Floor(oldPos))
	newRow := int(math.Floor(newPos))

	if newRow > d.info.Rows {
		d.reset()
		return
	}

	if oldRow != newRow {
		if rand.Intn(2) != 0 {
			newPos -= 1
			newRow -= 1
		} else {
			if oldRow == d.endRow {
				oldIndex := d.info.GetIndex(d.column, oldRow)
				if oldIndex !=-1 {
					pa, pb := d.ripple.coords[oldIndex*16+3], d.ripple.coords[oldIndex*16+16+3]
					p := Vector2 { X: (pa.X+pb.X)/2, Y : (pa.Y+pb.Y)/2 }
					d.ripple.RippleAt(p)
				}
				d.reset()
				return
			}
		}
	}
	if oldRow != newRow || math.Floor(oldPos+.25) != math.Floor(newPos+.25) {

		index := d.info.GetIndex(d.column, newRow)

		if index != -1 {
			b := d.brightness

			r := rand.Uint32()
			for i := 0; i< 32; i++ {
				if (1<<uint32(i)) & r != 0 {
					f.frame[index*16 + i] += b
				}
			}
		}
	}

	d.yPos = newPos
}

type RaindropFilter struct {
	drops []*raindrop
	buf *FrameBuffer
	ripple *RippleFilter
}

func symmetricTransform(pos Vector2) Vector2 {
	return Vector2 { X: (pos.X-600)*550/(pos.Y+500)/5, Y: 1100000/(pos.Y+500)/5 }
//    return Vector2 { X: pos.X/2.5, Y: pos.Y }
}


func NewRaindropFilter(info *ScreenInfo) Filter {
	f := new(RaindropFilter)
	f.drops = make([]*raindrop, len(columns))
	f.ripple = NewRippleFilter(.1, symmetricTransform, info)
	f.buf = NewFrameBuffer(info.Size*16)
	for i := range f.drops {
		f.drops[i] = newRaindrop(columns[i], f.ripple, info)
	}
	return f
}

func (r *RaindropFilter) Render(f *FrameBuffer, old *FrameBuffer, tick uint64) {

	for i := range r.buf.frame {
		r.buf.frame[i] *= .98
	}

	for i:= range r.drops {
		r.drops[i].Render(r.buf, nil, tick)
	}

	for i := range f.frame {
		f.frame[i] += r.buf.frame[i]
	}
	r.ripple.Render(f, old, tick)
}
