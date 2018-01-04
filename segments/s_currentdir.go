package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"log"
	"os"
	"strings"
)

// CurrentDirBlock segment prints current working directory with splitted folders.
type CurrentDirBlock struct {
	Style       format.Style
	StyleUnit   format.Style
	Separator   string
	MaxDepth    uint
	Ellipsis    string
	directories []string
}

// Create a CurrentDirBlock segment by splitting each directory and apply a style to it.
func NewCurrentDirBlock() *CurrentDirBlock {
	return &CurrentDirBlock{
		format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}),
		format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}),
		">", 0, "\u2026",
		[]string{}}
}

func (s *CurrentDirBlock) Load() {
	s.directories = getCwdEllipsed(s.MaxDepth, s.Ellipsis)
}

func (s CurrentDirBlock) Print(writer io.Writer, segments []Segment, current int) {
	FormatStringArrayBlock(writer, s.directories, s.GetStyle(segments, current), s.Separator, format.StyleChameleon{}, segments, current)
}

func (s CurrentDirBlock) GetStyle(segments []Segment, current int) format.Style {
	if len(s.directories) == 1 {
		return s.StyleUnit
	}
	return s.Style
}

// CurrentDirPlain segment prints current working directory.
type CurrentDirPlain struct {
	Style       format.Style
	Separator   string
	FgSeparator format.Color
	MaxDepth    uint
	Ellipsis    string
	directories []string
}

// Create a CurrentDirPlain segment with a single style for all folders.
func NewCurrentDirPlain() *CurrentDirPlain {
	return &CurrentDirPlain{
		format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}),
		">", format.White,
		0, "\u2026",
		[]string{}}
}

func (s *CurrentDirPlain) Load() {
	s.directories = getCwdEllipsed(s.MaxDepth, s.Ellipsis)
}

func (s CurrentDirPlain) Print(writer io.Writer, segments []Segment, current int) {
	ff := []PartFormatter{}
	for i := 0; i < len(s.directories)-1; i += 1 {
		ff = append(ff, PartFormatter{s.directories[i], nil, nil})
		ff = append(ff, PartFormatter{s.Separator, s.FgSeparator, nil})
	}
	ff = append(ff, PartFormatter{s.directories[len(s.directories)-1], nil, nil})
	FormatParts(writer, s.Style, segments, current, ff)
}

func (s CurrentDirPlain) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}

func getCwdEllipsed(maxDepth uint, ellipsis string) []string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	home_dir := os.Getenv("HOME")
	if strings.Index(dir, home_dir) == 0 {
		dir = strings.Replace(dir, home_dir, "~", 1)
	}

	directories := strings.Split(dir, "/")
	if maxDepth != 0 && len(directories) > int(maxDepth) {
		directories = directories[len(directories)-int(maxDepth):]
		directories[0] = ellipsis
	}

	for i, v := range directories {
		directories[i] = " " + v + " "
	}

	return directories
}
