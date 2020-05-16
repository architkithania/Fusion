// Provides an abstraction of working with fonts within the SDL2 library. Since SDL2 doesn't support
// dynamic resizing of fonts, a different version of the font's .ttf file must be opened for every size
// change. This package provides a way to efficiently manage and load different fonts by caching font
// objects internally so that the same font does not need to be repeatedly opened.
//
// All the fonts that will be used *MUST* be stored in the assets/fonts directory with only one '.' before
// the file extension. For example: awesome.font.ttf is an *INVALID* font name as it contains two periods.
//
// In the future an implementation of a caching strategy to free fonts that are not recently used can be created.
package fontmanager

import (
	"fmt"
	"github.com/veandco/go-sdl2/ttf"
	"path/filepath"
	"runtime"
	"strings"
)

// The set of fonts that should be loaded as the applications starts. Place the most commonly used fonts
// here so that the
var PRE_LOADED_FONTS = map[string]int{
	"universalfruitcake.ttf": 24,
}

type FontManager struct {
	fonts map[struct{string;int}]*ttf.Font
}

// Provided constructor
func New() (*FontManager, error) {
	err := ttf.Init()
	if err != nil {
		return nil, err
	}
	fManager := FontManager{make(map[struct{string;int}]*ttf.Font)}

	return &fManager, nil
}

// The main usage API of this package. The user can use this function to specify a font name and a
// size and the font manager would efficiently provide the font, be it by using a previously cached
// result or opens a new font instance if one wasn't previously cached
func (fManager *FontManager) GetFont(font string, size int) (*ttf.Font, bool) {
	assetDir := ""
	if runtime.GOOS != "android" {
		assetDir = filepath.Join( "assets")
	}

	if val, ok := fManager.fonts[struct{string;int}{font,size}]; ok {
		return val, true
	}

	var err error
	fontPack, err := ttf.OpenFont(filepath.Join(assetDir, "fonts", font + ".ttf"), size)

	if err != nil {
		fmt.Println(err)
		return fManager.fonts[struct{string;int}{"universalfruitcake", 24}], false
	}

	fManager.fonts[struct{string;int}{font,size}] = fontPack
	return fontPack, true
}

// Loads the library as well as pre-caches all the fonts in the PRE_LOADED_FONTS array
func (fManager *FontManager) Load() error {
	assetDir := ""
	if runtime.GOOS != "android" {
		assetDir = filepath.Join( "assets")
	}

	var err error
	var key struct{string;int}
	for font, size := range PRE_LOADED_FONTS {
		fontName:= strings.Split(font, ".")[0]
		key = struct{string;int}{fontName,size}
		fManager.fonts[key], err = ttf.OpenFont(filepath.Join(assetDir,"fonts", font), size)
		if err != nil {
			return err
		}
	}

	return nil
}

// Closes and frees all the fonts in the cache
func (fManager *FontManager) Close() {
	for _, font := range fManager.fonts {
		font.Close()
	}
}