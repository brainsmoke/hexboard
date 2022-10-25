
package screen

type ExitScreen struct {
	end, cur uint64
}

func NewExitScreen(fadeTime float64) Screen {

	return &ExitScreen{ end: uint64(fadeTime*Fps)+1 }
}

func (s *ExitScreen) NextFrame(f, old *FrameBuffer, tick uint64) bool {

	if s.cur < s.end {

		d := 1 - 1/float64(s.end-s.cur)
		for i,v := range old.frame {
			f.frame[i] = d*v
		}
		s.cur += 1
		return true

	} else {
		return false
	}
}
