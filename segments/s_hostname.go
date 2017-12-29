package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
)

type Hostname struct {
	style format.Style
}

func (s Hostname) Load() []Segment {
	return []Segment{s}
}

func (s Hostname) Print(writer io.Writer, segments []Segment, current int) {
	h, _ := os.Hostname()
	FormatString(writer, " "+h+" ", s.style, segments, current)
}

func (s Hostname) GetStyle(segments []Segment, current int) format.Style {
	return s.style
}

func NewHostname(style format.Style) *Hostname {
	return &Hostname{style}
}
