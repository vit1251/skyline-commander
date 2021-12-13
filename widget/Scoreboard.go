package widget

import (
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
)

type Scoreboard struct {
	IScoreboard
	menu      *MenuWidget
	buttonBer *ButtonBarWidget
	groups    []*WidgetGroup
}

func (self *Scoreboard) render(pTerm *tty.PTerm, area *Rect, skin *skin.Skin) {

	/* Step 0. Render widget groups */
	for _, group := range self.groups {
		group.Render(pTerm, area)
	}

	/* Step 1. Render menu */
	//	if self.menu != nil {
	//		self.menu.Render()
	//	}

	/* Step 2. Render button bar */
	if self.buttonBer != nil {
		self.buttonBer.Render(pTerm, area, skin)
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

func (self *Scoreboard) ProcessEvent(evt *event.Event) {

	/* Process menu on active */
	//if self.menu != nil {
	//	self.menu.ProcessEvent(evt)
	//}

	/* Process group widget */
	for _, group := range self.groups {
		group.ProcessEvent(evt)
	}

}
