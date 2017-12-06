package segments

import(
    "os"
)


type Text struct {
    style Style
    value string
}

func (t Text) Print(context Context, index int) {
    FormatString(os.ExpandEnv(t.value), t.style, context, index)
}

func (t Text) GetStyle(context Context, index int) Style {
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
