// Package chroma provides a way of finding names of ~2800 colors.
//
// These colors have been sourced from: https://github.com/ayushoriginal/Optimized-RGB-To-ColorName
package chroma

import "image/color"

// Nearest contains a Color, and distance from another color.
type Nearest struct {
	Color       Color
	RGBDistance RGBDistance
}

// NearestColor returns the nearest color, and distance from the given color.
func NearestColor(c color.Color) *Nearest {
	nearest := data.convert(c)
	if nearest == nil {
		return nil
	}

	distance := rgbDistance(c, nearest.Color)

	return &Nearest{
		*nearest,
		*distance,
	}
}
