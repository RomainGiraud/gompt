package segments

import(
    "io"
    "os"
)


type Hostname struct {
    style Style
}

func (s Hostname) Print(writer io.Writer, segments []Segment, current int) {
    n, _ := os.Hostname()
    FormatString(writer, " " + n + " ", s.style, segments, current)
}

func (s Hostname) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewHostname(style Style) Segment {
    return &Hostname{ style }
}