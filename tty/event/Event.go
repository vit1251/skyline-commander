package event

import (
	ncursesw "github.com/vit1251/go-ncursesw"
)

type EventType int

const (
	EventTypeKey    = EventType(1)
	EventTypeResize = EventType(2)
)

type Event struct {
	EvType EventType
	EvKey  ncursesw.Key
}

func NewEvent() *Event {
	evt := &Event{}
	return evt
}

func NewEventFromKey(key int) *Event {
	var evt *Event
	if key == ncursesw.KEY_RESIZE {
		evt = &Event{
			EvType: EventTypeResize,
		}
	} else {
		evt = &Event{
			EvType: EventTypeKey,
			EvKey:  ncursesw.Key(key),
		}
	}
	return evt
}
