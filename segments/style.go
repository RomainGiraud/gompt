package segments

import (
	"io"
	"unicode/utf8"
)

func FormatString(writer io.Writer, str string, style Style, segments []Segment, current int) {
	size := float32(len(str) - 1)

	var prevStyle, nextStyle StyleSnapshot = nil, nil

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

type PartFormatter struct {
	str string
	fg  Color
	bg  Color
}

func FormatParts(writer io.Writer, style Style, segments []Segment, current int, strs []PartFormatter) {
	sizeMax := 0
	for _, s := range strs {
		sizeMax += utf8.RuneCountInString(s.str)
	}

	var prevStyle, nextStyle StyleSnapshot = nil, nil

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

func FormatStringArrayBlock(writer io.Writer, strs []string, style Style, separator string, separatorStyle Style, segments []Segment, current int) {
	var prevStyle, nextStyle StyleSnapshot = nil, nil

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

type Style interface {
	ValueAt(t float32) StyleSnapshot
}

type StyleSnapshot interface {
	Format(io.Writer, string, StyleSnapshot, StyleSnapshot)
	GetFg() Color
	GetBg() Color
	OverrideFgBg(Color, Color) StyleSnapshot
}

type StyleStandard struct {
	Fg        Brush
	Bg        Brush
	Bold      bool
	Italic    bool
	Underline bool
}

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
	Colorize(writer, str, s.attributes...)
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

type StyleChameleon struct {
}

func NewStyleChameleon(fg Brush, bg Brush) *StyleChameleon {
	return &StyleChameleon{}
}

func (s StyleChameleon) ValueAt(t float32) StyleSnapshot {
	return StyleSnapshotChameleon{}
}

type StyleSnapshotChameleon struct {
}

func (s StyleSnapshotChameleon) Format(writer io.Writer, str string, prev StyleSnapshot, next StyleSnapshot) {
	fg := NewColor("default")
	if prev != nil {
		if c := prev.GetBg(); c != nil {
			fg = c
		}
	}

	bg := NewColor("default")
	if next != nil {
		if c := next.GetBg(); c != nil {
			bg = c
		}
	}

	Colorize(writer, str, Bg(bg), Fg(fg))
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
