package segments

import(
    "io"
)


type Git struct {
    style Style
    branch string
}

func (s Git) Load() []Segment {
    s.branch = ExecCommand("git", "rev-parse", "--abbrev-ref", "HEAD")

    if len(s.branch) == 0 {
        return []Segment{}
    }

    return []Segment{ s }
}

func (s Git) Print(writer io.Writer, segments []Segment, current int) {
    FormatString(writer, " " + s.branch + " ", s.style, segments, current)
}

func (s Git) GetStyle(segments []Segment, current int) Style {
    return s.style
}

func NewGit(style Style) Segment {
    return &Git{ style, "" }
}