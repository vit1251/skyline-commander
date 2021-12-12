package widget

import (
	"github.com/famz/SetLocale"
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/widget/event"
	"log"
)

type AppLauncher struct {
	running     bool
	scoreBoard  *Scoreboard
	stdscr      *goncurses.Window
	updateReady bool
}

func AppLauncherWithScoreboard(sb *Scoreboard) *AppLauncher {
	al := &AppLauncher{
		running:     false,
		scoreBoard:  sb,
		updateReady: true,
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

func (self *AppLauncher) ProcessEvent(evt *event.Event) {
	if evt.EvType == event.EventTypeKey {
		if evt.EvKey == goncurses.KEY_F10 {
			self.running = false
		}
	}
}

func (self *AppLauncher) reset() {
}

func (self *AppLauncher) Run() {

	/* Disable GetText encoding */
	SetLocale.SetLocale(SetLocale.LC_ALL, "")

	log.Printf("Initialize ncurses: version = %s", goncurses.CursesVersion())

	stdscr, err1 := goncurses.Init()
	if err1 != nil {
		log.Fatal("fail on Init", err1)
	}
	self.stdscr = stdscr
	defer goncurses.End()

	goncurses.CBreak(true)
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

	self.running = true
	for self.running {

		/* Render scoreboard */
		if self.updateReady {
			self.reset()
			self.render()
			stdscr.Refresh()
			self.updateReady = false
		}

		/* Process input */
		var key goncurses.Key = stdscr.GetChar()
		if key != 0 {
			evt := event.NewEventFromKey(int(key))
			log.Printf("Event: evt = %+v", evt)
			self.ProcessEvent(evt)
			self.updateReady = true
		}

	}

}
