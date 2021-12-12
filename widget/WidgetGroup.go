package widget

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/widget/event"
)

type WidgetGroup struct {
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

	//        match evt {
	//            Event::Key(
	//                Key::Char(ch)
	//            ) => {
	//                trace!("debug: key: char = {:?}", ch);
	//                for w in &mut self.widgets {
	//                    w.input(ch);
	//                }
	//            },
	//            _ => {},
	//        }

}

func (self *WidgetGroup) Render(stdscr *goncurses.Window, area *Rect) {
	//        for w in &mut self.widgets {
	//            w.render(area);
	//        }
}
