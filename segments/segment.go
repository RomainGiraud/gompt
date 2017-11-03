package segments

import(
    "github.com/fatih/color"
)


type Segment interface {
    GetFg() color.Attribute
    GetBg() color.Attribute
}

type Style struct {
    Fg color.Attribute
    Bg color.Attribute
}

func (s Style) GetFg() color.Attribute {
    return s.Fg
}

func (s Style) GetBg() color.Attribute {
    return s.Bg
}

