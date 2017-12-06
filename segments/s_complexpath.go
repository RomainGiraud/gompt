package segments

import(
    "os"
    "log"
    "strings"
)


type ComplexPath struct {
    style Style
    separator string
}

func (p ComplexPath) Print(context Context, index int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    home_dir := os.Getenv("HOME")
    if strings.Index(dir, home_dir) == 0 {
        dir = strings.Replace(dir, home_dir, "~", 1)
    }
    dir_s := strings.Split(dir, "/")
    FormatStringArray(dir_s, p.separator, p.style, context, index)
}

func (p ComplexPath) GetStyle(context Context, index int) Style {
    return p.style
}

func NewComplexPath(style Style, separator string) Segment {
    return &ComplexPath{ style, separator }
}

func LoadComplexPath(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    var sep, _  = config["separator"].(string);
    return &ComplexPath{ style, sep }
}

func init() {
    RegisterSegmentLoader("complex-path", LoadComplexPath)
}
