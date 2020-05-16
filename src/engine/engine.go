// The main game engine for the game. It jobs include the following:
// - Provide ways to initialise the SDL library
// - Loads the different asset managers and the event manager
// - Unloads the different asset managers when they are no longer being used
package engine

import (
	"Fusion/src/managers/eventmanager"
	"Fusion/src/managers/fontmanager"
	"Fusion/src/managers/imgmanager"
	"Fusion/src/screens"
	"path/filepath"
	"runtime"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Basic constants used for initialisation. Note that these constants do not impact the game screen
// on the android platform, but are instead used to set the view window for development.
//
// For example: Set these values to 1920, 1080 respectively for developing for 1920x1080p resolutions.
// this will provide a more realistic view of the app while developing
var (
	winTitle  string
	winWidth  int32
	winHeight int32
)

type Engine struct {
	State int

	// The main window object of the application
	Window *sdl.Window

	// The main renderer for the application. For simplicity, only one renderer
	// is used throughout the entire application.
	Renderer *sdl.Renderer

	// The image manager for the application. Refer to src/managers/imgmanager/imgmanager.go for more info
	Image *imgmanager.ImageManager

	// The font manager for the application. Refer to src/managers/fontmanager/fontmanager.go for more info
	Font *fontmanager.FontManager

	// The event manager for the application. The event managers are divided by screen as the event manager
	// use a simple linear scan to fire events. Separating the events by different screens allows for some
	// optimization and responsiveness in the application.
	//
	//Refer to src/managers/eventmanager/eventmanager.go for more info
	Event map[int]*eventmanager.EventManager

	// The default SDL implementation of the music API. No wrappers provided at the moment
	Music *mix.Music
	Sound *mix.Chunk

	// A variable keeping track of the current screen that is being rendered. The value of this variable
	// should be one provided by src/screens/screens.go
	CurrentScreen int

	// Indicates whether the application is running
	Running bool
}

// NewEngine returns new engine.
func New(title string, width, height int32) (e *Engine) {
	winTitle = title
	winWidth = width
	winHeight = height

	e = &Engine{}
	e.Running = true
	return
}

// Init initializes SDL as well as the custom wrappers for multiple SDL packages
func (e *Engine) Init() (err error) {
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return
	}

	e.Font, err = fontmanager.New()
	if err != nil {
		return
	}

	err = mix.Init(mix.INIT_MP3)
	if err != nil {
		return
	}

	err = mix.OpenAudio(mix.DEFAULT_FREQUENCY, mix.DEFAULT_FORMAT, mix.DEFAULT_CHANNELS, 3072)
	if err != nil {
		return
	}

	e.Window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return
	}

	e.Renderer, err = sdl.CreateRenderer(e.Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return
	}

	// An event manager for every screen
	e.Event = make(map[int]*eventmanager.EventManager)
	for _, screen := range screens.Screens {
		e.Event[screen] = eventmanager.New(screen)
	}

	e.Image, err = imgmanager.New(e.Renderer)
	if err != nil {
		return
	}

	e.CurrentScreen = screens.MainScreen

	return nil
}

// Load loads resources.
func (e *Engine) Load() {
	assetDir := ""
	if runtime.GOOS != "android" {
		assetDir = filepath.Join("assets")
	}

	var err error

	err = e.Font.Load()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "load font error: %s\n", err)
	}

	err = e.Image.Load()
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "load image error: %s\n", err)
	}

	e.Music, err = mix.LoadMUS(filepath.Join(assetDir, "music", "frantic-gameplay.mp3"))
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "LoadMUS: %s\n", err)
	}

	e.Sound, err = mix.LoadWAV(filepath.Join(assetDir, "sounds", "click.wav"))
	if err != nil {
		sdl.LogError(sdl.LOG_CATEGORY_APPLICATION, "LoadWAV: %s\n", err)
	}
}

// Destroy destroys SDL and releases the memory.
func (e *Engine) Destroy() {
	e.Renderer.Destroy()
	e.Window.Destroy()
	mix.CloseAudio()

	img.Quit()
	mix.Quit()
	ttf.Quit()
	sdl.Quit()
}

// Unload unloads resources.
func (e *Engine) Unload() {

	//e.Sprite.Destroy()
	e.Font.Close()
	e.Music.Free()
	e.Sound.Free()
}

// Quit exits main loop.
func (e *Engine) Quit() {
	e.Running = false
}
