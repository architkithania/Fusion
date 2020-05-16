// Use this package as an API to the different screens that will be used throughout the application. This
// package also provides an array with all the screens so that iteration may be done over the screens.
package screens

const (
	MainScreen = iota
	GameScreen
	SettingsScreen
)

var Screens = [...]int{
	MainScreen,
	GameScreen,
	SettingsScreen,
}
