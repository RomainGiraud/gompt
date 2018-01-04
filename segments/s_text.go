package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
	"strings"
)

func getenv(key string) string {
	// Escape $ by doubling it: $$
	if key == "$" {
		return "$"
	}

	// Execute "my_command" in ${cmd> my_command}
	if strings.HasPrefix(key, "cmd> ") {
		output, err := ExecCommand("bash", "-c", key[5:])
		if err != nil {
			return ""
		}
		return output
	}

	return os.Getenv(key)
}

// Text segment prints a stylized string.
type Text struct {
	Text  string
	Style format.Style
}

// Create a Text segment and interprets variables and commands.
// Environment values are expanded:
//   ${HOME} or $HOME
// Commands are evaluated:
// 	 ${cmd> ls -l | wl -l}
// To display dollar value, double it:
//   $$
func NewText(text string) *Text {
	return NewTextStylized(text, format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a Text segment with a style and interprets variables and commands.
func NewTextStylized(text string, style format.Style) *Text {
	return &Text{text, style}
}

func (s *Text) Load() {
}

func (s Text) Print(writer io.Writer, segments []Segment, current int) {
	text := os.Expand(s.Text, getenv)
	FormatString(writer, text, s.Style, segments, current)
}

func (s Text) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}
