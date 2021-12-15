package widget

type BaseWidget struct {
	IWidget
	x int /* Widget X position */
	y int /* Widget Y position */
}

func (self *BaseWidget) SetPos(y int, x int) {
	self.y = y
	self.x = x
}
