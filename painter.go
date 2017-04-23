package mochi

import "image/color"

type Painter interface {
	PaintOptions() PaintOptions
}

type PaintOptions struct {
	Alpha           float64
	BackgroundColor color.Color
	BorderColor     color.Color
	BorderWidth     float64
	CornerRadius    float64
	ShadowOpacity   float64
	ShadowRadius    float64
	ShadowOffset    float64
	ShadowColor     color.Color
	// Transform?
	// Mask
}

func (p PaintOptions) PaintOptions() PaintOptions {
	return p
}
