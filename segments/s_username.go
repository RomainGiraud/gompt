package segments

import(
    "io"
    "os/user"
)


type Username struct {
    style Style
}

func (s Username) Print(writer io.Writer, segments []Segment, current int) {
    uc, _ := user.Current()
    FormatString(writer, " " + uc.Username + " ", s.style, segments, current)
}

func (s Username) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewUsername(style Style) Segment {
    return &Username{ style }
}