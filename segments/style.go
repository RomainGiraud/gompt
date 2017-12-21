package segments

import(
    "io"
    "unicode/utf8"
)


func FormatString(writer io.Writer, str string, style Style, segments []Segment, current int) {
    size := float32(len(str) - 1)

    var prevStyle, nextStyle StyleSnapshot = nil, nil

    if current != 0 {
        prevStyle = segments[current - 1].GetStyle(segments, current - 1).ValueAt(1)
    }

    if (current + 1) < len(segments) {
        nextStyle = segments[current + 1].GetStyle(segments, current + 1).ValueAt(0)
    }

    for i, s := range str {
        style.ValueAt(float32(i) / size).Format(writer, string(s), prevStyle, nextStyle)
    }
}

type PartFormatter struct {
    str string
    fg Color
    bg Color
}

func FormatParts(writer io.Writer, style Style, segments []Segment, current int, strs []PartFormatter) {
    sizeMax := 0
    for _, s := range strs {
        sizeMax += utf8.RuneCountInString(s.str)
    }

    var prevStyle, nextStyle StyleSnapshot = nil, nil

    if current != 0 {
        prevStyle = segments[current - 1].GetStyle(segments, current - 1).ValueAt(1)
    }

    if (current + 1) < len(segments) {
        nextStyle = segments[current + 1].GetStyle(segments, current + 1).ValueAt(0)
    }

    i := 0
    for _, s := range strs {
        for _, c := range s.str {
            style.ValueAt(float32(i) / float32(sizeMax)).Override(s.fg, s.bg).Format(writer, string(c), prevStyle, nextStyle)
            i += 1
        }
    }
}

func FormatStringArrayBlock(writer io.Writer, strs []string, style Style, separator string, separatorStyle Style, segments []Segment, current int) {
    var prevStyle, nextStyle StyleSnapshot = nil, nil

    if current != 0 {
        prevStyle = segments[current - 1].GetStyle(segments, current - 1).ValueAt(1)
    }

    if (current + 1) < len(segments) {
        nextStyle = segments[current + 1].GetStyle(segments, current + 1).ValueAt(0)
    }

    size := float32(len(strs) - 1)
    for i, s := range strs {
        idx := float32(i)

        currentStyle := style.ValueAt(idx / size)
        currentStyle.Format(writer, s, prevStyle, nextStyle)

        if (int(idx) + 1) < len(strs) {
            separatorStyle.ValueAt(0).Format(writer, separator, currentStyle, style.ValueAt((idx + 1) / size))
        }
    }
}


type Style interface {
    ValueAt(t float32) StyleSnapshot
}

type StyleSnapshot interface {
    Format(io.Writer, string, StyleSnapshot, StyleSnapshot)
    GetFg() Color
    GetBg() Color
    Override(Color, Color) StyleSnapshot
}


type StyleStandard struct {
    Fg Brush
    Bg Brush
}

type StyleSnapshotStandard struct {
    fg Color
    bg Color
}

func (s StyleStandard) ValueAt(t float32) StyleSnapshot {
    var snapshot StyleSnapshotStandard
    if s.Fg != nil {
        snapshot.fg = s.Fg.ValueAt(t)
    }
    if s.Bg != nil {
        snapshot.bg = s.Bg.ValueAt(t)
    }
    return snapshot
}

func (s StyleSnapshotStandard) Format(writer io.Writer, str string, prev StyleSnapshot, next StyleSnapshot) {
    Colorize(writer, str, Bg(s.bg), Fg(s.fg))
}

func (s StyleSnapshotStandard) GetFg() Color {
    return s.fg
}

func (s StyleSnapshotStandard) GetBg() Color {
    return s.bg
}

func (s StyleSnapshotStandard) Override(fg Color, bg Color) StyleSnapshot {
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

type StyleSnapshotChameleon struct {
}

func (s StyleChameleon) ValueAt(t float32) StyleSnapshot {
    var snapshot StyleSnapshotChameleon
    return snapshot
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

func (s StyleSnapshotChameleon) Override(fg Color, bg Color) StyleSnapshot {
    return s
}