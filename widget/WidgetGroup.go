package widget

type WidgetGroup struct {
	widgets		[]IWidget
	active		IWidget
}

func NewWidgetGroup() (*WidgetGroup) {
	wg := WidgetGroup{}
	return wg
}

func (self *WidgetGroup) RegisterWidget(widget IWidget) {
	self.widgets = append(self.widgets, widget)
	if self.active == nil {
		self.active = widget
	}
}

func (self *WidgetGroup) ProcessEvent(evt *Event) {

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

func (self *WidgetGroup) Render(area *Rect) {
//        for w in &mut self.widgets {
//            w.render(area);
//        }
}
