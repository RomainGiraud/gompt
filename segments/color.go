package segments

import(
    "fmt"
    "bytes"
)

//fmt.Println(Colorize("toto", Bg24(0, 155, 0), Fg(30)))
//fmt.Println(fmt.Sprintf(Color.Bg(46).Fg(30)("toto")))


const escape = "\x1b"

type Background int
type Foreground int
type Decorator int
type Attribute func(*bytes.Buffer) error

const(
    FgBlack     Foreground = iota + 30
    FgRed       Foreground = iota + 31
    FgGreen     Foreground = iota + 32
    FgYellow    Foreground = iota + 33
    FgBlue      Foreground = iota + 34
    FgMagenta   Foreground = iota + 35
    FgCyan      Foreground = iota + 36
    FgWhite     Foreground = iota + 37
    FgDefault   Foreground = iota + 39
)

const(
    BgBlack     Background = iota + 40
    BgRed       Background = iota + 41
    BgGreen     Background = iota + 42
    BgYellow    Background = iota + 43
    BgBlue      Background = iota + 44
    BgMagenta   Background = iota + 45
    BgCyan      Background = iota + 46
    BgWhite     Background = iota + 47
    BgDefault   Background = iota + 49
)

func FgToBg(fg Foreground) Background {
    switch fg {
    case FgBlack    : return BgBlack;
    case FgRed      : return BgRed;
    case FgGreen    : return BgGreen;
    case FgYellow   : return BgYellow;
    case FgBlue     : return BgBlue;
    case FgMagenta  : return BgMagenta;
    case FgCyan     : return BgCyan;
    case FgWhite    : return BgWhite;
    case FgDefault  : return BgDefault;
    default         : return BgDefault;
    }
}

func BgToFg(bg Background) Foreground {
    switch bg {
    case BgBlack    : return FgBlack;
    case BgRed      : return FgRed;
    case BgGreen    : return FgGreen;
    case BgYellow   : return FgYellow;
    case BgBlue     : return FgBlue;
    case BgMagenta  : return FgMagenta;
    case BgCyan     : return FgCyan;
    case BgWhite    : return FgWhite;
    case BgDefault  : return FgDefault;
    default         : return FgDefault;
    }
}

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

func Bg(bg Background) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[%dm", escape, bg)
        return nil
    }
}

func Bg8(bg int) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[48;5;%dm", escape, bg)
        return nil
    }
}

func Bg24(r int, g int, b int) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[48;2;%d;%d;%dm", escape, r, g, b)
        return nil
    }
}

func Fg(fg Foreground) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[%dm", escape, fg)
        return nil
    }
}

func Fg8(fg int) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[38;5;%dm", escape, fg)
        return nil
    }
}

func Fg24(r int, g int, b int) Attribute {
    return func(buffer *bytes.Buffer) error {
        fmt.Fprintf(buffer, "%s[38;2;%d;%d;%dm", escape, r, g, b)
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
