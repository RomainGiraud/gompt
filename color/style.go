package color

import(
    "strconv"
)


type StyleConfig struct {
    Type string `json:"type"`
    Fg string `json:"fg"`
    Bg string `json:"bg"`
}

type StyleFmt func(string) string

func (s StyleConfig) GetFmt() StyleFmt {
    fg, _ := strconv.Atoi(s.Fg)
    bg, _ := strconv.Atoi(s.Bg)
    return ColorizeFn(Bg(Background(bg)), Fg(Foreground(fg)))
}
