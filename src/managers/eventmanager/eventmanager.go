// This is the event manager for the application. Since SDL's event managers requires constant polling,
// abstracting this process so that events can be added and triggered from anywhere seemed like a
// worthwhile feature.
//
// The event manager currently only supports ClickEvents but can be easily extended to support other types
// of events in the future. The events are added in a stack like manner and our processed in a LIFO order.
// This is done to give precedence to newly added events which feels more intuitive during game development.
package eventmanager

import (
	"Fusion/src/managers/eventmanager/events"

	"github.com/veandco/go-sdl2/sdl"
)

type EventManager struct {
	// Stores the screen id as provided by src/screens/screens.go
	screen int

	// Stores references to ClickEvent objects that have registered as an event provider
	// for this particular screen. Thought has been given to make this variable private
	// to truly encapsulate event management but in the end decided against it as sometimes
	// the need for upcoming events are a useful insight into app level state. Perhaps in the
	// future a getter can be provided to provide more control over the process.
	RegisteredClicks []events.ClickEvent
}

// Provided constructor
func New(screen int) *EventManager {
	return &EventManager{
		screen:           screen,
		RegisteredClicks: make([]events.ClickEvent, 0, 5),
	}
}

// Registers a ClickEvent object as an event provider for this particular screen.
func (em *EventManager) RegisterEvent(event events.ClickEvent) {
	em.RegisteredClicks = append(em.RegisteredClicks, event)
}

// Process and fires an event out of all the registered events of this screen. This method
// simply iterates (in reverse order) over all the objects and fires the *most recently*
// added ClickEvent. It does so by comparing the mouse position at the time of click with
// the the object positions of all the objects in the RegisteredClick slice. Since the scan
// is linear, it is best to divide events into multiple screens and scan accordingly
func (em *EventManager) ProcessClickEvents(mouseEv *sdl.MouseButtonEvent) error {
	for i := len(em.RegisteredClicks) - 1; i >= 0; i-- {
		e := em.RegisteredClicks[i]
		if mouseEv.X >= e.GetX() && mouseEv.X <= (e.GetX()+e.GetWidth()) &&
			mouseEv.Y >= e.GetY() && mouseEv.Y <= (e.GetY()+e.GetHeight()) {
			return e.RunCallback(e)
		}
	}
	return nil
}
