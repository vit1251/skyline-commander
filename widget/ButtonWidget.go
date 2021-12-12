package widget

type ButtonWidget struct {
	//    color: &'a dyn termion::color::Color,
	//    x: u16,
	//    y: u16,
	//    title: String,
	//    callback: Option<Box<dyn Fn()>>,
	//    focused: bool,
}

func NewButtonWidget() *ButtonWidget {
	bw := &ButtonWidget{
		//            color: &Cyan,
		//            x: 0,
		//            y: 0,
		//            callback: None,
		//            title: String::new(),
		//            focused: false,
	}
	return bw
}

func ProcessInput() {

	//        match ch {
	//            '\n' => {
	//                self.invoke();
	//            },
	//            _ => {
	//            }
	//        };

}

func (self *ButtonWidget) Render() {

	//        // Step 1. Select color
	//        if self.focused {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
	///        } else {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(White));
	//        }
	//        // Step 2. Set position
	//        print!("{}", termion::cursor::Goto(self.x, self.y));
	//        // Step 3. Draw button
	//        let out = format!("[ {} ]", self.title);
	//        print!("{}", out);

}

//    pub fn set_pos(&mut self, x: u16, y: u16) {
//        self.x = x;
//        self.y = y;
//    }

//    pub fn set_title(&mut self, title: &str) {
//        self.title = String::from(title);
//    }

//    pub fn set_color(&mut self, color: &'a dyn termion::color::Color) {
//        self.color = color;
//    }
