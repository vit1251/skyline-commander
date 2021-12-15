package widget

type LabelWidget struct {
	BaseWidget
}

func NewLabelWidget() *LabelWidget {
	lw := &LabelWidget{
		BaseWidget{
			x: 0,
			y: 0,
		},
	}
	return lw
}
