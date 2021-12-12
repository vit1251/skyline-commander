package main

import (
	"github.com/vit1251/skyline-commander/widget"
)

func main() {

	mainWidgetGroup := widget.NewWidgetGroup()

	mainBar := widget.NewButtonBarWidget()

	mainMenu := widget.NewMenuWidget()

	/* Create main Scoreboard */
	mainBoard := widget.NewScoreBoardBuilder().
		WithWidgetGroup(mainWidgetGroup).
		WithButtonBar(mainBar).
		WithMainMenu(mainMenu).
		Build()

	/* Create application with single scoreboard */
	app := widget.AppLauncherWithScoreboard(mainBoard)

	/* Run application */
	app.Run()

}
