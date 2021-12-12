package widget

type InputWidget struct {
//    color: &'a dyn termion::color::Color,
//    x: u16,
//    y: u16,
//    focused: bool,
//    callback: Option<Box<dyn Fn(&str)>>,
//    value: String,
//    placeholder: String,
//    is_password: bool,
//    point: u16,
//    width: u16,
}

func (self *InputWidget) ProcessInput(ch rune) {
//    match ch {
//        '\n' => {
//            self.invoke();
//        },
//        '\t' => {},
//        _ => {
//            if self.value.len() < self.width as usize {
//                self.value.push(ch);
//            }
//        }
//    };
}

func (self *InputWidget) Render(area *Rect) {
//        // Step 1. Select color
//        if self.focused {
//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
//        } else {
//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
//        }
//       // Step 2. Set position
//        print!("{}", termion::cursor::Goto(self.x, self.y));
//        // Step 3. Draw input
//        let mut out = format!("{}", self.value);
//        while out.len() < self.width as usize {
//            out.push(' ');
//        }
//        print!("{}", out);
        // Step 4. Show cursor (on focused case)
//        if self.focused {
            // Step 1. Calculate position bases on point
//            print!("{}", termion::cursor::Goto(self.x + self.point, self.y));
            // Step 2. Show cursor
//            print!("{}", termion::cursor::Show);
//        } else {
//            print!("{}", termion::cursor::Hide);
//        }
}


func NewInputWidget() *InputWidget {
	iw := &InputWidget{
//            color: &Cyan,
//            x: 0,
//            y: 0,
//            focused: true,
//            callback: None,
//            is_password: false,
//            point: 0,
//            width: 20,
//            value: String::from(""),
//            placeholder: String::from(""),
	}
	return iw
}

//    pub fn set_placeholder(&mut self, placeholder: &str) {
//        self.placeholder = String::from(placeholder);
//    }

//    pub fn set_color(&mut self, color: &'a dyn termion::color::Color) {
//        self.color = color;
//    }

//    pub fn set_pos(&mut self, x: u16, y: u16) {
//        self.x = x;
//        self.y = y;
//    }

//    pub fn set_width(&mut self, width: u16) {
//        self.width = width;
//    }

//    pub fn set_point(&mut self, point: u16) {
//        self.point = point;
//    }

//    pub fn clean(&mut self) {
//        self.value = String::from("");
//    }

//    pub fn is_empty(&self) -> bool {
//        self.value.len() == 0
//    }

//    fn invoke(&mut self) {
//        if let Some(callback) = &self.callback {
//            (callback)(&self.value);
//            trace!("Callback complete.");
//        } else {
//            trace!("No callback.");
//        }
//    }

//    pub fn set_callback(&mut self, callback: Box<dyn Fn(&str)>) {
//        self.callback = Some(callback);
//    }
