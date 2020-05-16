package buttons

import (
	"Fusion/src/managers/eventmanager/events"
)

// An interface that is implemented by every Button. Currently this interface is not very useful
// and a button could simply implement the events.ClickEvent interface but this is added mainly
// as a future safety if some aspects of a button need to be abstracted. It is therefore
// recommended to use the button interface for typing as it may be changed in the future.
type Button interface {
	events.ClickEvent
}
