package widget

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

//    pub fn launch(&mut self) {
//
//        let stdin = stdin();
//        let mut stdout = MouseTerminal::from(stdout().into_raw_mode().unwrap());
//
//        let mut events = stdin.events();
//
//        self.running = true;
//        while self.running {
//
//            self.reset();
//
//            self.render();
//
//            print!("{}", termion::cursor::Hide);
//
//            stdout.flush().unwrap();
//
//            let evt = events.next(); // Option<Result<...>>
//            let evt = evt.unwrap();  // Result<Event, Error>
//            let evt = evt.unwrap();  // Event
//
//            self.process_event(evt);
//
//        }
//
//        /* Application complete */
//        self.reset();
//
//    }
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
