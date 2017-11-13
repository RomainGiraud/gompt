package color

import(
    "strconv"
    "encoding/json"
)


type StyleFmt func(string) string

type Style interface {
    GetFmt() StyleFmt
}

func NewStyle(name string, options json.RawMessage) Style {
    var style Style
    switch name {
    case "uni":
        var tmp StyleConfigUni
        json.Unmarshal(options, &tmp)
        style = tmp
    case "prev-next":
        var tmp StyleConfigChameleon
        json.Unmarshal(options, &tmp)
        style = tmp
    default:
        style = nil
    }
    return style
}

type StyleConfigUni struct {
    Fg string `json:"fg"`
    Bg string `json:"bg"`
}

func (s StyleConfigUni) GetFmt() StyleFmt {
    fg, _ := strconv.Atoi(s.Fg)
    bg, _ := strconv.Atoi(s.Bg)
    return ColorizeFn(Bg(Background(bg)), Fg(Foreground(fg)))
}

type StyleConfigChameleon struct {
    DefaultFg string `json:"default-fg"`
    DefaultBg string `json:"default-bg"`
}

func (s StyleConfigChameleon) GetFmt() StyleFmt {
    fg, _ := strconv.Atoi(s.DefaultFg)
    bg, _ := strconv.Atoi(s.DefaultBg)
    return ColorizeFn(Bg(Background(bg)), Fg(Foreground(fg)))
}
