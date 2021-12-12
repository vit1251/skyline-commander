package widget

type Rect struct {
	x		uint
	y		uint
	width		uint
	height		uint
}

interface Widget {
    func input(&mut self, ch: char);
    func render(&mut self, area: &Rect);
}
