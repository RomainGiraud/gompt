package segments

import(
    "fmt"
    "os"
    "log"
    "encoding/json"
)


type Path struct {
    style Style
}

func (p Path) Print(context Context, name string) {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(p.style.Format(dir, context, name))
}

type pathConfig struct {
}

func NewPath(bytes json.RawMessage, style Style) Segment {
    var config pathConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Path{ style }
}
