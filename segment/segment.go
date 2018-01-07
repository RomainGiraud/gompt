package segment

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
)

// A segment is a displayed element of the prompt.
// It is advisable to create a segment in a New* function.
type Segment interface {
	// Initialize elements of segment needed for Print and GetStyle methods
	Load()
	// Print segment i in w. All segments are in s.
	Print(w io.Writer, s []Segment, i int)
	// Return style of segment i. All segments are in s.
	GetStyle(s []Segment, i int) format.Style
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
