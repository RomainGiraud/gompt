package format

import (
	"bytes"
	"errors"
	"io"
	"strconv"
)

// Color is the interface that represents any color type
type Color interface {
	Fprintf(io.Writer, Shell, bool)
	Lerp(Color, float32) Color
}

// Color4 represents a named color in 8 available.
type Color4 int

const (
	Black Color4 = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White

	BrightBlack
	BrightRed
	BrightGreen
	BrightYellow
	BrightBlue
	BrightMagenta
	BrightCyan
	BrightWhite

	Default
)

func (c Color4) Fprintf(w io.Writer, sh Shell, fg bool) {
	sh.PrintColor4(w, c, fg)
}

func (c Color4) Lerp(other Color, t float32) Color {
	return c
}

// Color8 represents a height bits color.
type Color8 struct {
	value uint8
}

func (c Color8) Fprintf(w io.Writer, sh Shell, fg bool) {
	sh.PrintColor8(w, c, fg)
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

func (c Color24) Fprintf(w io.Writer, sh Shell, fg bool) {
	sh.PrintColor24(w, c, fg)
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
		} else {
			err = errors.New("invalid length for color parsing")
		}
		if err != nil {
			panic("Error during color parsing: " + err.Error())
		}
		return Color24{uint8(r), uint8(g), uint8(b)}
	}

	switch str {
	case "black":
		return Black
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	case "magenta":
		return Magenta
	case "cyan":
		return Cyan
	case "white":
		return White
	case "default":
		return Default
	}

	panic("cannot parse color")
}

// Attribute is a type that represents a terminal ANSI escape code.
type Attribute func(io.Writer, Shell)

// Fformat formats the string surrounded by attributes and writes it to w.
func Fformat(w io.Writer, sh Shell, s string, a ...Attribute) {
	FformatFn(a...)(w, sh, s)
}

// FformatFn returns a method to format a string.
func FformatFn(a ...Attribute) func(io.Writer, Shell, string) {
	return func(w io.Writer, sh Shell, s string) {
		defer Reset(w, sh)
		for _, style := range a {
			style(w, sh)
		}
		io.WriteString(w, s)
	}
}

// Format returns a formatted string surrounded by attributes.
func Format(s string, sh Shell, a ...Attribute) string {
	return FormatFn(a...)(s, sh)
}

// FormatFn returns a method to format a string.
func FormatFn(a ...Attribute) func(string, Shell) string {
	return func(s string, sh Shell) string {
		var b bytes.Buffer
		FformatFn(a...)(&b, sh, s)
		return b.String()
	}
}

// Set background color.
func Bg(c Color) Attribute {
	return func(w io.Writer, sh Shell) {
		if c == nil {
			return
		}
		c.Fprintf(w, sh, false)
	}
}

// Set foreground color.
func Fg(c Color) Attribute {
	return func(w io.Writer, sh Shell) {
		if c == nil {
			return
		}
		c.Fprintf(w, sh, true)
	}
}

// // Reset all attributes.
func Reset(w io.Writer, sh Shell) {
	sh.Reset(w)
}

// Set bold.
func Bold(w io.Writer, sh Shell) {
	sh.Bold(w)
}

// Set faint.
func Faint(w io.Writer, sh Shell) {
	sh.Faint(w)
}

// Set italic.
func Italic(w io.Writer, sh Shell) {
	sh.Italic(w)
}

// Set underline.
func Underline(w io.Writer, sh Shell) {
	sh.Underline(w)
}

// Set a slow blink.
func SlowBlink(w io.Writer, sh Shell) {
	sh.SlowBlink(w)
}

// Set a rapid blink.
func RapidBlink(w io.Writer, sh Shell) {
	sh.RapidBlink(w)
}

// Reverse foreground and background colors.
func Reverse(w io.Writer, sh Shell) {
	sh.Reverse(w)
}