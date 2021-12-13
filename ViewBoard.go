package main

import (
	"github.com/vit1251/skyline-commander/tty/event"
	"github.com/vit1251/skyline-commander/widget"
)

type ViewBoard struct {
	widget.IWidget
	widget.Scoreboard
}

func NewViewBoard() *ViewBoard {
	return new(ViewBoard)
}

func (self *ViewBoard) ProcessEvent(evt *event.Event) {
	// Up - line up
	// Down - line down
	// Left - wrap left
	// Right - wrap right
	// F10 - quit on base scoreboard
}