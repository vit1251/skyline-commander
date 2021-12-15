package widget

type MenuEntry struct {
}

type Menu struct {
}

type MenuWidget struct {
	IWidget
}

func NewMenuWidget() *MenuWidget {
	mw := &MenuWidget{}
	return mw
}
