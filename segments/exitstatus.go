package segments

import(
    "github.com/fatih/color"
)


type ExitStatus struct {
    StyleOk Style
    StyleError Style
    Value int
}

func (e ExitStatus) String() string {
    return ""
}

func (e ExitStatus) GetFg() color.Attribute {
    if e.Value != 0 {
        return e.StyleError.Fg
    }
    return e.StyleOk.Fg
}

func (e ExitStatus) GetBg() color.Attribute {
    if e.Value != 0 {
        return e.StyleError.Bg
    }
    return e.StyleOk.Bg
}
