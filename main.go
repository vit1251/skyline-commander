package main

import (
	"github.com/gbin/goncurses"
	"log"
	"time"
)

func main() {

	stdscr, err1 := goncurses.Init()
	if err1 != nil {
		log.Fatal("init:", err1)
	}
	defer goncurses.End()

	err2 := goncurses.StartColor()
	if err2 != nil {
		log.Fatal("StartColor", err2)
	}

	stdscr.Print("Press enter to continue...")
	stdscr.Refresh()

	time.Sleep(1 * time.Minute)

}
