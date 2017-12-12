package segments

import(
    "os"
)


type Hostname struct {
    style Style
}

func (h Hostname) Print(segments []Segment, current int) {
    n, _ := os.Hostname()
    FormatString(n, h.style, segments, current)
}

func (h Hostname) GetStyle(segments []Segment, current int) Style {
    return h.style
}

func NewHostname(style Style) Segment {
    return &Hostname{ style }
}