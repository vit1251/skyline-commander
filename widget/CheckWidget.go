package widget

type CheckWidget struct {
	BaseWidget
}

func NewCheckWidget() *CheckWidget {
	cw := &CheckWidget{
		BaseWidget{
			x: 0,
			y: 0,
		},
	}
	return cw
}
