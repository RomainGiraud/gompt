package segments

import(
    _"log"
    _"encoding/json"
)


type Separator struct {
    style Style
    value string
}

func (s Separator) Print(segments []Segment, current int) {
    FormatString(s.value, s.style, segments, current)
}

func (s Separator) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewSeparator(text string, style Style) Segment {
    return &Separator{ style, text }
}