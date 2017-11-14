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

func (s Separator) Print(context Context, index int) {
    fmt.Print(s.style.Format(s.value, context, index))
}

func (s Separator) GetStyle() Style {
    return s.style
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
