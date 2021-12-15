package widget

type CheckWidget struct {
	Widget
}

func NewCheckWidget() *CheckWidget {
	cw := &CheckWidget{
		Widget{
			x: 0,
			y: 0,
		},
	}
	return cw
}
