package events

// The interface that must be implemented by any component that wishes to fire a callback upon being clicked.
// For more information refer to the documentation on the event manager in src/managers/eventmanager/eventmanager.go
type ClickEvent interface {
	GetX() int32
	GetY() int32
	GetWidth() int32
	GetHeight() int32
	RunCallback(...interface{}) error
}
