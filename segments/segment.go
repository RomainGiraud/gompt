package segments

import (
	"io"
)

type Arguments struct {
	Status     int
	ConfigPath string
}

type Segment interface {
	Load() []Segment
	Print(io.Writer, []Segment, int)
	GetStyle([]Segment, int) Style
}

type SegmentList []Segment

func (segments *SegmentList) Render(writer io.Writer) {
	if len(*segments) == 0 {
		panic("Empty prompt")
	}

	var toRender SegmentList
	for i, j := 0, 1; i < len(*segments); i, j = i+1, j+1 {
		seg := (*segments)[i]
		toRender = append(toRender, seg.Load()...)
	}

	for i, j := 0, 1; i < len(toRender); i, j = i+1, j+1 {
		seg := toRender[i]
		seg.Print(writer, toRender, i)
	}
}
