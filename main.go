package main

import (
	"github.com/vit1251/skyline-commander/widget"
	"io"
	"log"
	"os"
)

type BoardType int

const (
	BoardAbort = BoardType(1)
	UserMenu   = BoardType(2)
	BoardPanel = BoardType(3)
	Copy       = BoardType(4)
	Move       = BoardType(5)
	View       = BoardType(6)
	Edit       = BoardType(7)
	Remove     = BoardType(8)
)

func createPanelBoard() *widget.Scoreboard {

	mainWidgetGroup := widget.NewWidgetGroup()

	/* Left panel */
	leftPanel := widget.NewPanelWidget()
	mainWidgetGroup.RegisterWidget(leftPanel)

	/* Right panel */
	rightPanel := widget.NewPanelWidget()
	mainWidgetGroup.RegisterWidget(rightPanel)

	/* Input panel */
	mainInput := widget.NewInputWidget()
	mainInput.SetCallback(func(value string) {
		log.Printf("Execute shell operation: %q", value)
	})
	mainInput.SetPos(10, 0)
	mainWidgetGroup.RegisterWidget(mainInput)

	/**/
	mainBar := widget.NewButtonBarWidget()
	mainBar.SetLabel(1, "Help")
	mainBar.SetLabel(2, "Menu")
	mainBar.SetLabel(3, "View")
	mainBar.SetLabel(4, "Edit")
	mainBar.SetLabel(5, "Copy")
	mainBar.SetLabel(6, "Move")
	mainBar.SetLabel(7, "MkDir")
	mainBar.SetLabel(8, "Remove")
	mainBar.SetLabel(9, "PullDn")
	mainBar.SetLabel(10, "Quit")

	mainMenu := widget.NewMenuWidget()

	/* Create main Scoreboard */
	mainBoard := widget.NewScoreBoardBuilder().
		WithWidgetGroup(mainWidgetGroup).
		WithButtonBar(mainBar).
		WithMainMenu(mainMenu).
		Build()

	return mainBoard
}

func createTerminalBoard() *widget.Scoreboard {
	return nil
}

func createCpuBoard() *widget.Scoreboard {
	return nil
}

func createViewBoard() *widget.Scoreboard {
	return nil
}

func createEditBoard() *widget.Scoreboard {
	return nil
}

func createCopyBoard() *widget.Scoreboard {
	return nil
}

func createQuitBoard() *widget.Scoreboard {
	return nil
}

func main() {

	/* Setup debug output */
	stream, err1 := os.Create("debug.log")
	if err1 != nil {
		panic(err1)
	}
	defer func() {
		log.SetOutput(io.Discard)
		err2 := stream.Close()
		if err2 != nil {
			panic(err2)
		}
	}()
	log.SetOutput(stream)

	mainBoard := createPanelBoard()

	/* Create application with single scoreboard */
	app := widget.AppLauncherWithScoreboard(mainBoard)

	/* Run application */
	app.Run()

}
