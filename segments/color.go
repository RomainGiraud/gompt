package segments

import(
    "io"
    "fmt"
    "strconv"
)


const escape = "\x1b"

type Brush interface {
    ValueAt(float32) Color
}

type UniBrush struct {
    Color0 Color
}

func (b UniBrush) ValueAt(t float32) Color {
    return b.Color0
}

type GradientBrush struct {
    Color0 Color
    Color1 Color
}

func (b GradientBrush) ValueAt(t float32) Color {
    return b.Color0.Lerp(b.Color1, t)
}


type Color interface {
    Fprintf(writer io.Writer, fg bool)
    Lerp(Color, float32) Color
}


type Color4 struct {
    value uint8
}

func (c Color4) Fprintf(writer io.Writer, fg bool) {
    if fg {
        fmt.Fprintf(writer, "%s[%dm", escape, 30 + c.value)
    } else {
        fmt.Fprintf(writer, "%s[%dm", escape, 40 + c.value)
    }
}

func (c Color4) Lerp(other Color, t float32) Color  {
    switch o := other.(type) {
    case Color4:
        return Color4{ c.value + uint8(t * float32(int16(o.value) - int16(c.value))) }
    default:
        return c
    }
}

 
type Color8 struct {
    value uint8
}

func (c Color8) Fprintf(writer io.Writer, fg bool) {
    if fg {
        fmt.Fprintf(writer, "%s[38;5;%dm", escape, c.value)
    } else {
        fmt.Fprintf(writer, "%s[48;5;%dm", escape, c.value)
    }
}

func (c Color8) Lerp(other Color, t float32) Color {
    switch o := other.(type) {
    case Color8:
        return Color8{ c.value + uint8(t * float32(int16(o.value) - int16(c.value))) }
    default:
        return c
    }
}


type Color24 struct {
    r, g, b uint8
}

func (c Color24) Fprintf(writer io.Writer, fg bool) {
    if fg {
        fmt.Fprintf(writer, "%s[38;2;%d;%d;%dm", escape, c.r, c.g, c.b)
    } else {
        fmt.Fprintf(writer, "%s[48;2;%d;%d;%dm", escape, c.r, c.g, c.b)
    }
}

func (c Color24) Lerp(other Color, t float32) Color {
    switch o := other.(type) {
    case Color24:
        return Color24{
            c.r + uint8(t * float32(int16(o.r) - int16(c.r))),
            c.g + uint8(t * float32(int16(o.g) - int16(c.g))),
            c.b + uint8(t * float32(int16(o.b) - int16(c.b))),
        }
    default:
        return c
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
            r, _ = strconv.ParseUint(str[1:2] + str[1:2], 16, 8)
            g, _ = strconv.ParseUint(str[2:3] + str[2:3], 16, 8)
            b, _ = strconv.ParseUint(str[3:4] + str[3:4], 16, 8)
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


type Attribute func(io.Writer) error

func Colorize(writer io.Writer, str string, styles ...Attribute) {
    ColorizeFn(styles...)(writer, str)
}

func ColorizeFn(styles ...Attribute) func(io.Writer, string) {
    return func(writer io.Writer, str string) {
        defer Reset(writer)
        for _, style := range styles {
            err := style(writer)
            if err != nil {
                return
            }
        }
        io.WriteString(writer, str)
    }
}

func Bg(c Color) Attribute {
    return func(writer io.Writer) error {
        c.Fprintf(writer, false)
        return nil
    }
}

func Fg(c Color) Attribute {
    return func(writer io.Writer) error {
        c.Fprintf(writer, true)
        return nil
    }
}

func Bold(writer io.Writer) error {
    fmt.Fprintf(writer, "%s[1m", escape)
    return nil
}

func Underline(writer io.Writer) error {
    fmt.Fprintf(writer, "%s[4m", escape)
    return nil
}

func Reset(writer io.Writer) error {
    fmt.Fprintf(writer, "%s[0m", escape)
    return nil
}
