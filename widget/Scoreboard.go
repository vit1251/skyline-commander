package widget

import (
	"github.com/vit1251/skyline-commander/tty/event"
)

type Scoreboard struct {
	IScoreboard
	menu      *MenuWidget
	buttonBar *ButtonBarWidget
	widget    IWidget
}

func (self *Scoreboard) Draw() {

	/* Step 1. Render main widget */
	self.widget.Draw()

	/* Step 2. Render menu */
	if self.menu != nil {
		self.menu.Draw()
	}

	/* Step 3. Render button bar */
	if self.buttonBar != nil {
		self.buttonBar.Draw()
	}

}

func NewScoreboard() *Scoreboard {
	return new(Scoreboard)
}

func (self *Scoreboard) SetButtonBar(bar *ButtonBarWidget) {
	self.buttonBar = bar
}

func (self *Scoreboard) SetWidget(widget IWidget) {
	self.widget = widget
}

func (self *Scoreboard) ProcessEvent(evt *event.Event) {
	self.widget.ProcessEvent(evt)
}
