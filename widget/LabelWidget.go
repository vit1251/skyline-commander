package widget

type LabelWidget struct {
	Widget
}

func NewLabelWidget() *LabelWidget {
	lw := &LabelWidget{
		Widget{
			x: 0,
			y: 0,
		},
	}
	return lw
}
