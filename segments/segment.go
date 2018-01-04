package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
)

// A segment is a displayed element of the prompt.
type Segment interface {
	Load()
	Print(io.Writer, []Segment, int)
	GetStyle([]Segment, int) format.Style
}

// A list of segments.
type SegmentList []Segment

// Render SegmentList
func (segments SegmentList) Render(writer io.Writer) {
	// load all segments
	for _, s := range segments {
		s.Load()
	}

	// print all segments
	for i, j := 0, 1; i < len(segments); i, j = i+1, j+1 {
		seg := (segments)[i]
		seg.Print(writer, segments, i)
	}
}
