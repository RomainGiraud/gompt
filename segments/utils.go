package segments

import (
	"bytes"
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os/exec"
	"strings"
	"unicode/utf8"
)

// Execute a command and return standard output.
func ExecCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return strings.Trim(out.String(), "\n"), err
}

// Format a string with a style.
func FormatString(writer io.Writer, str string, style format.Style, segments []Segment, current int) {
	size := float32(len(str) - 1)

	var prevStyle, nextStyle format.StyleSnapshot = nil, nil

	if current != 0 {
		prevStyle = segments[current-1].GetStyle(segments, current-1).ValueAt(1)
	}

	if (current + 1) < len(segments) {
		nextStyle = segments[current+1].GetStyle(segments, current+1).ValueAt(0)
	}

	for i, s := range str {
		style.ValueAt(float32(i)/size).Format(writer, string(s), prevStyle, nextStyle)
	}
}

// Used to format elements with different styles.
type PartFormatter struct {
	str string
	fg  format.Color
	bg  format.Color
}

// Format a string with PartFormatter.
func FormatParts(writer io.Writer, style format.Style, segments []Segment, current int, strs []PartFormatter) {
	sizeMax := 0
	for _, s := range strs {
		sizeMax += utf8.RuneCountInString(s.str)
	}

	var prevStyle, nextStyle format.StyleSnapshot = nil, nil

	if current != 0 {
		prevStyle = segments[current-1].GetStyle(segments, current-1).ValueAt(1)
	}

	if (current + 1) < len(segments) {
		nextStyle = segments[current+1].GetStyle(segments, current+1).ValueAt(0)
	}

	i := 0
	for _, s := range strs {
		for _, c := range s.str {
			style.ValueAt(float32(i)/float32(sizeMax)).OverrideFgBg(s.fg, s.bg).Format(writer, string(c), prevStyle, nextStyle)
			i += 1
		}
	}
}

// Format a list of strings.
func FormatStringArrayBlock(writer io.Writer, strs []string, style format.Style, separator string, separatorStyle format.Style, segments []Segment, current int) {
	var prevStyle, nextStyle format.StyleSnapshot = nil, nil

	if current != 0 {
		prevStyle = segments[current-1].GetStyle(segments, current-1).ValueAt(1)
	}

	if (current + 1) < len(segments) {
		nextStyle = segments[current+1].GetStyle(segments, current+1).ValueAt(0)
	}

	size := float32(len(strs) - 1)
	for i := 0; i < len(strs)-1; i += 1 {
		idx := float32(i)

		currentStyle := style.ValueAt(idx / size)
		currentStyle.Format(writer, strs[i], prevStyle, nextStyle)

		separatorStyle.ValueAt(0).Format(writer, separator, currentStyle, style.ValueAt((idx+1)/size))
	}

	style.ValueAt(1).Format(writer, strs[len(strs)-1], prevStyle, nextStyle)
}
