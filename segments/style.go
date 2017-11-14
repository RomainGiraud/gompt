package segments

import(
    _ "fmt"
    "log"
    "strconv"
    "encoding/json"
)


type Style interface {
    Format(string, Context, int) string
    GetBg() Background
    GetFg() Foreground
}

func NewStyle(name string, options json.RawMessage) Style {
    var style Style
    switch name {
    case "uni":
        var conf StyleConfigUni
        json.Unmarshal(options, &conf)
        style = NewStyleUni(conf)
    case "prev-next":
        var conf StyleConfigChameleon
        json.Unmarshal(options, &conf)
        style = NewStyleChameleon(conf)
    default:
        style = nil
    }
    return style
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

type StyleConfigUni struct {
    Fg string `json:"fg"`
    Bg string `json:"bg"`
}

func NewStyleUni(config StyleConfigUni) Style {
    fg, _ := strconv.Atoi(config.Fg)
    bg, _ := strconv.Atoi(config.Bg)
    return &StyleUni{ Foreground(fg), Background(bg) }
}


type StyleChameleon struct {
    defaultFg Foreground
    defaultBg Background
}

func (s StyleChameleon) Format(str string, context Context, index int) string {
    if index == -1 {
        log.Panic("ERROR during style formatting")
    }

    prev, next := index - 1, -1
    if index + 1 < len(context.Order) {
        next = index + 1
    }

    fg := s.defaultFg
    if prev != -1 {
        tmp := context.Segments[context.Order[prev]].GetStyle().GetBg()
        fg = BgToFg(tmp)
    }

    bg := s.defaultBg
    if next != -1 {
        bg = context.Segments[context.Order[next]].GetStyle().GetBg()
    }

    return Colorize(str, Bg(bg), Fg(fg))
}

func (s StyleChameleon) GetBg() Background {
    return s.defaultBg
}

func (s StyleChameleon) GetFg() Foreground {
    return s.defaultFg
}

type StyleConfigChameleon struct {
    DefaultFg string `json:"default-fg,omitempty"`
    DefaultBg string `json:"default-bg,omitempty"`
}

func NewStyleChameleon(config StyleConfigChameleon) Style {
    fg, _ := strconv.Atoi(config.DefaultFg)
    bg, _ := strconv.Atoi(config.DefaultBg)
    return &StyleChameleon{ Foreground(fg), Background(bg) }
}