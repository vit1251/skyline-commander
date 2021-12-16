package widget

import (
	"fmt"
	"github.com/vit1251/goncurses"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/tty/event"
	"unicode"
	"unicode/utf8"
)

type InputWidget struct {
	Widget
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
		var charRune rune = rune(evt.EvKey)
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
		} else if unicode.IsNumber(charRune) || unicode.IsLetter(charRune) {
			self.value = fmt.Sprintf("%s%c", self.value, rune(evt.EvKey))
		}

	}
}

func (self *InputWidget) Draw() {

	pTerm := ctx.GetTerm()

	/* Step 1. Select color */
	//        if self.focused {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
	//        } else {
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(self.color));
	//        }

	/* Step 2. Set position */
	self.Widget.GotoYX(0, 0)

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
		Widget: Widget{
			X: 0,
			Y: 0,
		},
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

func (self *InputWidget) SetPlaceholder(placeholder string) {
	self.placeholder = placeholder
}

//    pub fn set_color(&mut self, color: &'a dyn termion::color::Color) {
//        self.color = color;
//    }

//    pub fn set_point(&mut self, point: u16) {
//        self.point = point;
//    }

func (self *InputWidget) Clean() {
	self.value = ""
}

func (self *InputWidget) IsEmpty() bool {
	return self.value == ""
}

func (self *InputWidget) SetCallback(callback func(string)) *InputWidget {
	self.callback = callback
	return self
}

func (self *InputWidget) SetWidth(width int) *InputWidget {
	self.width = width
	return self
}

func (self *InputWidget) SetValue(value string) *InputWidget {
	self.value = value
	return self
}
