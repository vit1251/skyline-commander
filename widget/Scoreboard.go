package widget

import "github.com/gbin/goncurses"

type Scoreboard struct {
	menu      *MenuWidget
	buttonBer *ButtonBarWidget
	groups    []*WidgetGroup
}

func (self *Scoreboard) render(stdscr *goncurses.Window, area *Rect) {
	for _, group := range self.groups {
		group.Render(stdscr, area)
	}
}

func NewScoreboard() *Scoreboard {
	return new(Scoreboard)
}
