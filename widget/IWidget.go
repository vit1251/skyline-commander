package widget

import (
	"github.com/vit1251/skyline-commander/tty/event"
)

type IWidget interface {
	Draw()
	ProcessEvent(evt *event.Event)
}
