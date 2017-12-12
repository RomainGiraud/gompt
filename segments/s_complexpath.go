package segments

import(
    "os"
    "log"
    "strings"
)


type ComplexPath struct {
    style Style
    separator string
    fgSeparator Color
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

    FormatStringArray(dir_s, p.separator, p.style, segments, current)
}

func (p ComplexPath) GetStyle(segments []Segment, current int) Style {
    return p.style
}

func NewComplexPath(style Style, separator string, fgSeparator Color) Segment {
    return &ComplexPath{ style, separator, fgSeparator }
}

func LoadComplexPath(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    var sep, _  = config["separator"].(string);
    var fg_sep, _  = config["fg-separator"].(string);
    return &ComplexPath{ style, sep, NewColor(fg_sep) }
}

func init() {
    RegisterSegmentLoader("complex-path", LoadComplexPath)
}
