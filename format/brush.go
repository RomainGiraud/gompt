package format

// Brush is an interface that represents a color mix.
// A single color value can be extracted with a percent value.
type Brush interface {
	ValueAt(float32) Color
}

// UniBrush contains a uniq color on all range.
type UniBrush struct {
	Color0 Color
}

func (b UniBrush) ValueAt(t float32) Color {
	return b.Color0
}

// GradientBrush contains gradient between two colors.
type GradientBrush struct {
	Color0 Color
	Color1 Color
}

func (b GradientBrush) ValueAt(t float32) Color {
	return b.Color0.Lerp(b.Color1, t)
}
