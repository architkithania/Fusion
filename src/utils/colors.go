// Use this colors file to develop a color palate for the application so that a consistent color theme
// can be developed throughout the application. THis also has the added benefit of caching the colors so
// that the same color does not need to be repeatedly created (Although the performance gain is negligible)
package utils

import "github.com/veandco/go-sdl2/sdl"

var (
	BLACK  = &sdl.Color{A: 1}
	WHITE  = &sdl.Color{R: 255, G: 255, B: 255, A: 1}
	GRAY   = &sdl.Color{R: 243, G: 241, B: 239, A: 1}
	GREEN  = &sdl.Color{R: 35, G: 203, B: 167, A: 1}
	SILVER = &sdl.Color{R: 191, G: 191, B: 191, A: 1}
	BRIGHT_GREEN = &sdl.Color{G: 230, B: 64, A: 1}
)
