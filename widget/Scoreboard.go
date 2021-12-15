package widget

type Scoreboard struct {
	IScoreboard
	menu      *MenuWidget
	buttonBar *ButtonBarWidget
	groups    []*WidgetGroup
}

func (self *Scoreboard) Draw() {

	/* Step 0. Render widget groups */
	for _, group := range self.groups {
		group.Draw()
	}

	/* Step 1. Render menu */
	if self.menu != nil {
		self.menu.Draw()
	}

	/* Step 2. Render button bar */
	if self.buttonBar != nil {
		self.buttonBar.Draw()
	}

}

func NewScoreboard() *Scoreboard {
	return new(Scoreboard)
}

func (self *Scoreboard) SetWidgetGroup(group *WidgetGroup) {
	self.groups = append(self.groups, group)
}

func (self *Scoreboard) SetButtonBar(bar *ButtonBarWidget) {
	self.buttonBar = bar
}

//func (self *Scoreboard) ProcessEvent(evt *event.Event) {
//
//	/* Process menu on active */
//	//if self.menu != nil {
//	//	self.menu.ProcessEvent(evt)
//	//}
//
//	/* Process group widget */
//	for _, group := range self.groups {
//		group.ProcessEvent(evt)
//	}
//
//}
