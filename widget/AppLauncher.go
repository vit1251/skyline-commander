package widget

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
)

type AppLauncher struct {
	running     bool
	scoreBoard  *Scoreboard
	scoreBoards []*Scoreboard
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

func (self *AppLauncher) SetBoard(scoreBoard *Scoreboard) {
	self.scoreBoard = scoreBoard
	self.scoreBoards = append(self.scoreBoards, scoreBoard)
	self.updateReady = true
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
		//		self.scoreBoard.ProcessEvent(evt)
	} else {
		self.running = false
	}

}

func (self *AppLauncher) makeDefaultSkin(pTerm *tty.PTerm) *skin.Skin {

	s := skin.NewSkin()

	/* Core */
	s.Register("core", "_default_", pTerm.InitColor("lightgray", "blue"))
	s.Register("core", "selected", pTerm.InitColor("black", "cyan"))
	s.Register("core", "marked", pTerm.InitColor("yellow", "blue"))
	s.Register("core", "markselect", pTerm.InitColor("yellow", "cyan"))
	s.Register("core", "gauge", pTerm.InitColor("white", "black"))
	s.Register("core", "input", pTerm.InitColor("black", "cyan"))
	s.Register("core", "reverse", pTerm.InitColor("black", "lightgray"))

	/* Button bar */
	s.Register("buttonbar", "hotkey", pTerm.InitColor("white", "black"))
	s.Register("buttonbar", "button", pTerm.InitColor("black", "cyan"))

	return s
}

func (self *AppLauncher) Run() {

	/* Initialize Terminal */
	pTerm := tty.NewPTerm()
	err1 := pTerm.Init()
	if err1 != nil {
		panic(err1)
	}
	defer pTerm.End()
	ctx.SetTerm(pTerm)

	/* Initialize Skin */
	mainSkin := self.makeDefaultSkin(pTerm)
	mainSkin.Dump()
	ctx.SetSkin(mainSkin)

	self.running = true
	for self.running {

		/* Render scoreboard */
		if self.updateReady {
			pTerm.Touch()
			self.scoreBoard.Draw()
			pTerm.Refresh()
			log.Printf("Update: err = %v", goncurses.Update())
			self.updateReady = false
		}

		/* Process event */
		select {
		case evt := <-pTerm.C:
			log.Printf("pTerm: evt = %+v", evt)
			self.ProcessEvent(&evt)
			self.updateReady = true
		}

	}

}
