package widget

type IWidget interface {
	Input(ch rune)
	Render(area *Rect)
}
