package chroma

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNearestColor(t *testing.T) {
	tests := []struct {
		color   color.Color
		nearest Color
	}{
		{
			color.RGBA{1, 1, 1, 255},
			Color{
				Color: color.RGBA{0, 0, 0, 255},
				Names: []string{"black", "gray0"},
			},
		},
		{
			color.RGBA{90, 135, 197, 255},
			Color{
				Color: color.RGBA{90, 135, 197, 255},
				Names: []string{"IBM Blue"},
			},
		},
		{
			color.RGBA{180, 25, 15, 255},
			Color{
				Color: color.RGBA{181, 31, 17, 255},
				Names: []string{"Red FS 21105"},
			},
		},
	}

	for _, test := range tests {
		nearest := NearestColor(test.color)
		assert.NotNil(t, nearest)
		assert.Equal(t, test.nearest.Names, nearest.Color.Names)
		assert.Equal(t, test.nearest, nearest.Color)
	}

	// Copy palette so we can re-assign later
	paletteCopy := data

	// Test an empty palette
	Palette{}.Load()
	nearest := NearestColor(tests[0].color)
	assert.Nil(t, nearest)

	// Re-assign palette
	paletteCopy.Load()
}

func ExampleNearestColor() {
	// Pick a color close to IBM Blue
	blue := color.RGBA{84, 138, 196, 255}

	// Find the nearest color
	nearest := NearestColor(blue)

	// Print what we got
	fmt.Printf("%+v\n", nearest)
	// Output: &{Color:{Color:{R:90 G:135 B:197 A:255} Names:[IBM Blue]} RGBDistance:{R:6 G:-3 B:1}}
}

// We need to benchmark different colours, as the current implementation
// is slice-order dependent, and performs differently depending on the colour.
func benchmarkNearestColor(c color.Color, b *testing.B) {
	for i := 0; i < b.N; i++ {
		NearestColor(c)
	}
}

func BenchmarkNearestColorWhite(b *testing.B) {
	benchmarkNearestColor(color.RGBA{255, 255, 255, 255}, b)
}

func BenchmarkNearestColorGrey(b *testing.B) {
	benchmarkNearestColor(color.RGBA{155, 155, 155, 255}, b)
}

func BenchmarkNearestColorBlack(b *testing.B) {
	benchmarkNearestColor(color.RGBA{0, 0, 0, 255}, b)
}
