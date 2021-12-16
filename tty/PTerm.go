package tty

import (
	"fmt"
	ncursesw "github.com/vit1251/go-ncursesw"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type PTerm struct {
	running bool             /* Main event processing marker   */
	stdscr  *ncursesw.Window /* Use libncursesw C binding      */
	C       chan event.Event /* Channel with event             */
	resized bool             /* Resize marker                  */
}

func NewPTerm() *PTerm {
	pTerm := &PTerm{}
	return pTerm
}

func (self *PTerm) Init() error {

	/* Process SIGWINCH system event */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGWINCH)

	log.Printf("Initialize ncurses: version = %s", ncursesw.CursesVersion())

	stdscr, err1 := ncursesw.Init()
	if err1 != nil {
		return err1
	}
	self.stdscr = stdscr

	ncursesw.CBreak(true)
	ncursesw.Raw(true)
	ncursesw.Echo(false)
	err2 := ncursesw.Cursor(0)
	if err2 != nil {
		return err2
	}
	err3 := stdscr.Keypad(true)
	if err3 != nil {
		return err3
	}

	err4 := ncursesw.StartColor()
	if err4 != nil {
		return err4
	}

	console := make(chan int, 1)

	self.running = true

	go func() {
		for self.running {
			select {
			case <-sigs:
				self.resized = true
			}
		}
	}()

	resize := make(chan bool, 1)
	go func() {
		for self.running {
			time.Sleep(500 * time.Millisecond)
			if self.resized {
				resize <- true
				self.resized = false
			}
		}
		close(resize)
	}()

	go func() {

		runtime.LockOSThread()

		/* Set timeout */
		stdscr.Timeout(100)

		for self.running {

			/* Wait input */
			startWait := time.Now()
			log.Printf("waitInput in.")
			waitInput()
			log.Printf("waitInput is out.")
			log.Printf("waitInput is %q msec.", time.Since(startWait))

			/* Process input */
			var key ncursesw.Key = stdscr.GetChar()
			if key != 0 {
				console <- int(key)
			}
		}

		close(console)

	}()

	self.C = make(chan event.Event)

	go func() {

		runtime.LockOSThread()

		for self.running {
			select {
			case <-resize:
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
	ncursesw.End()
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

func (self *PTerm) GotoYX(y int, x int) {
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
	err2 := ncursesw.ResizeTerm(height, width)
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
		colorIndex = ncursesw.C_BLACK
	} else if colorName == "blue" {
		colorIndex = ncursesw.C_BLUE
	} else if colorName == "lightgray" {
		colorIndex = ncursesw.C_WHITE
	} else if colorName == "white" {
		colorIndex = ncursesw.C_WHITE // | goncurses.
	} else if colorName == "cyan" {
		colorIndex = ncursesw.C_CYAN
	} else if colorName == "yellow" {
		colorIndex = ncursesw.C_YELLOW
	} else {
		log.Panicf("wong color name %s", colorName)
	}
	return colorIndex
}

func (self *PTerm) InitColor(fg string, bg string) skin.ColorPair {
	pairIndex := nextPairIndex
	bgColorIndex := colorToIndex(bg)
	fgColorIndex := colorToIndex(fg)
	err1 := ncursesw.InitPair(pairIndex, fgColorIndex, bgColorIndex)
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

func (self *PTerm) FillRegion(x int, y int, rows int, cols int, ch rune) {

	// TODO -     if (!tty_clip (&y, &x, &rows, &cols))
	// TODO -         return;

	var i int
	for i = 0; i < rows; i++ {
		self.stdscr.HLine(y+i, x, ncursesw.Char(ch), cols)
	}

	self.Move(y, x)

}

const (
	FRM_VERT  = 0x0001
	FRM_HORIZ = 0x0002

	FRM_ULCORNER = 0x0101
	FRM_LLCORNER = 0x0102
	FRM_URCORNER = 0x0103
	FRM_LRCORNER = 0x0104

	FRM_LTEE = 0x0201
	FRM_RTEE = 0x0202
)

func (self *PTerm) DrawBox(y int, x int, ys int, xs int, single bool) {

	var drawRune rune
	var y2 int
	var x2 int

	if ys <= 0 || xs <= 0 {
		return
	}

	ys--
	xs--

	y2 = y + ys
	x2 = x + xs

	drawRune = self.GetAltChar(FRM_VERT, single) // mc_tty_frm[single ? MC_TTY_FRM_VERT : MC_TTY_FRM_DVERT]
	self.DrawVLine(y, x, drawRune, ys)
	self.DrawVLine(y, x2, drawRune, ys)

	drawRune = self.GetAltChar(FRM_HORIZ, single) // mc_tty_frm[single ? MC_TTY_FRM_HORIZ : MC_TTY_FRM_DHORIZ]
	self.DrawHLine(y, x, drawRune, xs)
	self.DrawHLine(y2, x, drawRune, xs)

	/* Draw upper left */
	self.GotoYX(y, x)
	drawRune = self.GetAltChar(FRM_ULCORNER, single)
	self.Print(fmt.Sprintf("%c", drawRune))

	/* Draw lower left*/
	self.GotoYX(y2, x)
	drawRune = self.GetAltChar(FRM_LLCORNER, single)
	self.Print(fmt.Sprintf("%c", drawRune))

	/* Draw upper right */
	self.GotoYX(y, x2)
	drawRune = self.GetAltChar(FRM_URCORNER, single)
	self.Print(fmt.Sprintf("%c", drawRune))

	/* Draw lower right */
	self.GotoYX(y2, x2)
	drawRune = self.GetAltChar(FRM_LRCORNER, single)
	self.Print(fmt.Sprintf("%c", drawRune))
}

func (self *PTerm) GetAltChar(runeName int, single bool) rune {
	if single {
		switch runeName {
		case FRM_VERT:
			return '│'
		case FRM_HORIZ:
			return '─'
		case FRM_LLCORNER:
			return '└'
		case FRM_URCORNER:
			return '┐'
		case FRM_ULCORNER:
			return '┌'
		case FRM_LRCORNER:
			return '┘'
		case FRM_LTEE:
			return '├'
		case FRM_RTEE:
			return '┤'
		}
	} else {
		switch runeName {
		case FRM_VERT:
			return '║'
		case FRM_HORIZ:
			return '═'
		case FRM_LLCORNER:
			return '╚'
		case FRM_URCORNER:
			return '╗'
		case FRM_ULCORNER:
			return '╔'
		case FRM_LRCORNER:
			return '╝'
		}
	}
	return '?'
}

func (self *PTerm) DrawVLine(y int, x int, ch rune, len int) {

	var y1 int = y

	maxY, maxX := self.MaxYX()

	if x < 0 || y < 0 || x >= maxX || y >= maxY {
		return
	}

	// TODO - self.stdscr.VLine(y, x, goncurses.Char(ch), len)
	var curY int
	for curY = y; curY < y+len && curY < maxY; curY++ {
		self.Move(curY, x)
		self.Print(fmt.Sprintf("%c", ch))
	}

	/* Restore position */
	self.Move(y1, x)

}

func (self *PTerm) DrawHLine(y int, x int, ch rune, len int) {

	var x1 int = x

	maxY, maxX := self.MaxYX()

	if x < 0 || y < 0 || y >= maxY || x >= maxX {
		return
	}

	//TODO - self.stdscr.HLine(y, x, goncurses.Char(ch), len)
	var curX int
	for curX = x; curX < x+len && curX < maxX; curX++ {
		self.Move(y, curX)
		self.Print(fmt.Sprintf("%c", ch))
	}

	self.GotoYX(y, x1)

}

func (self *PTerm) GetYX() (int, int) {
	return self.stdscr.CursorYX()
}
