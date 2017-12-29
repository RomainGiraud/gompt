package format

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
