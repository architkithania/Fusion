// Some commonly used utility functions are defined in this file
package utils

import "github.com/veandco/go-sdl2/sdl"

// Centers an image on a surface of dimensions surfaceX by surfaceY
func CenterTexture(texture *sdl.Texture, surfaceX, surfaceY int32) *sdl.Rect {
	x, y := surfaceX/2, surfaceY/2
	_, _, w, h, _ := texture.Query()

	return &sdl.Rect{
		X: x - (w/2),
		Y: y - (h/2),
		W: w,
		H: h,
	}
}

func GetCenterCoordinates(width, height , surfaceX, surfaceY int32) (centerX, centerY int32) {
	x, y := surfaceX/2, surfaceY/2

	return x - (width/2), y - (height/2)
}

func Percent(value int32, percent float32) float32 {
	return float32(value) * (percent/100)
}
