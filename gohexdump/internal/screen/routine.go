
package screen

import (
	"time"
//"fmt"
)

const Fps = 50

type Output interface {

	Write(dst []float64) (int, error)
}

func DisplayRoutine(out Output, s Screen, quit <-chan bool) {

	var counter uint64
	var frames = []*FrameBuffer { NewFrameBuffer(), NewFrameBuffer() }
	cur, old := 0, 1

	if !s.NextFrame(frames[cur], frames[old], counter) {
		return
	}

	tick := time.NewTicker(time.Second / time.Duration(Fps))

	loop: for {
		select {

			case <-quit:

				tick.Stop()
				break loop

			case <-tick.C:

				out.Write(frames[cur].frame)
				cur, old = old, cur
				counter++
				if !s.NextFrame(frames[cur], frames[old], counter) {
					return
				}
		}
	}

}


