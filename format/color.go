package format

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// Print escaped attribute.
func escapedPrint(w io.Writer, a ...interface{}) {
	fmt.Fprint(w, "\\[\\e[")
	fmt.Fprint(w, a...)
	fmt.Fprint(w, "m\\]")
}

// Color is the interface that represents any color type
type Color interface {
	Fprintf(io.Writer, bool)
	Lerp(Color, float32) Color
}

// Color4 represents a named color in 8 available.
type Color4 struct {
	value uint8
}

func (c Color4) Fprintf(w io.Writer, fg bool) {
	if fg {
		escapedPrint(w, 30+c.value)
	} else {
		escapedPrint(w, 40+c.value)
	}
}

func (c Color4) Lerp(other Color, t float32) Color {
	switch o := other.(type) {
	case Color4:
		return Color4{c.value + uint8(t*float32(int16(o.value)-int16(c.value)))}
	default:
		return c
	}
}

// Color8 represents a height bits color.
type Color8 struct {
	value uint8
}

func (c Color8) Fprintf(w io.Writer, fg bool) {
	if fg {
		escapedPrint(w, "38;5;", c.value)
	} else {
		escapedPrint(w, "48;5;", c.value)
	}
}

func (c Color8) Lerp(other Color, t float32) Color {
	switch o := other.(type) {
	case Color8:
		return Color8{c.value + uint8(t*float32(int16(o.value)-int16(c.value)))}
	default:
		return c
	}
}

// Color24 represents a true color in 24 bits format.
type Color24 struct {
	r, g, b uint8
}

func (c Color24) Fprintf(w io.Writer, fg bool) {
	if fg {
		escapedPrint(w, "38;2;", c.r, ";", c.g, ";", c.b)
	} else {
		escapedPrint(w, "48;2;", c.r, ";", c.g, ";", c.b)
	}
}

func (c Color24) Lerp(other Color, t float32) Color {
	switch o := other.(type) {
	case Color24:
		return Color24{
			c.r + uint8(t*float32(int16(o.r)-int16(c.r))),
			c.g + uint8(t*float32(int16(o.g)-int16(c.g))),
			c.b + uint8(t*float32(int16(o.b)-int16(c.b))),
		}
	default:
		return c
	}
}

// Named color
var (
	Black   = Color4{0}
	Red     = Color4{1}
	Green   = Color4{2}
	Yellow  = Color4{3}
	Blue    = Color4{4}
	Magenta = Color4{5}
	Cyan    = Color4{6}
	White   = Color4{7}
	Default = Color4{9}
)

// Construct a color from a string.
//   format.NewColor("black")
//   format.NewColor("250")
//   format.NewColor("#f0f")
//   format.NewColor("#ff00ff")
func NewColor(str string) Color {
	if i, err := strconv.Atoi(str); err == nil {
		return Color8{uint8(i)}
	}

	if str[0] == '#' {
		var r, g, b uint64
		var err error
		if len(str) == 4 {
			r, err = strconv.ParseUint(str[1:2]+str[1:2], 16, 8)
			g, err = strconv.ParseUint(str[2:3]+str[2:3], 16, 8)
			b, err = strconv.ParseUint(str[3:4]+str[3:4], 16, 8)
		} else if len(str) == 7 {
			r, err = strconv.ParseUint(str[1:3], 16, 8)
			g, err = strconv.ParseUint(str[3:5], 16, 8)
			b, err = strconv.ParseUint(str[5:7], 16, 8)
		}
		if err != nil {
			panic("Error during color parsing: " + err.Error())
		}
		return Color24{uint8(r), uint8(g), uint8(b)}
	}

	switch str {
	case "black":
		return Color4{0}
	case "red":
		return Color4{1}
	case "green":
		return Color4{2}
	case "yellow":
		return Color4{3}
	case "blue":
		return Color4{4}
	case "magenta":
		return Color4{5}
	case "cyan":
		return Color4{6}
	case "white":
		return Color4{7}
	case "default":
		return Color4{9}
	}

	panic("cannot parse color")
}

// Attribute is a type that represents a terminal ANSI escape code.
type Attribute func(io.Writer)

// Fformat formats the string surrounded by attributes and writes it to w.
func Fformat(w io.Writer, s string, a ...Attribute) {
	FformatFn(a...)(w, s)
}

// FformatFn returns a method to format a string.
func FformatFn(a ...Attribute) func(io.Writer, string) {
	return func(w io.Writer, s string) {
		defer Reset(w)
		for _, style := range a {
			style(w)
		}
		io.WriteString(w, s)
	}
}

// Format returns a formatted string surrounded by attributes.
func Format(s string, a ...Attribute) string {
	return FormatFn(a...)(s)
}

// FormatFn returns a method to format a string.
func FormatFn(a ...Attribute) func(string) string {
	return func(s string) string {
		var b bytes.Buffer
		FformatFn(a...)(&b, s)
		return b.String()
	}
}

// Set background color.
func Bg(c Color) Attribute {
	return func(w io.Writer) {
		if c == nil {
			return
		}
		c.Fprintf(w, false)
	}
}

// Set foreground color.
func Fg(c Color) Attribute {
	return func(w io.Writer) {
		if c == nil {
			return
		}
		c.Fprintf(w, true)
	}
}

// Reset all attributes.
func Reset(w io.Writer) {
	escapedPrint(w, 0)
}

// Set bold.
func Bold(w io.Writer) {
	escapedPrint(w, 1)
}

// Set faint.
func Faint(w io.Writer) {
	escapedPrint(w, 2)
}

// Set italic.
func Italic(w io.Writer) {
	escapedPrint(w, 3)
}

// Set underline.
func Underline(w io.Writer) {
	escapedPrint(w, 4)
}

// Set a slow blink.
func SlowBlink(w io.Writer) {
	escapedPrint(w, 5)
}

// Set a rapid blink.
func RapidBlink(w io.Writer) {
	escapedPrint(w, 6)
}

// Reverse foreground and background colors.
func Reverse(w io.Writer) {
	escapedPrint(w, 7)
}
