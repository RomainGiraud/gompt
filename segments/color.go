package segments

import(
    "fmt"
    "bytes"
    "strconv"
)

//fmt.Println(Colorize("toto", Bg24(0, 155, 0), Fg(30)))
//fmt.Println(fmt.Sprintf(Color.Bg(46).Fg(30)("toto")))

const escape = "\x1b"


type Color interface {
    Fprintf(buffer *bytes.Buffer, fg bool)
}


type Color4 struct {
    value uint8
}

func (c Color4) Fprintf(buffer *bytes.Buffer, fg bool) {
    if fg {
        fmt.Fprintf(buffer, "%s[%dm", escape, 30 + c.value)
    } else {
        fmt.Fprintf(buffer, "%s[%dm", escape, 40 + c.value)
    }
}


type Color8 struct {
    value uint8
}

func (c Color8) Fprintf(buffer *bytes.Buffer, fg bool) {
    if fg {
        fmt.Fprintf(buffer, "%s[38;5;%dm", escape, c.value)
    } else {
        fmt.Fprintf(buffer, "%s[48;5;%dm", escape, c.value)
    }
}


type Color24 struct {
    r, g, b uint8
}

func (c Color24) Fprintf(buffer *bytes.Buffer, fg bool) {
    if fg {
        fmt.Fprintf(buffer, "%s[38;2;%d;%d;%dm", escape, c.r, c.g, c.b)
    } else {
        fmt.Fprintf(buffer, "%s[48;2;%d;%d;%dm", escape, c.r, c.g, c.b)
    }
}

var Black    = Color4{ 0 }
var Red      = Color4{ 1 }
var Green    = Color4{ 2 }
var Yellow   = Color4{ 3 }
var Blue     = Color4{ 4 }
var Magenta  = Color4{ 5 }
var Cyan     = Color4{ 6 }
var White    = Color4{ 7 }
var Default  = Color4{ 9 }

func NewColor(str string) Color {
    if i, err := strconv.Atoi(str); err == nil {
        return Color8{ uint8(i) }
    }

    if str[0] == '#' {
        var r, g, b uint64
        if len(str) == 4 {
            r, _ = strconv.ParseUint(str[1:2], 16, 8)
            g, _ = strconv.ParseUint(str[2:3], 16, 8)
            b, _ = strconv.ParseUint(str[3:4], 16, 8)
        } else if len(str) == 7 {
            r, _ = strconv.ParseUint(str[1:3], 16, 8)
            g, _ = strconv.ParseUint(str[3:5], 16, 8)
            b, _ = strconv.ParseUint(str[5:7], 16, 8)
        }
        return Color24{ uint8(r), uint8(g), uint8(b) }
    }

    switch str {
    case "black"    : return Color4{ 0 }
    case "red"      : return Color4{ 1 }
    case "green"    : return Color4{ 2 }
    case "yellow"   : return Color4{ 3 }
    case "blue"     : return Color4{ 4 }
    case "magenta"  : return Color4{ 5 }
    case "cyan"     : return Color4{ 6 }
    case "white"    : return Color4{ 7 }
    case "default"  : return Color4{ 9 }
    }

    panic("cannot parse color");
}


type Attribute func(*bytes.Buffer) error

func Colorize(str string, styles ...Attribute) string {
    return ColorizeFn(styles...)(str)
}

func ColorizeFn(styles ...Attribute) func(str string) string {
    return func(str string) string {
        var buffer bytes.Buffer
        for _, style := range styles {
            err := style(&buffer)
            if err != nil {
                return ""
            }
        }
        buffer.WriteString(str)
        Reset(&buffer)
        return buffer.String()
    }
}

func Bg(c Color) Attribute {
    return func(buffer *bytes.Buffer) error {
        c.Fprintf(buffer, false)
        return nil
    }
}

func Fg(c Color) Attribute {
    return func(buffer *bytes.Buffer) error {
        c.Fprintf(buffer, true)
        return nil
    }
}

func Bold(buffer *bytes.Buffer) error {
    fmt.Fprintf(buffer, "%s[1m", escape)
    return nil
}

func Underline(buffer *bytes.Buffer) error {
    fmt.Fprintf(buffer, "%s[4m", escape)
    return nil
}

func Reset(buffer *bytes.Buffer) error {
    fmt.Fprintf(buffer, "%s[0m", escape)
    return nil
}



/*
type Color struct {
    Background
}

func (c *Color) Bg(bg Background) *Color {
    return c
}

//return fmt.Sprintf("\x1b[%d;47m%s\x1b[0m", 30, dir)
func Bg(bg Background) (func(string) string) {
    return func(str string) string {
        return fmt.Sprintf("%s[%dm%s%s[0m", escape, bg, str, escape)
    }
}

func Fg(fg Foreground) (func(string) string) {
    return func(str string) string {
        return fmt.Sprintf("%s[%dm%s%s[0m", escape, fg, str, escape)
    }
}
*/
