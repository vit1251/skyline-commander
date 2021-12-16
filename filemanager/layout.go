package filemanager

import (
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/widget"
)

//var leftPanel *PanelWidget
//var rightPanel *PanelWidget

func CreatePanel(mainWidgetGroup *widget.WidgetGroup) {

	mainTerm := ctx.GetTerm()
	maxY, maxX := mainTerm.MaxYX()

	var panelWidth = maxX / 2
	var panelWidthMod = maxX % 2

	/* Left panel */
	leftPanel := NewPanelWidget()
	leftPanel.SetYX(0, 0)
	leftPanel.Lines = maxY - 2
	leftPanel.Cols = panelWidth
	mainWidgetGroup.RegisterWidget(leftPanel)

	/* Right panel */
	rightPanel := NewPanelWidget()
	rightPanel.SetYX(0, maxX/2)
	rightPanel.Lines = maxY - 2
	rightPanel.Cols = panelWidth + panelWidthMod
	mainWidgetGroup.RegisterWidget(rightPanel)

}
