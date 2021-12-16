package widget

import (
	"github.com/vit1251/skyline-commander/ctx"
	"log"
)

type LabelWidget struct {
	Widget
	title string
}

func NewLabelWidget() *LabelWidget {
	lw := &LabelWidget{
		Widget: Widget{
			X: 0,
			Y: 0,
		},
		title: "",
	}
	return lw
}

func (self *LabelWidget) Draw() {
	log.Printf("LabelWidget: Draw: title = %s", self.title)

	mainTerm := ctx.GetTerm()

	self.Widget.GotoYX(0, 0)

	mainTerm.Print(self.title)

}

func (self *LabelWidget) SetTitle(title string) {
	self.title = title
}
