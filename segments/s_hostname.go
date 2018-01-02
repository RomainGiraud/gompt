package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"os"
)

// HostnameLoader create a segment containing the machine hostname.
type HostnameLoader struct {
	Style format.Style
}

func (s HostnameLoader) Load() []Segment {
	h, err := os.Hostname()
	if err != nil {
		return []Segment{}
	}
	return []Segment{Text{s.Style, " " + h + " "}}
}
