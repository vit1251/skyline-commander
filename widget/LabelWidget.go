package widget

type LabelWidget struct {
	Widget
}

func NewLabelWidget() *LabelWidget {
	lw := &LabelWidget{
		Widget{
			X: 0,
			Y: 0,
		},
	}
	return lw
}
