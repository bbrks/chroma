# chroma [![Go Reference](https://pkg.go.dev/badge/github.com/bbrks/chroma.svg)](https://pkg.go.dev/github.com/bbrks/chroma) [![GitHub tag](https://img.shields.io/github/tag/bbrks/chroma.svg)](https://github.com/bbrks/chroma/releases) [![license](https://img.shields.io/github/license/bbrks/chroma.svg)](https://github.com/bbrks/chroma/blob/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/bbrks/chroma)](https://goreportcard.com/report/github.com/bbrks/chroma) [![codecov](https://codecov.io/gh/bbrks/chroma/branch/master/graph/badge.svg)](https://codecov.io/gh/bbrks/chroma)

Find names of ~2800 colors from a Go package.

Inspiration and color data have been sourced from [ayushoriginal/Optimized-RGB-To-ColorName](https://github.com/ayushoriginal/Optimized-RGB-To-ColorName)

---

Usage:

```go
// Pick a color close to IBM Blue
blue := color.RGBA{84, 138, 196, 255}

// Find the nearest color
nearest := chroma.NearestColor(blue)

// Print what we got
fmt.Printf("%+v\n", nearest)
```

```
&{Color:{Color:{R:90 G:135 B:197 A:255} Names:[IBM Blue]} RGBDistance:{R:6 G:-3 B:1}}
```

## Todo

- Use faster algorithm for `NearestColor()`
- Optimize data structures.

## Contributing

Issues/PRs are much appreciated!

Feature requests/improvements welcome.

## License
This project is licensed under the [MIT License](LICENSE.md).
