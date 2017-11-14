package segments

import(
    "fmt"
    "log"
    "encoding/json"
)


type Text struct {
    style Style
    value string
}

func (t Text) Print(context Context, name string) {
    fmt.Print(t.style.Format(t.value, context, name))
}

func (t Text) GetStyle() Style {
    return t.style
}

type textConfig struct {
    Text string `json:"text"`
}

func NewText(bytes json.RawMessage, style Style) Segment {
    var config textConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Text{ style, config.Text }
}
