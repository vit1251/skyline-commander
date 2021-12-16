package event

import "github.com/vit1251/goncurses"

type EventType int

const (
	EventTypeKey    = EventType(1)
	EventTypeResize = EventType(2)
)

type Event struct {
	EvType EventType
	EvKey  goncurses.Key
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
			EvKey:  goncurses.Key(key),
		}
	}
	return evt
}
