package tty

import (
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type PTerm struct {
	running bool
	stdscr  *goncurses.Window
	C       chan event.Event
}

func NewPTerm() *PTerm {
	pTerm := &PTerm{}
	return pTerm
}

func (self *PTerm) Init() error {

	/* Process SIGWINCH system event */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGWINCH)

	/* Disable GetText encoding */
	//res := SetLocale.SetLocale(SetLocale.LC_ALL, "")
	//log.Printf("SetLocale: res = %+v", res)

	log.Printf("Initialize ncurses: version = %s", goncurses.CursesVersion())

	stdscr, err1 := goncurses.Init()
	if err1 != nil {
		return err1
	}
	self.stdscr = stdscr

	goncurses.CBreak(true)
	goncurses.Raw(true)
	goncurses.Echo(false)
	err2 := goncurses.Cursor(0)
	if err2 != nil {
		return err2
	}
	err3 := stdscr.Keypad(true)
	if err3 != nil {
		return err3
	}

	stdscr.Timeout(100)

	err4 := goncurses.StartColor()
	if err4 != nil {
		return err4
	}

	console := make(chan int, 1)

	self.running = true

	go func() {

		for self.running {
			/* Process input */
			var key goncurses.Key = stdscr.GetChar()
			if key != 0 {
				console <- int(key)
			}
		}

		close(console)

	}()

	self.C = make(chan event.Event)

	go func() {
		for self.running {
			select {
			case <-sigs:
				log.Printf("Linux console resize event")
				self.updateTermSize()
				evt := event.NewEvent()
				evt.EvType = event.EventTypeResize
				self.C <- *evt
			case key := <-console:
				log.Printf("Console event: ch = %d", key)
				evt := event.NewEventFromKey(key)
				log.Printf("Event: evt = %+v", evt)
				self.C <- *evt
				// TODO - mouse ...
			}
		}
	}()

	return nil
}

func (self *PTerm) End() {
	goncurses.End()
}

func (self *PTerm) Erase() {
	self.stdscr.Erase()
}

func (self *PTerm) Refresh() {
	self.stdscr.Refresh()
}

func (self *PTerm) MaxYX() (int, int) {
	return self.stdscr.MaxYX()
}

func (self *PTerm) Move(y int, x int) {
	self.stdscr.Move(y, x)
}

func (self *PTerm) Print(msg string) {
	self.stdscr.Print(msg)
}

func (self *PTerm) updateTermSize() error {

	prevHeight, prevWidth := self.stdscr.MaxYX()
	height, width, err1 := osTermSize()
	if err1 != nil {
		return err1
	}

	/* Debug message */
	log.Printf("prevWidth = %d prevHeight = %d width = %d height = %d",
		prevWidth, prevHeight,
		width, height,
	)

	/* Set new PTerm size */
	err2 := goncurses.ResizeTerm(height, width)
	if err2 != nil {
		return err2
	}
	self.stdscr.ClearOk(true)

	return nil
}

var nextPairIndex int16 = 1

func colorToIndex(colorName string) int16 {
	var colorIndex int16 = 0
	if colorName == "black" {
		colorIndex = goncurses.C_BLACK
	} else if colorName == "blue" {
		colorIndex = goncurses.C_BLUE
	} else if colorName == "lightgray" {
		colorIndex = goncurses.C_WHITE
	} else if colorName == "white" {
		colorIndex = goncurses.C_WHITE // | goncurses.
	} else if colorName == "cyan" {
		colorIndex = goncurses.C_CYAN
	} else if colorName == "yellow" {
		colorIndex = goncurses.C_YELLOW
	} else {
		log.Panicf("wong color name %s", colorName)
	}
	return colorIndex
}

func (self *PTerm) InitColor(fg string, bg string) skin.ColorPair {
	pairIndex := nextPairIndex
	bgColorIndex := colorToIndex(bg)
	fgColorIndex := colorToIndex(fg)
	err1 := goncurses.InitPair(pairIndex, fgColorIndex, bgColorIndex)
	if err1 != nil {
		panic("init ncurses color error")
	}
	nextPairIndex += 1

	return skin.ColorPair(pairIndex)
}

func (self *PTerm) ColorOn(pair skin.ColorPair) {
	log.Printf("PTerm: set color: pair = %d", pair)
	self.stdscr.ColorOn(int16(pair))
}
func (self *PTerm) ColorOff(pair skin.ColorPair) {
	log.Printf("PTerm: set color: pair = %d", pair)
	self.stdscr.ColorOff(int16(pair))
}

func (self *PTerm) Touch() {
	self.stdscr.Touch()
}

func (self *PTerm) FillRegion(x int, y int, lines int, cols int, ch rune) {
	//TODO - self.stdscr.Re
}

func (self *PTerm) DrawBox(y int, x int, lines int, cols int, b bool) {
	// TODO -
}
