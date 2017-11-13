package segments

import(
    "fmt"
    "log"
    "encoding/json"
    "prompt/color"
)


type Separator struct {
    style color.StyleFmt
    value string
}

func (s Separator) Print(context Context, name string) {
    fmt.Print(s.style(s.value))
}

type separatorConfig struct {
    Text string `json:"text"`
}

func NewSeparator(bytes json.RawMessage, style color.Style) Segment {
    var config separatorConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Separator{ style.GetFmt(), config.Text }
}
