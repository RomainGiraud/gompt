package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"log"
	"os"
	"strings"
)

// CurrentDir segment display current working directory.
type CurrentDir struct {
	style       format.Style
	styleUnit   format.Style
	isPlain     bool
	separator   string
	fgSeparator format.Color
	maxDepth    uint
	ellipsis    string
	directories []string
}

func (s CurrentDir) Load() []Segment {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	home_dir := os.Getenv("HOME")
	if strings.Index(dir, home_dir) == 0 {
		dir = strings.Replace(dir, home_dir, "~", 1)
	}

	s.directories = strings.Split(dir, "/")
	if s.maxDepth != 0 && len(s.directories) > int(s.maxDepth) {
		s.directories = s.directories[len(s.directories)-int(s.maxDepth):]
		s.directories[0] = s.ellipsis
	}

	return []Segment{s}
}

func (s CurrentDir) Print(writer io.Writer, segments []Segment, current int) {
	for i, v := range s.directories {
		s.directories[i] = " " + v + " "
	}

	if s.isPlain {
		ff := []PartFormatter{}
		for i := 0; i < len(s.directories)-1; i += 1 {
			ff = append(ff, PartFormatter{s.directories[i], nil, nil})
			ff = append(ff, PartFormatter{s.separator, s.fgSeparator, nil})
		}
		ff = append(ff, PartFormatter{s.directories[len(s.directories)-1], nil, nil})
		FormatParts(writer, s.style, segments, current, ff)
	} else {
		FormatStringArrayBlock(writer, s.directories, s.style, s.separator, format.StyleChameleon{}, segments, current)
	}
}

func (s CurrentDir) GetStyle(segments []Segment, current int) format.Style {
	if !s.isPlain && len(s.directories) == 1 {
		return s.styleUnit
	}
	return s.style
}

// Create a CurrentDir segment with a unique style for all path.
// Separators are simple string.
func NewCurrentDirPlain(style format.Style, styleUnit format.Style, separator string, fgSeparator format.Color, maxDepth uint, ellipsis string) *CurrentDir {
	return &CurrentDir{style, styleUnit, true, separator, fgSeparator, maxDepth, ellipsis, []string{}}
}

// Create a CurrentDir segment by spliting each directory and apply a style to it.
// Separators are transition characters.
func NewCurrentDirSplitted(style format.Style, styleUnit format.Style, separator string, maxDepth uint, ellipsis string) *CurrentDir {
	return &CurrentDir{style, styleUnit, false, separator, nil, maxDepth, ellipsis, []string{}}
}
