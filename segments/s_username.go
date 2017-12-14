package segments

import(
    "io"
    "os/user"
)


type Username struct {
    style Style
}

func (u Username) Print(writer io.Writer, segments []Segment, current int) {
    uc, _ := user.Current()
    FormatString(writer, " " + uc.Username + " ", u.style, segments, current)
}

func (u Username) GetStyle(segments []Segment, current int) Style {
    return u.style
}

func NewUsername(style Style) Segment {
    return &Username{ style }
}