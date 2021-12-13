package widget

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
)

type AppLauncher struct {
	running     bool
	scoreBoard  *Scoreboard
	scoreBoards []*Scoreboard
	pTerm       *tty.PTerm
	updateReady bool
	skin        *skin.Skin
}

func AppLauncherWithScoreboard(sb *Scoreboard) *AppLauncher {
	al := &AppLauncher{
		running:     false,
		scoreBoard:  sb,
		updateReady: true,
	}
	return al
}

func (self *AppLauncher) SetBoard(scoreBoard *Scoreboard) {
	self.scoreBoard = scoreBoard
	self.scoreBoards = append(self.scoreBoards, scoreBoard)
	self.updateReady = true
}

func (self *AppLauncher) size() *Rect {
	y, x := self.pTerm.MaxYX()
	return NewRect(0, 0, uint(x), uint(y))
}

func (self *AppLauncher) render() {

	area := self.size()
	self.scoreBoard.render(self.pTerm, area, self.skin)
}

func (self *AppLauncher) ProcessEvent(evt *event.Event) {

	if evt.EvType == event.EventTypeKey {
		log.Printf("key = %+v", goncurses.KeyString(goncurses.Key(evt.EvKey)))
		if evt.EvKey == goncurses.KEY_F10 {
			var scoreBoardCount int = len(self.scoreBoards)

			/* Notify board about exit */
			//var activeScoreBoard = self.scoreBoards[scoreBoardCount - 1] // TODO - process exit event ...

			/* Remove in board stack registry */
			if scoreBoardCount > 0 {
				self.scoreBoards = self.scoreBoards[:scoreBoardCount-1]
			}

			/* Select previous scoreboard */
			var prevScoreBoard *Scoreboard = nil
			if scoreBoardCount > 0 {
				prevScoreBoard = self.scoreBoards[scoreBoardCount-1]
			}
			self.scoreBoard = prevScoreBoard
		}
	}

	if self.scoreBoard != nil {
		self.scoreBoard.ProcessEvent(evt)
	} else {
		self.running = false
	}

}

func (self *AppLauncher) reset() {
	self.pTerm.Erase()
}

func (self *AppLauncher) makeDefaultSkin() *skin.Skin {

	s := skin.NewSkin()

	/* Core */
	s.Register("core", "_default_", self.pTerm.InitColor("lightgray", "blue"))
	s.Register("core", "selected", self.pTerm.InitColor("black", "cyan"))
	s.Register("core", "marked", self.pTerm.InitColor("yellow", "blue"))
	s.Register("core", "markselect", self.pTerm.InitColor("yellow", "cyan"))
	s.Register("core", "gauge", self.pTerm.InitColor("white", "black"))
	s.Register("core", "input", self.pTerm.InitColor("black", "cyan"))
	s.Register("core", "reverse", self.pTerm.InitColor("black", "lightgray"))

	/* Button bar */
	s.Register("buttonbar", "hotkey", self.pTerm.InitColor("white", "black"))
	s.Register("buttonbar", "button", self.pTerm.InitColor("black", "cyan"))

	return s
}

func (self *AppLauncher) Run() {

	self.pTerm = tty.NewPTerm()
	err1 := self.pTerm.Init()
	if err1 != nil {
		panic(err1)
	}
	defer self.pTerm.End()

	/* Prepare main Skin */
	self.skin = self.makeDefaultSkin()
	self.skin.Dump()

	self.running = true
	for self.running {

		/* Render scoreboard */
		if self.updateReady {
			self.pTerm.Touch()
			self.reset()
			self.render()
			self.pTerm.Refresh()
			log.Printf("Update: err = %v", goncurses.Update())
			self.updateReady = false
		}

		/* Process event */
		select {
		case evt := <-self.pTerm.C:
			log.Printf("pTerm: evt = %+v", evt)
			self.ProcessEvent(&evt)
			self.updateReady = true
		}

	}

}
