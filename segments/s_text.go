package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
	"strings"
)

// Text segment is used to print a string.
// Environment values are expanded:
//   ${HOME} or $HOME
// Commands are evaluated:
// 	 ${cmd> ls -l | wl -l}
// To display dollar value, double it:
//   $$
type Text struct {
	style format.Style
	value string
}

func (s Text) Load() []Segment {
	return []Segment{s}
}

func (s Text) Print(writer io.Writer, segments []Segment, current int) {
	text := os.Expand(s.value, getenv)
	FormatString(writer, text, s.style, segments, current)
}

func (s Text) GetStyle(segments []Segment, current int) format.Style {
	return s.style
}

func NewText(style format.Style, text string) *Text {
	return &Text{style, text}
}

func getenv(key string) string {
	// Escape $ by doubling it: $$
	if key == "$" {
		return "$"
	}

	// Execute "my_command" from ${cmd> my_command}
	if strings.HasPrefix(key, "cmd> ") {
		return ExecCommand("bash", "-c", key[5:])
	}

	return os.Getenv(key)
}
