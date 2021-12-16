package widget

import (
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
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

//func (self *WidgetGroup) ProcessEvent(evt *event.Event) {
//
//	/* Process widget group actions */
//	if evt.EvType == event.EventTypeKey {
//		if evt.EvKey == goncurses.KEY_RETURN {
//
//		}
//	}
//
//	/* Process every widget */
//	for _, w := range self.widgets {
//		w.ProcessEvent(evt)
//	}
//
//}

func (self *WidgetGroup) Draw() {
	for _, w := range self.widgets {
		log.Printf("WidgetGroup: Draw: w = %#v", w)
		w.Draw()
	}
}

func (self *WidgetGroup) ProcessEvent(evt *event.Event) {
	// TODO - TAB is switch active module ...
}
