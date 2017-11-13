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

func (t Text) String() string {
    return t.style(t.value)
}

type textConfig struct {
    Text string `json:"text"`
}

func NewText(bytes json.RawMessage, style color.Style) fmt.Stringer {
    var config textConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Text{ style.GetFmt(), config.Text }
}
