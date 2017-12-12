package segments

import(
    "io"
    "os"
    "log"
)


type Path struct {
    style Style
}

func (p Path) Print(writer io.Writer, segments []Segment, current int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    FormatString(writer, dir, p.style, segments, current)
}

func (p Path) GetStyle(segments []Segment, current int) Style {
    return p.style
}

func NewPath(style Style) Segment {
    return &Path{ style }
}