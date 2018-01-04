package format

import (
	"io"
)

// Style is an interface that store a format.
// A snapshot can be extracted at a given percent.
type Style interface {
	ValueAt(t float32) StyleSnapshot
}

// StyleSnapshot represents an extract of a Style at a given percent.
type StyleSnapshot interface {
	// Format a string
	Format(io.Writer, string, StyleSnapshot, StyleSnapshot)
	// Return foreground
	GetFg() Color
	// Return background
	GetBg() Color
	// Replace foreground and/or background
	OverrideFgBg(Color, Color) StyleSnapshot
}

// Most basic style.
type StyleStandard struct {
	Fg        Brush
	Bg        Brush
	Bold      bool
	Italic    bool
	Underline bool
}

// Create a StyleStandard.
func NewStyleStandard(fg Brush, bg Brush) *StyleStandard {
	return &StyleStandard{fg, bg, false, false, false}
}

func (s StyleStandard) ValueAt(t float32) StyleSnapshot {
	var snapshot StyleSnapshotStandard
	if s.Fg != nil {
		snapshot.fg = s.Fg.ValueAt(t)
	}
	if s.Bg != nil {
		snapshot.bg = s.Bg.ValueAt(t)
	}
	if s.Bold {
		snapshot.attributes = append(snapshot.attributes, Bold)
	}
	if s.Italic {
		snapshot.attributes = append(snapshot.attributes, Italic)
	}
	if s.Underline {
		snapshot.attributes = append(snapshot.attributes, Underline)
	}
	return snapshot
}

type StyleSnapshotStandard struct {
	fg         Color
	bg         Color
	attributes []Attribute
}

func (s StyleSnapshotStandard) Format(writer io.Writer, str string, prev StyleSnapshot, next StyleSnapshot) {
	s.attributes = append(s.attributes, Bg(s.bg))
	s.attributes = append(s.attributes, Fg(s.fg))
	Fformat(writer, str, s.attributes...)
}

func (s StyleSnapshotStandard) GetFg() Color {
	return s.fg
}

func (s StyleSnapshotStandard) GetBg() Color {
	return s.bg
}

func (s StyleSnapshotStandard) OverrideFgBg(fg Color, bg Color) StyleSnapshot {
	var newSs StyleSnapshotStandard = s
	if fg != nil {
		newSs.fg = fg
	}
	if bg != nil {
		newSs.bg = bg
	}
	return newSs
}

// StyleChameleon can be used with Powerline Font separtors.
type StyleChameleon struct {
}

// Create a StyleChameleon.
func NewStyleChameleon() *StyleChameleon {
	return &StyleChameleon{}
}

func (s StyleChameleon) ValueAt(t float32) StyleSnapshot {
	return StyleSnapshotChameleon{}
}

type StyleSnapshotChameleon struct {
}

func (s StyleSnapshotChameleon) Format(writer io.Writer, str string, prev StyleSnapshot, next StyleSnapshot) {
	var fg Color = White
	if prev != nil {
		if c := prev.GetBg(); c != nil {
			fg = c
		}
	}

	var bg Color = Black
	if next != nil {
		if c := next.GetBg(); c != nil {
			bg = c
		}
	}

	Fformat(writer, str, Bg(bg), Fg(fg))
}

func (s StyleSnapshotChameleon) GetFg() Color {
	return nil
}

func (s StyleSnapshotChameleon) GetBg() Color {
	return nil
}

func (s StyleSnapshotChameleon) OverrideFgBg(fg Color, bg Color) StyleSnapshot {
	return s
}
