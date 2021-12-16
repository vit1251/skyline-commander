package widget

import (
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
)

type Widget struct {
	IWidget
	X        int /* Widget X position */
	Y        int /* Widget Y position */
	Lines    int
	Cols     int
	callback func(msg WidgetMsg)
}

func (self *Widget) SetYX(y int, x int) {
	self.Y = y
	self.X = x
}

func (self *Widget) defaultCallback(msg WidgetMsg) {
	log.Printf("Widget: defaultCallback: msg = %+v", msg)
}

func (self *Widget) Erase() {
	pTerm := ctx.GetTerm()
	pTerm.FillRegion(self.X, self.Y, self.Lines, self.Cols, ' ')
}

func (self *Widget) GotoYX(y int, x int) {
	var newY int = self.Y + y
	var newX int = self.X + x
	log.Printf("Widget: GotoXY: newX = %d newY = %d", newX, newY)
	pTerm := ctx.GetTerm()
	pTerm.Move(newY, newX)
}

func (self *Widget) Draw() {
	panic("not yet implemented")
}

func (self *Widget) ProcessEvent(evt *event.Event) {
	self.callback(MsgKey)
}
