package widget

import (
	"github.com/gbin/goncurses"
	"log"
)

type AppLauncher struct {
	running    bool
	scoreBoard *Scoreboard
	stdscr     *goncurses.Window
}

func AppLauncherWithScoreboard(sb *Scoreboard) *AppLauncher {
	al := &AppLauncher{
		running:    false,
		scoreBoard: sb,
	}
	return al
}

func (self *AppLauncher) size() *Rect {
	y, x := self.stdscr.MaxYX()
	return NewRect(0, 0, uint(x), uint(y))
}

func (self *AppLauncher) render() {
	area := self.size()
	self.scoreBoard.render(self.stdscr, area)
}

//    fn process_event(&mut self, evt: Event) {
//        match evt {
//            Event::Key(
//                Key::F(10)
//            ) => {
//                self.running = false;
//            },
//            _ => {
//                self.view.process_event(evt);
//            }
//        };
//    }

func (self *AppLauncher) Run() {

	stdscr, err1 := goncurses.Init()
	if err1 != nil {
		log.Fatal("fail on Init", err1)
	}
	self.stdscr = stdscr
	defer goncurses.End()

	goncurses.Raw(true)
	goncurses.Echo(false)
	err2 := goncurses.Cursor(0)
	if err2 != nil {
		log.Fatal("fail on Cursor", err2)
	}
	err3 := stdscr.Keypad(true)
	if err3 != nil {
		log.Fatal("fail on Keypad", err3)
	}

	stdscr.Timeout(100)

	err4 := goncurses.StartColor()
	if err4 != nil {
		log.Fatal("fail on StartColor", err4)
	}

	stdscr.Print("Press enter to continue...")
	stdscr.Refresh()

	self.running = true
	for self.running {

		/* Render scoreboard */
		self.render()

		/* Process input */
		key := stdscr.GetChar()
		switch key {
		case goncurses.KEY_F10:
			log.Printf("Press F10 key\n")
			self.running = false
		}

	}

}

//
//    fn reset(&self) {
//        print!("{}{}{}{}",
//               termion::color::Fg(termion::color::Reset),
//               termion::color::Bg(termion::color::Reset),
//               termion::clear::All,
//               termion::cursor::Show,
//            );
//    }
//
