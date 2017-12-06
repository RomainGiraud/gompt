package segments

import(
    _ "fmt"
    _ "log"
    "errors"
    _ "strconv"
    _ "encoding/json"
)


type Style interface {
    Format(string, Context, int, float32) string
    GetBg() Color
    GetFg() Color
}


func init() {
    RegisterStyleLoader("uni", LoadStyleUni)
    RegisterStyleLoader("cham", LoadStyleChameleon)
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

func (s StyleUni) Format(str string, context Context, index int, t float32) string {
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

func (s StyleChameleon) Format(str string, context Context, index int, t float32) string {
    prev, next := index - 1, -1
    if index + 1 < len(context.Segments) {
        next = index + 1
    }

    fg := NewColor("default")
    if prev != -1 {
        fg = context.Segments[prev].GetStyle(context, index).GetBg()
    }

    bg := NewColor("default")
    if next != -1 {
        bg = context.Segments[next].GetStyle(context, index).GetBg()
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
