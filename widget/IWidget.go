package widget

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/widget/event"
)

type IWidget interface {
	ProcessEvent(evt *event.Event)
	Render(stdscr *goncurses.Window, area *Rect)
}
