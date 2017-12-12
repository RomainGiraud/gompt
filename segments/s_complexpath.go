package segments

import(
    "os"
    "log"
    "strings"
)


type ComplexPath struct {
    style Style
    separator string
    fgSeparator Brush
}

func (p ComplexPath) Print(segments []Segment, current int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    home_dir := os.Getenv("HOME")
    if strings.Index(dir, home_dir) == 0 {
        dir = strings.Replace(dir, home_dir, "~", 1)
    }
    dir_s := strings.Split(dir, "/")

    FormatStringArray(dir_s, p.style, p.separator, p.fgSeparator, segments, current)
}

func (p ComplexPath) GetStyle(segments []Segment, current int) Style {
    return p.style
}

func NewComplexPath(style Style, separator string, fgSeparator Brush) Segment {
    return &ComplexPath{ style, separator, fgSeparator }
}