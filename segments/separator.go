package segments

import(
    "fmt"
    "log"
    "encoding/json"
)


type Separator struct {
    style Style
    value string
}

func (s Separator) Print(context Context, name string) {
    fmt.Print(s.style.Format(s.value, context, name))
}

type separatorConfig struct {
    Text string `json:"text"`
}

func NewSeparator(bytes json.RawMessage, style Style) Segment {
    var config separatorConfig
    err := json.Unmarshal(bytes, &config)
    if err != nil {
        log.Fatal(err)
    }
    return &Separator{ style, config.Text }
}
