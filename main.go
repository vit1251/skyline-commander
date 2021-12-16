package main

import (
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

	/* Create main File Manager widget */
	mainFileManager := filemanager.NewFileManagerWidget()

	/* Create main Scoreboard */
	mainBoard := widget.NewScoreBoardBuilder().
		WithWidget(mainFileManager).
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
