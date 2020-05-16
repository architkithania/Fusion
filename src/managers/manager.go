package managers

// A simple abstraction of an app level manager. This interface currently is not strongly forced onto the managers
// as the main game engine simply uses the underlying types directly but rather this interface is added to provide
// somewhat of a guideline as to what is expected from a manager. Please refer to the individual managers in the
// src/managers folder for more details
type Manager interface {
	Load() error
	Close()
}