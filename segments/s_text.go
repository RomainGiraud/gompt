package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
	"strings"
)

// TextLoader loads a Text segment and interprets some characters.
// Environment values are expanded:
//   ${HOME} or $HOME
// Commands are evaluated:
// 	 ${cmd> ls -l | wl -l}
// To display dollar value, double it:
//   $$
type TextLoader struct {
	Style format.Style
	Value string
}

func (s TextLoader) Load() []Segment {
	text := os.Expand(s.Value, getenv)
	return []Segment{Text{s.Style, text}}
}

func getenv(key string) string {
	// Escape $ by doubling it: $$
	if key == "$" {
		return "$"
	}

	// Execute "my_command" from ${cmd> my_command}
	if strings.HasPrefix(key, "cmd> ") {
		output, err := ExecCommand("bash", "-c", key[5:])
		if err != nil {
			return ""
		}
		return output
	}

	return os.Getenv(key)
}

// Text segment is used to print a stylized string.
type Text struct {
	style format.Style
	text  string
}

func (s Text) Print(writer io.Writer, segments []Segment, current int) {
	FormatString(writer, s.text, s.style, segments, current)
}

func (s Text) GetStyle(segments []Segment, current int) format.Style {
	return s.style
}
