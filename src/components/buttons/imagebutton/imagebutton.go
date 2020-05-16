package imagebutton

import "github.com/veandco/go-sdl2/sdl"

// An implementation of the button interface. It provides an image that can run a callback upon being
// clicked
type ImageButton struct {
	// Check rectbutton.RectangularButton for more details on attributes
	Width  int32
	Height int32
	X      int32
	Y      int32

	imageTexture *sdl.Texture

	CallBack func(...interface{}) error
}

// Provided Constructor
func New(image *sdl.Texture) *ImageButton {
	_, _, imageW, imageH, _ := image.Query()
	imageBtn := ImageButton{
		Width:        imageW,
		Height:       imageH,
		X:            0,
		Y:            0,
		imageTexture: image,
		CallBack:     nil,
	}

	return &imageBtn
}

func (btn *ImageButton) Draw(x, y int32, renderer *sdl.Renderer) error {
	btn.X, btn.Y = x, y

	rect := sdl.Rect{
		X: btn.X,
		Y: btn.Y,
		W: btn.Width,
		H: btn.Height,
	}

	return renderer.Copy(btn.imageTexture, nil, &rect)
}

// Getters and Setters required by the ClickEvent interface
func (btn *ImageButton) GetX() int32 {
	return btn.X
}

func (btn *ImageButton) GetY() int32 {
	return btn.Y
}

func (btn *ImageButton) GetWidth() int32 {
	return btn.Width
}

func (btn *ImageButton) GetHeight() int32 {
	return btn.Height
}

func (btn *ImageButton) RunCallback(i ...interface{}) error {
	return btn.CallBack(i)
}
