package segments

import(
    "os/user"
)


type Username struct {
    style Style
}

func (u Username) Print(segments []Segment, current int) {
    uc, _ := user.Current()
    FormatString(uc.Username, u.style, segments, current)
}

func (u Username) GetStyle(segments []Segment, current int) Style {
    return u.style
}

func NewUsername(style Style) Segment {
    return &Username{ style }
}