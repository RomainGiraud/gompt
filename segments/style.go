package segments

import(
    "strconv"
    "encoding/json"
)


type Style interface {
    Format(string, Context, string) string
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

func (s StyleConfigUni) Format(str string, context Context, name string) string {
    fg, _ := strconv.Atoi(s.Fg)
    bg, _ := strconv.Atoi(s.Bg)
    return Colorize(str, Bg(Background(bg)), Fg(Foreground(fg)))
}

type StyleConfigChameleon struct {
    DefaultFg string `json:"default-fg"`
    DefaultBg string `json:"default-bg"`
}

func (s StyleConfigChameleon) Format(str string, context Context, name string) string {
    fg, _ := strconv.Atoi(s.DefaultFg)
    bg, _ := strconv.Atoi(s.DefaultBg)
    return Colorize(str, Bg(Background(bg)), Fg(Foreground(fg)))
}
