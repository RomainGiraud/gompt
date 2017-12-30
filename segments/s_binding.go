package segments

// BindingLoader links two segment loaders.
// If one of them is empty, none is displayed.
type BindingLoader struct {
	Segment1 SegmentLoader
	Segment2 SegmentLoader
}

func (s BindingLoader) Load() []Segment {
	s1 := s.Segment1.Load()
	s2 := s.Segment2.Load()
	if len(s1) == 0 || len(s2) == 0 {
		return []Segment{}
	}
	return append(s1, s2...)
}
