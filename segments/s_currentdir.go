package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"log"
	"os"
	"strings"
)

// CurrentDir segment display current working directory.
type currentDir struct {
	style       format.Style
	styleUnit   format.Style
	isPlain     bool
	separator   string
	fgSeparator format.Color
	directories []string
}

func (s currentDir) Print(writer io.Writer, segments []Segment, current int) {
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

func (s currentDir) GetStyle(segments []Segment, current int) format.Style {
	if !s.isPlain && len(s.directories) == 1 {
		return s.styleUnit
	}
	return s.style
}

// Create a CurrentDir segment with a unique style for all path.
// Separators are simple string.
type CurrentDirPlainLoader struct {
	Style       format.Style
	StyleUnit   format.Style
	Separator   string
	FgSeparator format.Color
	MaxDepth    uint
	Ellipsis    string
}

func (s CurrentDirPlainLoader) Load() []Segment {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	home_dir := os.Getenv("HOME")
	if strings.Index(dir, home_dir) == 0 {
		dir = strings.Replace(dir, home_dir, "~", 1)
	}

	directories := strings.Split(dir, "/")
	if s.MaxDepth != 0 && len(directories) > int(s.MaxDepth) {
		directories = directories[len(directories)-int(s.MaxDepth):]
		directories[0] = s.Ellipsis
	}

	return []Segment{currentDir{s.Style, s.StyleUnit, true, s.Separator, s.FgSeparator, directories}}
}

// Create a CurrentDir segment by spliting each directory and apply a style to it.
// Separators are transition characters.
type CurrentDirBlockLoader struct {
	Style     format.Style
	StyleUnit format.Style
	Separator string
	MaxDepth  uint
	Ellipsis  string
}

func (s CurrentDirBlockLoader) Load() []Segment {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	home_dir := os.Getenv("HOME")
	if strings.Index(dir, home_dir) == 0 {
		dir = strings.Replace(dir, home_dir, "~", 1)
	}

	directories := strings.Split(dir, "/")
	if s.MaxDepth != 0 && len(directories) > int(s.MaxDepth) {
		directories = directories[len(directories)-int(s.MaxDepth):]
		directories[0] = s.Ellipsis
	}

	return []Segment{currentDir{s.Style, s.StyleUnit, false, s.Separator, nil, directories}}
}
