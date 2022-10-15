
package screen

import (
	"time"
	"fmt"
)

const Fps = 60

type Output interface {

	Write(dst []float64) (int, error)
}

func DisplayRoutine(out Output, s Screen, info ScreenInfo, quit <-chan bool) {

	var counter, prev_counter uint64
	var frames = []*FrameBuffer { NewFrameBuffer(info.SegmentCount()), NewFrameBuffer(info.SegmentCount()) }
	cur, old := 0, 1

	if !s.NextFrame(frames[cur], frames[old], counter) {
		return
	}

	tick := time.NewTicker(time.Second / time.Duration(Fps))
	seconds := time.NewTicker(time.Second)

	loop: for {
		select {

			case <-quit:

				tick.Stop()
				break loop

			case <-tick.C:

				out.Write(frames[cur].frame)
				cur, old = old, cur
				counter++
				frames[cur].Clear()
				if !s.NextFrame(frames[cur], frames[old], counter) {
					return
				}

			case <-seconds.C:
				fmt.Printf("fps: %d\n", counter-prev_counter)
				prev_counter = counter
		}
	}

}


