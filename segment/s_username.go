package segment

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

// FullUsername segment prints current username with options.
type FullUsername struct {
	Style     format.Style
	RootStyle format.Style
	username  string
}

// Create a FullUsername segment.
func NewFullUsername() *FullUsername {
	return NewFullUsernameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a FullUsername segment with a style.
func NewFullUsernameStylized(style format.Style) *FullUsername {
	return &FullUsername{style, format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}), ""}
}

func (s *FullUsername) Load() {
	u, err := user.Current()
	if err != nil {
		return
	}
	s.username = u.Username
}

func (s FullUsername) Print(writer io.Writer, segments []Segment, current int) {
	FormatString(writer, " "+s.username+" ", s.GetStyle(segments, current), segments, current)
}

func (s FullUsername) GetStyle(segments []Segment, current int) format.Style {
	if s.username == "root" {
		return s.RootStyle
	}
	return s.Style
}
