package segments

import(
    _"log"
    _"encoding/json"
)


type Separator struct {
    style Style
    value string
}

func (s Separator) Print(segments []Segment, current int) {
    FormatString(s.value, s.style, segments, current)
}

func (s Separator) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewSeparator(text string, style Style) Segment {
    return &Separator{ style, text }
}

func LoadSeparator(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    return &Separator{ style, config["text"].(string) }
}

func init() {
    RegisterSegmentLoader("separator", LoadSeparator)
}
