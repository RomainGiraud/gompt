package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os/user"
)

type Username struct {
	style format.Style
}

func (s Username) Load() []Segment {
	return []Segment{s}
}

func (s Username) Print(writer io.Writer, segments []Segment, current int) {
	u, _ := user.Current()
	FormatString(writer, " "+u.Username+" ", s.style, segments, current)
}

func (s Username) GetStyle(segments []Segment, current int) format.Style {
	return s.style
}

func NewUsername(style format.Style) *Username {
	return &Username{style}
}
