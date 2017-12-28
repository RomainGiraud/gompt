package segments

import(
    "io"
)


type Binding struct {
    segment1 Segment
    segment2 Segment
}

func (s Binding) Load() []Segment {
    s1 := s.segment1.Load()
    s2 := s.segment2.Load()
    if len(s1) == 0 || len(s2) == 0 {
        return []Segment{}
    }
    return append(s1, s2...)
}

func (s Binding) Print(writer io.Writer, segments []Segment, current int) {
}

func (s Binding) GetStyle(segments []Segment, current int) Style {
    return s.segment1.GetStyle(segments, current)
}

func NewBinding(segment1 Segment, segment2 Segment) *Binding {
    return &Binding{ segment1, segment2 }
}