package widget

type WidgetMsg int

const (
	MsgInit    = WidgetMsg(1)
	MsgDraw    = WidgetMsg(2)
	MsgFocus   = WidgetMsg(3)
	MsgUnfocus = WidgetMsg(4)
	MsgKey     = WidgetMsg(5)
	MsgAction  = WidgetMsg(6)
	MsgDestroy = WidgetMsg(7)
)
