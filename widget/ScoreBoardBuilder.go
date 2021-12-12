package widget

type ScoreBoardBuilder struct {
	group     *WidgetGroup
	buttonBar *ButtonBarWidget
}

func NewScoreBoardBuilder() *ScoreBoardBuilder {
	return new(ScoreBoardBuilder)
}

func (self *ScoreBoardBuilder) Build() *Scoreboard {
	sb := NewScoreboard()
	sb.SetWidgetGroup(self.group)
	sb.SetButtonBar(self.buttonBar)
	return sb
}

func (self *ScoreBoardBuilder) WithWidgetGroup(group *WidgetGroup) *ScoreBoardBuilder {
	self.group = group
	return self
}

func (self *ScoreBoardBuilder) WithButtonBar(buttonBar *ButtonBarWidget) *ScoreBoardBuilder {
	self.buttonBar = buttonBar
	return self
}

func (self *ScoreBoardBuilder) WithMainMenu(menu *MenuWidget) *ScoreBoardBuilder {
	return self
}
