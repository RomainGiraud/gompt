package segments

import(
    "fmt"
    _ "log"
    "errors"
    _ "strconv"
    _ "encoding/json"
)

type Style interface {
    Format(string, float32, Style, Style) string
    GetBg() Color
    GetFg() Color
}

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

func FormatStringArray(strs []string, separator string, style Style, segments []Segment, current int) {
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
        fmt.Print(style.Format(separator, float32(i) / size, prevStyle, nextStyle))
    }
}


func init() {
    RegisterStyleLoader("uni", LoadStyleUni)
    RegisterStyleLoader("cham", LoadStyleChameleon)
    RegisterStyleLoader("gradient", LoadStyleGradient)
}

type StyleLoader func(map[string]interface{}) Style

var styleLoaders = map[string]StyleLoader{}

func RegisterStyleLoader(name string, fn StyleLoader) {
    styleLoaders[name] = fn
}

func LoadStyle(conf interface{}) (Style, error) {
    config, ok := conf.(map[string]interface{})
    if ! ok {
        return nil, errors.New("LoadStyle: cannot parse configuration")
    }

    typeName, ok := config["type"].(string);
    if ! ok {
        return nil, errors.New("LoadStyle: key 'type' does not exists in configuration")
    }

    val, ok := styleLoaders[typeName];
    if ! ok {
        panic("unknown style type: " + typeName)
        return nil, errors.New("unknown style type: " + typeName)
    }

    return val(config), nil
}

type StyleUni struct {
    fg Color
    bg Color
}

func (s StyleUni) Format(str string, t float32, prevStyle Style, nextStyle Style) string {
    return Colorize(str, Bg(s.bg), Fg(s.fg))
}

func (s StyleUni) GetBg() Color {
    return s.bg
}

func (s StyleUni) GetFg() Color {
    return s.fg
}

func NewStyleUni(fg Color, bg Color) Style {
    return &StyleUni{ fg, bg }
}

func LoadStyleUni(config map[string]interface{}) Style {
    var fg, _ = config["fg"].(string)
    var bg, _ = config["bg"].(string)
    return &StyleUni{ NewColor(fg), NewColor(bg) }
}


type StyleChameleon struct {
    defaultFg Color
    defaultBg Color
}

func (s StyleChameleon) Format(str string, t float32, prevStyle Style, nextStyle Style) string {
    fg := NewColor("default")
    if prevStyle != nil {
        fg = prevStyle.GetBg()
    }

    bg := NewColor("default")
    if nextStyle != nil {
        bg = prevStyle.GetBg()
    }

    return Colorize(str, Bg(bg), Fg(fg))
}

func (s StyleChameleon) GetBg() Color {
    return s.defaultBg
}

func (s StyleChameleon) GetFg() Color {
    return s.defaultFg
}

func NewStyleChameleon() Style {
    return &StyleChameleon{ NewColor("default"), NewColor("default") }
}

func LoadStyleChameleon(config map[string]interface{}) Style {
    var fg, _ = config["default-fg"].(string)
    var bg, _ = config["default-bg"].(string)
    return &StyleChameleon{ NewColor(fg), NewColor(bg) }
}


type StyleGradient struct {
    fgStart Color
    fgEnd   Color
    bgStart Color
    bgEnd   Color
}

func (s StyleGradient) Format(str string, t float32, prevStyle Style, nextStyle Style) string {
    return Colorize(str, Bg(s.bgStart.Lerp(s.bgEnd, t)), Fg(s.fgStart.Lerp(s.fgEnd, t)))
}

func (s StyleGradient) GetBg() Color {
    return s.bgStart
}

func (s StyleGradient) GetFg() Color {
    return s.fgEnd
}

func NewStyleGradient(fgStart Color, fgEnd Color, bgStart Color, bgEnd Color) Style {
    return &StyleGradient{ fgStart, fgEnd, bgStart, bgEnd }
}

func LoadStyleGradient(config map[string]interface{}) Style {
    var fgStart, _ = config["fg-start"].(string)
    var fgEnd  , _ = config["fg-end"].(string)
    var bgStart, _ = config["bg-start"].(string)
    var bgEnd  , _ = config["bg-end"].(string)
    return &StyleGradient{ NewColor(fgStart), NewColor(fgEnd), NewColor(bgStart), NewColor(bgEnd) }
}