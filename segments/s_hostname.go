package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
)

// Hostname segment prints machine hostname.
type Hostname struct {
	Style format.Style
}

// Create a Hostname segment.
func NewHostname() *Hostname {
	return NewHostnameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a Hostname segment with a style.
func NewHostnameStylized(style format.Style) *Hostname {
	return &Hostname{style}
}

func (s *Hostname) Load() {
}

func (s Hostname) Print(writer io.Writer, segments []Segment, current int) {
	h, err := os.Hostname()
	if err != nil {
		return
	}
	FormatString(writer, " "+h+" ", s.Style, segments, current)
}

func (s Hostname) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}
