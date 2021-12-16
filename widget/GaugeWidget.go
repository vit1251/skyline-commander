package widget

import "github.com/vit1251/skyline-commander/tty/event"

type GaugeWidget struct {
	Widget
	max     uint64
	current uint64
	width   uint16
}

func (self *GaugeWidget) ProcessEvent(evt *event.Event) {
}

func (self *GaugeWidget) Render(area *Rect) {
	//        let width: u16 = self.width - 2;
	//        // Step 1. Select color
	//        // Step 2. Set position
	//        print!("{}", termion::cursor::Goto(self.x, self.y));
	//        // Step 3. Draw bar
	//        print!("{}{}", termion::color::Fg(Black), termion::color::Bg(White));
	//        print!("[");
	//        // Active
	//        print!("{}", termion::color::Bg(Black));
	//        let active = (width as u64 * self.current) / self.max;
	//        let mut out = String::new();
	//        for _ in 0..active {
	//            out.push(' ');
	//        }
	//        print!("{}", out);
	//        // Backlog
	//        let backlog = width as u64 - active;
	//        print!("{}", termion::color::Bg(White));
	//        let mut out = String::new();
	//        for _ in 0..backlog {
	//            out.push(' ');
	//        }
	//        print!("{}", out);
	//        //
	//        print!("{}{}", termion::color::Fg(Black), termion::color::Bg(White));
	//        print!("]");
}

func NewGaugeWidget() *GaugeWidget {
	gw := &GaugeWidget{
		Widget: Widget{
			X: 0,
			Y: 0,
		},
		current: 0,
		max:     100,
		width:   20,
	}
	return gw
}

func (self *GaugeWidget) SetWidth(width uint16) {
	self.width = width
}

//func (self *GaugeWidget) SetColor(color string) {
//	self.color = color
//}

func (self *GaugeWidget) SetValue(max uint64, current uint64) {
	self.max = max
	self.current = current
}
