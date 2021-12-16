package widget

type ScoreBoardBuilder struct {
	mainMenu  *MenuWidget
	buttonBar *ButtonBarWidget
	widget    IWidget
}

func NewScoreBoardBuilder() *ScoreBoardBuilder {
	return new(ScoreBoardBuilder)
}

func (self *ScoreBoardBuilder) Build() *Scoreboard {
	sb := NewScoreboard()
	sb.SetButtonBar(self.buttonBar)
	sb.SetWidget(self.widget)
	return sb
}

func (self *ScoreBoardBuilder) WithButtonBar(buttonBar *ButtonBarWidget) *ScoreBoardBuilder {
	self.buttonBar = buttonBar
	return self
}

func (self *ScoreBoardBuilder) WithMainMenu(menu *MenuWidget) *ScoreBoardBuilder {
	return self
}

func (self *ScoreBoardBuilder) WithWidget(widget IWidget) *ScoreBoardBuilder {
	self.widget = widget
	return self
}
