package widget

type ButtonWidget struct {
	Widget
	title    string
	callback func()
	focused  bool
}

func NewButtonWidget() *ButtonWidget {
	bw := &ButtonWidget{
		Widget: Widget{
			X: 0,
			Y: 0,
		},
		callback: nil,
		title:    "",
		focused:  false,
	}
	return bw
}

func (self *ButtonWidget) Draw() {

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

func (self *ButtonWidget) SetTitle(title string) {
	self.title = title
}
