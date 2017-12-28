package segments

import(
    "io"
    "os"
    "log"
    "strings"
)


type ComplexPath struct {
    style Style
    styleUnit Style
    isPlain bool
    separator string
    fgSeparator Color
    maxDepth uint
    ellipsis string
    directories []string
}

func (s ComplexPath) Load() []Segment {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    home_dir := os.Getenv("HOME")
    if strings.Index(dir, home_dir) == 0 {
        dir = strings.Replace(dir, home_dir, "~", 1)
    }

    s.directories = strings.Split(dir, "/")
    if s.maxDepth != 0 && len(s.directories) > int(s.maxDepth) {
        s.directories = s.directories[len(s.directories) - int(s.maxDepth):]
        s.directories[0] = s.ellipsis
    }

    return []Segment{ s }
}

func (s ComplexPath) Print(writer io.Writer, segments []Segment, current int) {
    for i, v := range s.directories {
        s.directories[i] = " " + v + " "
    }

    if s.isPlain {
        ff := []PartFormatter{}
        for i := 0; i < len(s.directories) - 1; i += 1 {
            ff = append(ff, PartFormatter{ s.directories[i], nil, nil })
            ff = append(ff, PartFormatter{ s.separator, s.fgSeparator, nil })
        }
        ff = append(ff, PartFormatter{ s.directories[len(s.directories) - 1], nil, nil })
        FormatParts(writer, s.style, segments, current, ff)
    } else {
        FormatStringArrayBlock(writer, s.directories, s.style, s.separator, StyleChameleon{ }, segments, current)
    }
}

func (s ComplexPath) GetStyle(segments []Segment, current int) Style {
    if !s.isPlain && len(s.directories) == 1 {
        return s.styleUnit
    }
    return s.style
}

func NewComplexPathPlain(style Style, styleUnit Style, separator string, fgSeparator Color, maxDepth uint, ellipsis string) *ComplexPath {
    return &ComplexPath{ style, styleUnit, true, separator, fgSeparator, maxDepth, ellipsis, []string{} }
}

func NewComplexPathSplitted(style Style, styleUnit Style, separator string, maxDepth uint, ellipsis string) *ComplexPath {
    return &ComplexPath{ style, styleUnit, false, separator, nil, maxDepth, ellipsis, []string{} }
}