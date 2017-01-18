package chroma

import "image/color"

// Palette is a palette of colors.
type Palette []Color

// Color embeds a color.Color interface and a slice of associated names.
type Color struct {
	color.Color
	Names []string
}

// Load will replace the current palette data with the palette specified.
func (p Palette) Load() {
	data = p
}

// convert returns the palette color closest to c in Euclidean R,G,B space.
func (p Palette) convert(c color.Color) *Color {
	if len(p) == 0 {
		return nil
	}
	return &p[p.index(c)]
}

// index returns the index of the palette color closest to c in Euclidean
// R,G,B,A space.
func (p Palette) index(c color.Color) int {
	// A batch version of this computation is in image/draw/draw.go.

	cr, cg, cb, ca := c.RGBA()
	ret, bestSum := 0, uint32(1<<32-1)
	for i, v := range p {
		vr, vg, vb, va := v.Color.RGBA()
		sum := sqDiff(cr, vr) + sqDiff(cg, vg) + sqDiff(cb, vb) + sqDiff(ca, va)
		if sum < bestSum {
			if sum == 0 {
				return i
			}
			ret, bestSum = i, sum
		}
	}
	return ret
}

// sqDiff returns the squared-difference of x and y, shifted by 2 so that
// adding four of those won't overflow a uint32.
//
// x and y are both assumed to be in the range [0, 0xffff].
func sqDiff(x, y uint32) uint32 {
	var d uint32
	if x > y {
		d = x - y
	} else {
		d = y - x
	}
	return (d * d) >> 2
}
