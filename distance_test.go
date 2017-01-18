package chroma

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGBADistance(t *testing.T) {
	tests := []struct {
		given   color.RGBA
		nearest color.RGBA
		exp     RGBDistance
	}{
		{
			color.RGBA{0, 0, 0, 0},
			color.RGBA{0, 0, 0, 0},
			RGBDistance{0, 0, 0},
		},
		{
			color.RGBA{221, 18, 137, 0},
			color.RGBA{218, 24, 132, 0},
			RGBDistance{-3, 6, -5},
		},
		{
			color.RGBA{0, 0, 0, 0},
			color.RGBA{255, 255, 255, 255},
			RGBDistance{255, 255, 255},
		},
		{
			color.RGBA{255, 255, 255, 255},
			color.RGBA{0, 0, 0, 0},
			RGBDistance{-255, -255, -255},
		},
	}

	for _, test := range tests {
		d := rgbDistance(test.given, test.nearest)
		assert.NotNil(t, d)
		assert.Equal(t, test.exp.R, d.R, "Red distance incorrect")
		assert.Equal(t, test.exp.G, d.G, "Green distance incorrect")
		assert.Equal(t, test.exp.B, d.B, "Blue distance incorrect")
	}
}
