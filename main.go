package main

import (
	"github.com/vit1251/skyline-commander/widget"
)

func main() {

	/* Create main Scoreboard */
	mainBoard := ScoreBoardBuilder().
		WithWidgetGroup(mainWidgetGroup).
		WithButtonBar(mainBar).
		WithMainMenu(mainMenu).
		Build()

	/* Create application with single scoreboard */
	app := widget.AppLauncherWithScoreboard(mainBoard)

	/* Run application */
	arr.Run()

}
