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

func (s ComplexPath) Load() []Segment {
    return []Segment{ s }
}

func (s ComplexPath) Print(writer io.Writer, segments []Segment, current int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    home_dir := os.Getenv("HOME")
    if strings.Index(dir, home_dir) == 0 {
        dir = strings.Replace(dir, home_dir, "~", 1)
    }

    dir_s := strings.Split(dir, "/")
    if s.maxDepth != 0 && len(dir_s) > int(s.maxDepth) {
        dir_s = dir_s[len(dir_s) - int(s.maxDepth):]
        dir_s[0] = s.ellipsis
    }
    for i, v := range dir_s {
        dir_s[i] = " " + v + " "
    }

    if s.isPlain {
        FormatStringArrayPlain(writer, dir_s, s.style, s.separator, s.fgSeparator, segments, current)
    } else {
        FormatStringArrayBlock(writer, dir_s, s.style, s.separator, StyleChameleon{ }, segments, current)
    }
}

func (s ComplexPath) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewComplexPathPlain(style Style, separator string, fgSeparator Color, maxDepth uint, ellipsis string) Segment {
    return &ComplexPath{ style, true, separator, fgSeparator, maxDepth, ellipsis }
}

func NewComplexPathSplitted(style Style, separator string, maxDepth uint, ellipsis string) Segment {
    return &ComplexPath{ style, false, separator, nil, maxDepth, ellipsis }
}