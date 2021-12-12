package widget

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/widget/event"
)

type WidgetGroup struct {
	IWidget
	widgets []IWidget
	active  IWidget
}

func NewWidgetGroup() *WidgetGroup {
	wg := &WidgetGroup{}
	return wg
}

func (self *WidgetGroup) RegisterWidget(widget IWidget) {
	self.widgets = append(self.widgets, widget)
	if self.active == nil {
		self.active = widget
	}
}

func (self *WidgetGroup) ProcessEvent(evt *event.Event) {
	if evt.EvType == event.EventTypeKey {
		if evt.EvKey == goncurses.KEY_RETURN {

		}
	}
}

func (self *WidgetGroup) Render(stdscr *goncurses.Window, area *Rect) {
	for _, w := range self.widgets {
		w.Render(stdscr, area)
	}
}
