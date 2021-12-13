package widget

import (
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
)

type IWidget interface {
	ProcessEvent(evt *event.Event)
	Render(stdscr *tty.PTerm, area *Rect)
}
