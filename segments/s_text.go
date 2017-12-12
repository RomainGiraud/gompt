package segments

import(
    "os"
)


type Text struct {
    style Style
    value string
}

func (t Text) Print(segments []Segment, current int) {
    FormatString(os.ExpandEnv(t.value), t.style, segments, current)
}

func (t Text) GetStyle(segments []Segment, current int) Style {
    return t.style
}

func NewText(style Style, text string) Segment {
    return &Text{ style, text }
}

func LoadText(config map[string]interface{}) Segment {
    var style, _ = LoadStyle(config["style"])
    var text, _  = config["text"].(string)
    return &Text{ style, text }
}

func init() {
    RegisterSegmentLoader("text", LoadText)
}
