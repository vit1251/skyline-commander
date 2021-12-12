package widget

type Rect struct {
	x      uint
	y      uint
	width  uint
	height uint
}

func NewRect(x uint, y uint, width uint, height uint) *Rect {
	return &Rect{
		x, y, width, height,
	}
}
