package widget

import (
	"github.com/vit1251/skyline-commander/ctx"
	"log"
)

type Widget struct {
	IWidget
	x        int /* Widget X position */
	y        int /* Widget Y position */
	lines    int
	cols     int
	callback func(msg WidgetMsg)
}

func (self *Widget) SetPos(y int, x int) {
	self.y = y
	self.x = x
}

func (self *Widget) defaultCallback(msg WidgetMsg) {
	log.Printf("Widget: defaultCallback: msg = %+v", msg)
}

func (self *Widget) Erase() {
	pTerm := ctx.GetTerm()
	pTerm.FillRegion(self.x, self.y, self.lines, self.cols, ' ')
}

func (self *Widget) GotoYX(y int, x int) {
	// TODO - ...
}

func (self *Widget) Draw() {
	if self.callback != nil {
		self.callback(MsgDraw)
	} else {
		log.Printf("Widget: no Draw callback is defined.")
	}
}
