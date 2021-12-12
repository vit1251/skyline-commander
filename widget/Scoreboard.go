package widget

import "github.com/gbin/goncurses"

type Scoreboard struct {
	menu      *MenuWidget
	buttonBer *ButtonBarWidget
	groups    []*WidgetGroup
}

func (self *Scoreboard) render(stdscr *goncurses.Window, area *Rect) {

	/* Step 0. Render widget groups */
	for _, group := range self.groups {
		group.Render(stdscr, area)
	}

	/* Step 1. Render menu */
	//	if self.menu != nil {
	//		self.menu.Render()
	//	}

	/* Step 2. Render button bar */
	if self.buttonBer != nil {
		self.buttonBer.Render(stdscr, area)
	}

}

func NewScoreboard() *Scoreboard {
	return new(Scoreboard)
}

func (self *Scoreboard) SetWidgetGroup(group *WidgetGroup) {
	self.groups = append(self.groups, group)
}

func (self *Scoreboard) SetButtonBar(bar *ButtonBarWidget) {
	self.buttonBer = bar
}
