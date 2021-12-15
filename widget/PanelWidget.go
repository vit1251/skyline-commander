package widget

type PanelWidget struct {
	IWidget
}

func NewPanelWidget() *PanelWidget {
	pw := &PanelWidget{}
	return pw
}

func (self *PanelWidget) GetPanelItems() int {
	return 0
}
