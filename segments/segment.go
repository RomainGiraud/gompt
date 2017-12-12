package segments

import(
    "fmt"
)


type Arguments struct {
    Status int
    ConfigPath string
}

type Segment interface {
    Print([]Segment, int)
    GetStyle([]Segment, int) Style
}

type SegmentList []Segment

func (segments *SegmentList) Display() {
    if len(*segments) == 0 {
        panic("Empty prompt")
    }

    for i, j := 0, 1; i < len(*segments); i, j = i+1, j+1 {
        seg := (*segments)[i]
        seg.Print(*segments, i)
    }
    fmt.Printf("\n")
}