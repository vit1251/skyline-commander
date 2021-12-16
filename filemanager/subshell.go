package filemanager

import (
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/tty/event"
	"github.com/vit1251/skyline-commander/widget"
	"log"
)

type CommandWidget struct {
	labelWidget *widget.LabelWidget
	inputWidget *widget.InputWidget
}

func NewCommandWidget() *CommandWidget {

	mainTerm := ctx.GetTerm()
	maxY, maxX := mainTerm.MaxYX()

	cw := &CommandWidget{}

	//
	execPrompt := widget.NewLabelWidget()
	execPrompt.SetYX(maxY-2, 0)
	execPrompt.SetTitle("$ ")
	cw.labelWidget = execPrompt

	//
	execInput := widget.NewInputWidget()
	execInput.SetValue("touch debug.log")
	execInput.SetCallback(func(value string) {
		log.Printf("Execute shell operation: %q", value)
	})
	execInput.SetYX(maxY-2, 3)
	execInput.SetWidth(maxX - 3)

	cw.inputWidget = execInput

	return cw
}

func (self *CommandWidget) Draw() {
	self.labelWidget.Draw()
	self.inputWidget.Draw()
}

func (self *CommandWidget) ProcessEvent(evt *event.Event) {

	if evt.EvType == event.EventTypeResize {
		mainTerm := ctx.GetTerm()
		maxY, _ := mainTerm.MaxYX()
		self.labelWidget.Y = maxY - 2
		self.inputWidget.Y = maxY - 2
	}

}
