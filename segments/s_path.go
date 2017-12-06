package segments

import(
    "fmt"
    "os"
    "log"
)


type Path struct {
    style Style
}

func (p Path) Print(context Context, index int) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(p.style.Format(dir, context, index, 0))
}

func (p Path) GetStyle(context Context, index int) Style {
    return p.style
}

func NewPath(style Style) Segment {
    return &Path{ style }
}

func LoadPath(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    return &Path{ style }
}

func init() {
    RegisterSegmentLoader("path", LoadPath)
}
