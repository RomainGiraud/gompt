package segments

import(
    "fmt"
    "os"
    "log"
    "encoding/json"
    "prompt/color"
)


type Path struct {
    style color.StyleFmt
}

func (p Path) Print(context Context, name string) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(p.style(dir))
}

type pathConfig struct {
}

func NewPath(bytes json.RawMessage, style color.Style) Segment {
    var config pathConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Path{ style.GetFmt() }
}
