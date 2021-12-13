package widget

import (
	"fmt"
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
	"unicode/utf8"
)

type InputWidget struct {
	IWidget
	x           int
	y           int
	focused     bool
	value       string
	placeholder string
	isPassword  bool
	point       int
	width       int
	callback    func(value string)
}

func (self *InputWidget) ProcessEvent(evt *event.Event) {
	if evt.EvType == event.EventTypeKey {
		if evt.EvKey == goncurses.KEY_RETURN {
			if self.callback != nil {
				self.callback(self.value)
			}
		} else if evt.EvKey == goncurses.KEY_BACKSPACE {
			var runes []rune = []rune(self.value)
			var runeCount int = len(runes)
			if runeCount > 0 {
				runes = runes[:runeCount-1]
				self.value = string(runes)
			}
		} else {
			self.value = fmt.Sprintf("%s%c", self.value, rune(evt.EvKey))
		}

	}
}

func (self *InputWidget) Render(pTerm *tty.PTerm, area *Rect) {

	/* Step 1. Select color */
	//        if self.focused {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
	//        } else {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
	//        }

	/* Step 2. Set position */
	pTerm.Move(self.y, self.x)

	/* Step 3. Draw input */
	var out string = fmt.Sprintf("%s", self.value)
	for utf8.RuneCountInString(out) < self.width {
		out = fmt.Sprintf("%s ", out)
	}
	//pTerm.ColorOn()
	pTerm.Print(out)
	//pTerm.ColorOff()

	/* Step 4. Show cursor (on focused case) */
	//        if self.focused {
	//            print!("{}", termion::cursor::Goto(self.x + self.point, self.y));
	//            print!("{}", termion::cursor::Show);
	//        } else {
	//            print!("{}", termion::cursor::Hide);
	//        }
}

func NewInputWidget() *InputWidget {
	iw := &InputWidget{
		x:           0,
		y:           0,
		focused:     true,
		callback:    nil,
		isPassword:  false,
		point:       0,
		width:       20,
		value:       "",
		placeholder: "",
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

func (self *InputWidget) SetCallback(callback func(string)) *InputWidget {
	self.callback = callback
	return self
}

func (self *InputWidget) SetPos(y int, x int) *InputWidget {
	self.y = y
	self.x = x
	return self
}
