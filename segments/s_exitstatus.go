package segments

import(
    "fmt"
    "log"
    "encoding/json"
)


type ExitStatus struct {
    style Style
}

func (e ExitStatus) Print(context Context, index int) {
    if context.Args.Status == 0 {
        fmt.Print(e.style.Format("OK", context, index))
    } else {
        fmt.Print(e.style.Format("NO", context, index))
    }
}

func (e ExitStatus) GetStyle() Style {
    return e.style
}

type exitStatusConfig struct {
}

func NewExitStatus(bytes json.RawMessage, style Style) Segment {
    var config exitStatusConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &ExitStatus{ style }
}
