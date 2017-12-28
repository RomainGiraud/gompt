package segments

import(
    "io"
    "os"
)


type Hostname struct {
    style Style
}

func (s Hostname) Load() []Segment {
    return []Segment{ s }
}

func (s Hostname) Print(writer io.Writer, segments []Segment, current int) {
    h, _ := os.Hostname()
    FormatString(writer, " " + h + " ", s.style, segments, current)
}

func (s Hostname) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewHostname(style Style) *Hostname {
    return &Hostname{ style }
}