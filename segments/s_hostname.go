package segments

import(
    "io"
    "os"
)


type Hostname struct {
    style Style
}

func (h Hostname) Print(writer io.Writer, segments []Segment, current int) {
    n, _ := os.Hostname()
    FormatString(writer, " " + n + " ", h.style, segments, current)
}

func (h Hostname) GetStyle(segments []Segment, current int) Style {
    return h.style
}

func NewHostname(style Style) Segment {
    return &Hostname{ style }
}