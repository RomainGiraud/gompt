package segments

import(
    "fmt"
    "log"
    "encoding/json"
    "prompt/color"
)


type Text struct {
    value string
}

func (t Text) String() string {
    return t.value
}

type textConfig struct {
    Text string `json:"text"`
}

func NewText(bytes json.RawMessage, style color.StyleConfig) fmt.Stringer {
    var config textConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Text{ config.Text }
}
