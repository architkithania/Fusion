package rectbutton

import (
	"Fusion/src/components/text"
	"Fusion/src/utils"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// An implementation of the button interface. This is the most basic form of the button, providing
// a rectangle with text on top.
type RectangularButton struct {
	// The text that would be displayed on top of the button. Currently the text would
	// always be black color but changes can be made to Draw function in order to change
	// this behavior
	BtnText string

	// Size attributes of the button
	Width  int32
	Height int32

	// Position where the button should be rendered on the screen
	X int32
	Y int32

	// Color of the button background (the enclosing rectangle of the button)
	Color *sdl.Color
	Font  *ttf.Font

	// The callback function that gets called when the button is clicked. Note that the function
	// isn't directly called by the EventManager but rather through the RunCallback method
	CallBack func(...interface{}) error
}

// Provided Constructor
func New(text string, width, height int32, color *sdl.Color, font *ttf.Font) *RectangularButton {
	button := &RectangularButton{
		BtnText: text,
		Width:   width,
		Height:  height,
		Color:   color,
		Font:    font,
	}

	return button
}

// Method used to draw the button on the screen through the use of the provided
// sdl.Renderer
func (btn *RectangularButton) Draw(x, y int32, renderer *sdl.Renderer) error {
	rect := sdl.Rect{
		X: x,
		Y: y,
		W: btn.Width,
		H: btn.Height,
	}

	_ = renderer.SetDrawColor(btn.Color.R, btn.Color.G, btn.Color.B, btn.Color.A)
	_ = renderer.FillRect(&rect)
	textTexture, _ := text.New(btn.BtnText, btn.Font, renderer, sdl.Color{})
	defer textTexture.Destroy()

	_, _, tW, tH, _ := textTexture.Query()
	cenX, cenY := utils.GetCenterCoordinates(tW, tH, btn.Width, btn.Height)

	textRect := &sdl.Rect{
		X: x + cenX,
		Y: y + cenY,
		W: tW,
		H: tH,
	}

	btn.X = textRect.X
	btn.Y = textRect.Y

	return renderer.Copy(textTexture, nil, textRect)
}

// Getters and setters required by the ClickEvent interface
func (btn *RectangularButton) GetX() int32 {
	return btn.X
}

func (btn *RectangularButton) GetY() int32 {
	return btn.Y
}

func (btn *RectangularButton) GetWidth() int32 {
	return btn.Width
}

func (btn *RectangularButton) GetHeight() int32 {
	return btn.Height
}

func (btn *RectangularButton) RunCallback(i ...interface{}) error {
	return btn.CallBack(i)
}
