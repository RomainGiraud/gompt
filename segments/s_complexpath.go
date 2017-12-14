package segments

import(
    "io"
    "os"
    "log"
    "strings"
)


type ComplexPath struct {
    style Style
    isPlain bool
    separator string
    fgSeparator Color
    maxDepth uint
    ellipsis string
}

func (p ComplexPath) Print(writer io.Writer, segments []Segment, current int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    home_dir := os.Getenv("HOME")
    if strings.Index(dir, home_dir) == 0 {
        dir = strings.Replace(dir, home_dir, "~", 1)
    }

    dir_s := strings.Split(dir, "/")
    if p.maxDepth != 0 && len(dir_s) > int(p.maxDepth) {
        dir_s = dir_s[len(dir_s) - int(p.maxDepth):]
        dir_s[0] = p.ellipsis
    }
    for i, v := range dir_s {
        dir_s[i] = " " + v + " "
    }

    if p.isPlain {
        FormatStringArrayPlain(writer, dir_s, p.style, p.separator, p.fgSeparator, segments, current)
    } else {
        FormatStringArrayBlock(writer, dir_s, p.style, p.separator, StyleChameleon{ }, segments, current)
    }
}

func (p ComplexPath) GetStyle(segments []Segment, current int) Style {
    return p.style
}

func NewComplexPathPlain(style Style, separator string, fgSeparator Color, maxDepth uint, ellipsis string) Segment {
    return &ComplexPath{ style, true, separator, fgSeparator, maxDepth, ellipsis }
}

func NewComplexPathSplitted(style Style, separator string, maxDepth uint, ellipsis string) Segment {
    return &ComplexPath{ style, false, separator, nil, maxDepth, ellipsis }
}