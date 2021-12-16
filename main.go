package main

import (
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/filemanager"
	"github.com/vit1251/skyline-commander/widget"
	"io"
	"log"
	"os"
	"runtime"
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

	mainTerm := ctx.GetTerm()
	maxY, maxX := mainTerm.MaxYX()

	mainWidgetGroup := widget.NewWidgetGroup()

	/* Create filemanager panel */
	filemanager.CreatePanel(mainWidgetGroup)

	/* Execute input group */
	execWidgetGroup := widget.NewWidgetGroup()
	execPrompt := widget.NewLabelWidget()
	execPrompt.SetYX(maxY-2, 0)
	execPrompt.SetTitle("$ ")
	execWidgetGroup.RegisterWidget(execPrompt)

	execInput := widget.NewInputWidget()
	execInput.SetValue("touch debug.log")
	execInput.SetCallback(func(value string) {
		log.Printf("Execute shell operation: %q", value)
	})
	execInput.SetYX(maxY-2, 3)
	execInput.SetWidth(maxX - 3)
	execWidgetGroup.RegisterWidget(execInput)
	mainWidgetGroup.RegisterWidget(execWidgetGroup)

	/* Create hotkey bar */
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
		WithWidget(mainWidgetGroup).
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

	/* Use one OS invoke source */
	runtime.LockOSThread()

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

	/* Create application with single scoreboard */
	app := widget.AppLauncherWithScoreboard(func() *widget.Scoreboard {
		return createPanelBoard()
	})

	/* Run application */
	app.Run()

}
