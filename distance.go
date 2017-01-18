package chroma

import "image/color"

// RGBDistance holds a signed distance of 8-bit RGB values
// to get from one color to another. Range: -255 through 255.
type RGBDistance struct {
	R, G, B int16
}

// rgbDistance calculates the distance from one color to the other.
func rgbDistance(a, b color.Color) *RGBDistance {
	r1, g1, b1 := a.(color.RGBA).R,
		a.(color.RGBA).G,
		a.(color.RGBA).B
	r2, g2, b2 := b.(color.RGBA).R,
		b.(color.RGBA).G,
		b.(color.RGBA).B

	return &RGBDistance{
		R: int16(r2) - int16(r1),
		G: int16(g2) - int16(g1),
		B: int16(b2) - int16(b1),
	}
}
