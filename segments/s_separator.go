package segments

import(
    "io"
)


type Separator struct {
    style Style
    value string
}

func (s Separator) Print(writer io.Writer, segments []Segment, current int) {
    FormatString(writer, s.value, s.style, segments, current)
}

func (s Separator) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewSeparator(text string, style Style) Segment {
    return &Separator{ style, text }
}