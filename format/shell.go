package format

import (
	"io"
	"fmt"
)

type Shell interface {
	PrintColor4(io.Writer, Color4, bool)
	PrintColor8(io.Writer, Color8, bool)
	PrintColor24(io.Writer, Color24, bool)
	Reset(io.Writer)
	Bold(w io.Writer)
	Faint(w io.Writer)
	Italic(w io.Writer)
	Underline(w io.Writer)
	SlowBlink(w io.Writer)
	RapidBlink(w io.Writer)
	Reverse(w io.Writer)
}

type Bash struct {
}

// Print escaped attribute.
func (b Bash) EscapedPrint(w io.Writer, a ...interface{}) {
	fmt.Fprint(w, "\\[\\e[")
	fmt.Fprint(w, a...)
	fmt.Fprint(w, "m\\]")
}

func (b Bash) PrintColor4(w io.Writer, c Color4, fg bool) {
	var offset = 0
	if (!fg) {
		offset = 10
	}

	switch (c) {
		case Default:	b.EscapedPrint(w, 39 + offset);
		case Black:		b.EscapedPrint(w, 30 + offset);
		case Red:		b.EscapedPrint(w, 31 + offset);
		case Green:		b.EscapedPrint(w, 32 + offset);
		case Yellow:	b.EscapedPrint(w, 33 + offset);
		case Blue:		b.EscapedPrint(w, 34 + offset);
		case Magenta:	b.EscapedPrint(w, 35 + offset);
		case Cyan:		b.EscapedPrint(w, 36 + offset);
		case White:		b.EscapedPrint(w, 37 + offset);
		case BrightBlack:		b.EscapedPrint(w, 90 + offset);
		case BrightRed:			b.EscapedPrint(w, 91 + offset);
		case BrightGreen:		b.EscapedPrint(w, 92 + offset);
		case BrightYellow:		b.EscapedPrint(w, 93 + offset);
		case BrightBlue:		b.EscapedPrint(w, 94 + offset);
		case BrightMagenta:		b.EscapedPrint(w, 95 + offset);
		case BrightCyan:		b.EscapedPrint(w, 96 + offset);
		case BrightWhite:		b.EscapedPrint(w, 97 + offset);
	}
}

func (b Bash) PrintColor8(w io.Writer, c Color8, fg bool) {
	if fg {
		b.EscapedPrint(w, "38;5;", c.value)
	} else {
		b.EscapedPrint(w, "48;5;", c.value)
	}
}

func (b Bash) PrintColor24(w io.Writer, c Color24, fg bool) {
	if fg {
		b.EscapedPrint(w, "38;2;", c.r, ";", c.g, ";", c.b)
	} else {
		b.EscapedPrint(w, "48;2;", c.r, ";", c.g, ";", c.b)
	}
}

func (b Bash) Reset(w io.Writer) {
	b.EscapedPrint(w, 0)
}

func (b Bash) Bold(w io.Writer) {
	b.EscapedPrint(w, 1)
}

func (b Bash) Faint(w io.Writer) {
	b.EscapedPrint(w, 2)
}

func (b Bash) Italic(w io.Writer) {
	b.EscapedPrint(w, 3)
}

func (b Bash) Underline(w io.Writer) {
	b.EscapedPrint(w, 4)
}

func (b Bash) SlowBlink(w io.Writer) {
	b.EscapedPrint(w, 5)
}

func (b Bash) RapidBlink(w io.Writer) {
	b.EscapedPrint(w, 6)
}

func (b Bash) Reverse(w io.Writer) {
	b.EscapedPrint(w, 7)
}

type Zsh struct {
}

func (z Zsh) EscapedPrint(w io.Writer, a ...interface{}) {
	fmt.Fprint(w, "\\[\\e[")
	fmt.Fprint(w, a...)
	fmt.Fprint(w, "m\\]")
}

func (z Zsh) PrintColor4(w io.Writer, c Color4, fg bool) {
	var ground = "F"
	if (!fg) {
		ground = "K"
	}

	switch (c) {
		case Default:	fmt.Fprint(w, "%", ground, "{default}")
		case Black:		fmt.Fprint(w, "%", ground, "{black}")
		case Red:		fmt.Fprint(w, "%", ground, "{red}")
		case Green:		fmt.Fprint(w, "%", ground, "{green}")
		case Yellow:	fmt.Fprint(w, "%", ground, "{yellow}")
		case Blue:		fmt.Fprint(w, "%", ground, "{blue}")
		case Magenta:	fmt.Fprint(w, "%", ground, "{magenta}")
		case Cyan:		fmt.Fprint(w, "%", ground, "{cyan}")
		case White:		fmt.Fprint(w, "%", ground, "{white}")
		case BrightBlack:		fmt.Fprint(w, "%", ground, "{black}")
		case BrightRed:			fmt.Fprint(w, "%", ground, "{red}")
		case BrightGreen:		fmt.Fprint(w, "%", ground, "{green}")
		case BrightYellow:		fmt.Fprint(w, "%", ground, "{yellow}")
		case BrightBlue:		fmt.Fprint(w, "%", ground, "{blue}")
		case BrightMagenta:		fmt.Fprint(w, "%", ground, "{magenta}")
		case BrightCyan:		fmt.Fprint(w, "%", ground, "{cyan}")
		case BrightWhite:		fmt.Fprint(w, "%", ground, "{white}")
	}
}

func (z Zsh) PrintColor8(w io.Writer, c Color8, fg bool) {
	var ground = "F"
	if (!fg) {
		ground = "K"
	}
	fmt.Fprint(w, "%", ground, "{", c.value, "}")
}

func (z Zsh) PrintColor24(w io.Writer, c Color24, fg bool) {
	var ground = "F"
	if (!fg) {
		ground = "K"
	}
	fmt.Fprint(w, "%", ground, "{#")
	fmt.Fprintf(w, "%02x%02x%02x", c.r, c.g, c.b)
	fmt.Fprint(w, "}")
}

func (z Zsh) Reset(w io.Writer) {
	fmt.Fprint(w, "%f%k%b%u")
}

func (z Zsh) Bold(w io.Writer) {
	fmt.Fprint(w, "%B")
}

func (z Zsh) Faint(w io.Writer) {
}

func (z Zsh) Italic(w io.Writer) {
}

func (z Zsh) Underline(w io.Writer) {
	fmt.Fprint(w, "%U")
}

func (z Zsh) SlowBlink(w io.Writer) {
}

func (z Zsh) RapidBlink(w io.Writer) {
}

func (z Zsh) Reverse(w io.Writer) {
}