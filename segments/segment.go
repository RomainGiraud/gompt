package segments

import(
    "io"
)


type Arguments struct {
    Status int
    ConfigPath string
}

type Segment interface {
    Print(io.Writer, []Segment, int)
    GetStyle([]Segment, int) Style
}

type SegmentList []Segment

func (segments *SegmentList) Display(writer io.Writer) {
    if len(*segments) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(*segments); i, j = i+1, j+1 {
        seg := (*segments)[i]
        seg.Print(writer, *segments, i)
    }
}