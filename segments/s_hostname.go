package segments

import (
	"github.com/RomainGiraud/gompt/format"
	"io"
	"os"
)

// Hostname segment prints machine hostname.
type Hostname struct {
	Style format.Style
}

// Create a Hostname segment.
func NewHostname() *Hostname {
	return NewHostnameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a Hostname segment with a style.
func NewHostnameStylized(style format.Style) *Hostname {
	return &Hostname{style}
}

func (s *Hostname) Load() {
}

func (s Hostname) Print(writer io.Writer, segments []Segment, current int) {
	h, err := os.Hostname()
	if err != nil {
		return
	}
	FormatString(writer, " "+h+" ", s.Style, segments, current)
}

func (s Hostname) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}

// FullHostname segment prints current hostname with options.
type FullHostname struct {
	Style     format.Style
	SshSymbol string
	SshColor  format.Color
}

// Create a FullHostname segment.
func NewFullHostname() *FullHostname {
	return NewFullHostnameStylized(format.NewStyleStandard(format.UniBrush{format.White}, format.UniBrush{format.Black}))
}

// Create a FullHostname segment with a style.
func NewFullHostnameStylized(style format.Style) *FullHostname {
	return &FullHostname{style, "\uf44c", format.Black}
}

func (s *FullHostname) Load() {
}

func (s FullHostname) Print(writer io.Writer, segments []Segment, current int) {
	h, err := os.Hostname()
	if err != nil {
		return
	}
	e := os.Getenv("SSH_CLIENT")

	ff := []PartFormatter{}
	ff = append(ff, PartFormatter{" ", nil, nil})
	if len(e) != 0 {
		ff = append(ff, PartFormatter{s.SshSymbol, s.SshColor, nil})
	}
	ff = append(ff, PartFormatter{h, nil, nil})
	ff = append(ff, PartFormatter{" ", nil, nil})
	FormatParts(writer, s.Style, segments, current, ff)
}

func (s FullHostname) GetStyle(segments []Segment, current int) format.Style {
	return s.Style
}
