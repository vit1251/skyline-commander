package filemanager

import (
	ncursesw "github.com/vit1251/go-ncursesw"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/tty/event"
	"github.com/vit1251/skyline-commander/widget"
)

type FileManagerWidget struct {
	leftPanel     *PanelWidget
	rightPanel    *PanelWidget
	buttonBar     *widget.ButtonBarWidget
	commandWidget widget.IWidget
}

func NewFileManagerWidget() *FileManagerWidget {

	fm := &FileManagerWidget{}

	/* Left panel */
	leftPanel := NewPanelWidget()
	leftPanel.active = true

	fm.leftPanel = leftPanel

	/* Right panel */
	rightPanel := NewPanelWidget()
	rightPanel.active = false

	fm.rightPanel = rightPanel

	/* Update panel size */
	fm.updatePanelSize()

	/* Execute input group */
	commandWidget := NewCommandWidget()

	fm.commandWidget = commandWidget

	/* Create hotkey bar */
	mainBar := widget.NewButtonBarWidget()
	//mainBar.SetLabel(1, "Help")
	//mainBar.SetLabel(2, "Menu")
	//mainBar.SetLabel(3, "View")
	//mainBar.SetLabel(4, "Edit")
	//mainBar.SetLabel(5, "Copy")
	//mainBar.SetLabel(6, "Move")
	//mainBar.SetLabel(7, "MkDir")
	//mainBar.SetLabel(8, "Remove")
	//mainBar.SetLabel(9, "PullDn")
	mainBar.SetLabel(10, "Quit")

	fm.buttonBar = mainBar

	return fm
}

func (self *FileManagerWidget) Draw() {
	self.leftPanel.Draw()
	self.rightPanel.Draw()
	self.commandWidget.Draw()
	self.buttonBar.Draw()
}

func (self *FileManagerWidget) ProcessEvent(evt *event.Event) {

	if evt.EvType == event.EventTypeKey {
		if evt.EvKey == ncursesw.KEY_TAB {
			self.leftPanel.active = !self.leftPanel.active
			self.rightPanel.active = !self.rightPanel.active
		} else {
			/* Delivery input */
			self.commandWidget.ProcessEvent(evt)
			/* Delivery cursor */
			if self.leftPanel.active {
				self.leftPanel.ProcessEvent(evt)
			}
			if self.rightPanel.active {
				self.rightPanel.ProcessEvent(evt)
			}
		}
	} else if evt.EvType == event.EventTypeResize {
		self.updatePanelSize()
		self.commandWidget.ProcessEvent(evt)
	}
}

func (self *FileManagerWidget) updatePanelSize() {

	mainTerm := ctx.GetTerm()
	maxY, maxX := mainTerm.MaxYX()

	var panelWidth = maxX / 2
	var panelWidthMod = maxX % 2

	self.leftPanel.X = 0
	//self.leftPanel.Y =
	self.leftPanel.Cols = panelWidth
	self.leftPanel.Lines = maxY - 2

	//
	self.rightPanel.X = panelWidth
	//self.rightPanel.Y =
	self.rightPanel.Cols = panelWidth + panelWidthMod
	self.rightPanel.Lines = maxY - 2

}
