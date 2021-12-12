package event

import "github.com/gbin/goncurses"

type EventType int

const (
	EventTypeKey    = EventType(1)
	EventTypeResize = EventType(2)
)

type Event struct {
	EvType EventType
	EvKey  uint
}

func NewEvent() *Event {
	evt := &Event{}
	return evt
}

func NewEventFromKey(key int) *Event {
	var evt *Event
	if key == goncurses.KEY_RESIZE {
		evt = &Event{
			EvType: EventTypeResize,
		}
	} else {
		evt = &Event{
			EvType: EventTypeKey,
			EvKey:  uint(key),
		}
	}
	return evt
}
