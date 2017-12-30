package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
)

// A SegmentLoader is an interface that loads zero, one or more segments.
// A loader can check the presence of an environment (for example a repository).
type SegmentLoader interface {
	Load() []Segment
}

// A list of segment loaders.
type LoaderList []SegmentLoader

// Load segments from loaders.
func (loaders LoaderList) Load() SegmentList {
	if len(loaders) == 0 {
		panic("Empty prompt")
	}

	var segments SegmentList
	for i, j := 0, 1; i < len(loaders); i, j = i+1, j+1 {
		loader := loaders[i]
		segments = append(segments, loader.Load()...)
	}
	return segments
}

// A segment is a displayed element of the prompt.
type Segment interface {
	Print(io.Writer, []Segment, int)
	GetStyle([]Segment, int) format.Style
}

// A list of segments.
type SegmentList []Segment

// Render SegmentList
func (segments SegmentList) Render(writer io.Writer) {
	for i, j := 0, 1; i < len(segments); i, j = i+1, j+1 {
		seg := (segments)[i]
		seg.Print(writer, segments, i)
	}
}
