package widget

type CheckWidget struct {
	Widget
}

func NewCheckWidget() *CheckWidget {
	cw := &CheckWidget{
		Widget{
			X: 0,
			Y: 0,
		},
	}
	return cw
}
