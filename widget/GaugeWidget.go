package widget

type GaugeWidget struct {
//    color: &'a dyn termion::color::Color,
//    x: u16,
//    y: u16,
//    max: u64,
//    current: u64,
//    width: u16,
}

func (self *GaugeWidget) ProcessInput(ch rune) {
//        match ch {
//            '+' => {
//                self.current = self.current + 1;
//            },
//            '-' => {
//                self.current = self.current - 1;
//            },
//            _ => {},
//        };
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
            color: &White,
            x: 0,
            y: 0,
            current: 0,
            max: 100,
            width: 20,
	}
	return gw
}

//    pub fn set_width(&mut self, width: u16) {
//        self.width = width;
//    }

//    pub fn set_color(&mut self, color: &'a dyn termion::color::Color) {
//        self.color = color;
//    }

//    pub fn set_pos(&mut self, x: u16, y: u16) {
//        self.x = x;
//        self.y = y;
//    }

//    pub fn set_value(&mut self, max: u64, current: u64) {
//        self.max = max;
//        self.current = current;
//    }
