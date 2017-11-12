package segments

import(
    "os"
    "log"
    "fmt"
    "encoding/json"
    "prompt/color"
)


type Path struct {
}

func (p Path) String() string {
    dir, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    return dir
}

type pathConfig struct {
}

func NewPath(bytes json.RawMessage, style color.StyleConfig) fmt.Stringer {
    var config pathConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Path{}
}
