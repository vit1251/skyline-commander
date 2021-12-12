package widget

import (
	"github.com/gbin/goncurses"
	"log"
	"time"
)

type AppLauncher struct {
//    state: T,
//    view: WidgetGroup,
//    running: bool,
}

func AppLauncherWithScoreboard(sb *Scoreboard) *AppLauncher {
	al := &AppLauncher {
//            view: view,
//            state: state,
//            running: false,
	}
	return al
}

//    fn size(&self) -> std::io::Result<Rect> {
//        let terminal = termion::terminal_size()?;
//        Ok(Rect::new(0, 0, terminal.0, terminal.1))
//    }

//    fn render(&mut self) {
//        let area = self.size().unwrap();
//        self.view.render(&area);
//    }

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
		log.Fatal("init:", err1)
	}
	defer goncurses.End()

	err2 := goncurses.StartColor()
	if err2 != nil {
		log.Fatal("StartColor", err2)
	}

	stdscr.Print("Press enter to continue...")
	stdscr.Refresh()

	time.Sleep(1 * time.Minute)
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
