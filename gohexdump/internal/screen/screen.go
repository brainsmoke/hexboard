
package screen

type Screen interface {

	NextFrame(f *FrameBuffer, old *FrameBuffer, tick uint64) bool
}

type MultiScreen struct {
	c <-chan Screen
	s Screen
}

func NewMultiScreen() (Screen, chan<-Screen) {
	c := make(chan Screen, 1)
	return &MultiScreen{ c: c }, c
}

func (m *MultiScreen) NextFrame(f *FrameBuffer, old *FrameBuffer, tick uint64) bool {
	var ok bool

	if (m.s == nil) {
		m.s, ok = <-m.c
		if !ok {
			return false
		}
	}

	for {
		select {
			case m.s, ok = <-m.c:
				if !ok {
					return false
				}
			default:
				return m.s.NextFrame(f, old, tick)
		}
	}
}

type filterScreen struct {
	s Screen
	filters []Filter
}

func NewFilterScreen(s Screen, filters []Filter) Screen {
	return &filterScreen{s:s, filters:filters}
}


func AfterGlow(s Screen, factor float64) Screen {
	return NewFilterScreen(s, []Filter{ NewAfterGlowFilter(factor) })
}

func (s *filterScreen) NextFrame(f *FrameBuffer, old *FrameBuffer, tick uint64) bool {

	if !s.s.NextFrame(f, old, tick) {
		return false
	}

	for i := range s.filters {
		s.filters[i].Render(f, old, tick)
	}

	return true
}
