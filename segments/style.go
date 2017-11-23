package segments

import(
    _ "fmt"
    _ "log"
    "errors"
    _ "strconv"
    _ "encoding/json"
)


type Style interface {
    Format(string, Context, int) string
    GetBg() Background
    GetFg() Foreground
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
    fg Foreground
    bg Background
}

func (s StyleUni) Format(str string, context Context, index int) string {
    return Colorize(str, Bg(s.bg), Fg(s.fg))
}

func (s StyleUni) GetBg() Background {
    return s.bg
}

func (s StyleUni) GetFg() Foreground {
    return s.fg
}

func NewStyleUni(fg Foreground, bg Background) Style {
    return &StyleUni{ fg, bg }
}

func LoadStyleUni(config map[string]interface{}) Style {
    var fg, _ = config["fg"].(string)
    var bg, _ = config["bg"].(string)
    return &StyleUni{ StrToFg(fg), StrToBg(bg) }
}


type StyleChameleon struct {
    defaultFg Foreground
    defaultBg Background
}

func (s StyleChameleon) Format(str string, context Context, index int) string {
    prev, next := index - 1, -1
    if index + 1 < len(context.Segments) {
        next = index + 1
    }

    fg := FgDefault
    if prev != -1 {
        tmp := context.Segments[prev].GetStyle(context, index).GetBg()
        fg = BgToFg(tmp)
    }

    bg := BgDefault
    if next != -1 {
        bg = context.Segments[next].GetStyle(context, index).GetBg()
    }

    return Colorize(str, Bg(bg), Fg(fg))
}

func (s StyleChameleon) GetBg() Background {
    return s.defaultBg
}

func (s StyleChameleon) GetFg() Foreground {
    return s.defaultFg
}

func NewStyleChameleon() Style {
    return &StyleChameleon{ FgDefault, BgDefault }
}

func LoadStyleChameleon(config map[string]interface{}) Style {
    var fg, _ = config["default-fg"].(string)
    var bg, _ = config["default-bg"].(string)
    return &StyleChameleon{ StrToFg(fg), StrToBg(bg) }
}
