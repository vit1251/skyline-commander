package widget

type ScoreBoardBuilder struct {
}

func NewScoreBoardBuilder() *ScoreBoardBuilder {
	return new(ScoreBoardBuilder)
}

func (self *ScoreBoardBuilder) Build() *Scoreboard {
	sb := &Scoreboard{}
	return sb
}

func (self *ScoreBoardBuilder) WithWidgetGroup(group *WidgetGroup) *ScoreBoardBuilder {
	return self
}

func (self *ScoreBoardBuilder) WithButtonBar(buttonBar *ButtonBarWidget) *ScoreBoardBuilder {
	return self
}

func (self *ScoreBoardBuilder) WithMainMenu(menu *MenuWidget) *ScoreBoardBuilder {
	return self
}
