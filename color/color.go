package color

import(
    "fmt"
    "bytes"
)

const escape = "\x1b"

type Background int
type Foreground int
type Decorator int
type Attribute func(*bytes.Buffer) error

type Color struct {
    element fmt.Stringer
    styles []Attribute
}

func (c Color) String() string {
    return ColorizeFn(c.styles...)(c.element.String())
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
