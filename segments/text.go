package segments

import(
    "fmt"
    "log"
    "encoding/json"
    "prompt/color"
)


type Text struct {
    style color.StyleFmt
    value string
}

func (t Text) Print(context Context, name string) {
    fmt.Print(t.style(t.value))
}

type textConfig struct {
    Text string `json:"text"`
}

func NewText(bytes json.RawMessage, style color.Style) Segment {
    var config textConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Text{ style.GetFmt(), config.Text }
}
