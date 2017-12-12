package segments

import(
    "fmt"
)


func FormatString(str string, style Style, segments []Segment, current int) {
    size := float32(len(str))
    for i, s := range str {
        var prevStyle, nextStyle Style = nil, nil
        if current != 0 {
            prevStyle = segments[current - 1].GetStyle(segments, current - 1)
        }
        if (current + 1) < len(segments) {
            nextStyle = segments[current + 1].GetStyle(segments, current + 1)
        }
        fmt.Print(style.Format(string(s), float32(i) / size, prevStyle, nextStyle))
    }
}

func FormatStringArray(strs []string, style Style, separator string, separatorStyle Brush, segments []Segment, current int) {
    size := float32(len(strs))
    for i, s := range strs {
        var prevStyle, nextStyle Style = nil, nil
        if current != 0 {
            prevStyle = segments[current - 1].GetStyle(segments, current - 1)
        }
        if (current + 1) < len(segments) {
            nextStyle = segments[current + 1].GetStyle(segments, current + 1)
        }
        fmt.Print(style.Format(s, float32(i) / size, prevStyle, nextStyle))
        fmt.Print(style.Override(separatorStyle, nil).Format(separator, float32(i) / size, prevStyle, nextStyle))
    }
}


type Style interface {
    Format(string, float32, Style, Style) string
    GetBg() Brush
    GetFg() Brush
    Override(Brush, Brush) Style
}


type StyleStandard struct {
    Fg Brush
    Bg Brush
}

func (s StyleStandard) GetBg() Brush {
    return s.Bg
}

func (s StyleStandard) GetFg() Brush {
    return s.Fg
}

func (s StyleStandard) Override(fg Brush, bg Brush) Style {
    var newStyle StyleStandard = s
    if fg != nil {
        newStyle.Fg = fg
    }
    if bg != nil {
        newStyle.Bg = bg
    }
    return newStyle
}

func (s StyleStandard) Format(str string, t float32, prevStyle Style, nextStyle Style) string {
    return Colorize(str, Bg(s.Bg.ValueAt(t)), Fg(s.Fg.ValueAt(t)))
}


type StyleChameleon struct {
}

func (s StyleChameleon) GetBg() Brush {
    return nil
}

func (s StyleChameleon) GetFg() Brush {
    return nil
}

func (s StyleChameleon) Override(fg Brush, bg Brush) Style {
    return s
}

func (s StyleChameleon) Format(str string, t float32, prevStyle Style, nextStyle Style) string {
    fg := NewColor("default")
    if prevStyle != nil {
        if style := prevStyle.GetBg(); style != nil {
            fg = style.ValueAt(1)
        }
    }

    bg := NewColor("default")
    if nextStyle != nil {
        if style := nextStyle.GetBg(); style != nil {
            bg = style.ValueAt(0)
        }
    }

    return Colorize(str, Bg(bg), Fg(fg))
}