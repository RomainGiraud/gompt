package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os/user"
)

// Username segment prints current username.
type Username struct {
	Style format.Style
}

// Create a Username segment.
func NewUsername() *Username {
	return NewUsernameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a Username segment with a style.
func NewUsernameStylized(style format.Style) *Username {
	return &Username{style}
}

func (s *Username) Load() {
}

func (s Username) Print(writer io.Writer, segments []Segment, current int) {
	u, err := user.Current()
	if err != nil {
		return
	}
	FormatString(writer, " "+u.Username+" ", s.Style, segments, current)
}

func (s Username) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}
