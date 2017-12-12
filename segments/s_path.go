package segments

import(
    "os"
    "log"
)


type Path struct {
    style Style
}

func (p Path) Print(segments []Segment, current int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    FormatString(dir, p.style, segments, current)
}

func (p Path) GetStyle(segments []Segment, current int) Style {
    return p.style
}

func NewPath(style Style) Segment {
    return &Path{ style }
}