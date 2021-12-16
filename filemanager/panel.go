package filemanager

import (
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/strutil"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/util"
	"github.com/vit1251/skyline-commander/widget"
)

type PanelWidget struct {
	widget.Widget
	active         bool /* If panel currently selected    */
	dirty          bool /* Should we redisplay the panel? */
	dir            int  /**/
	dirStat        int  /**/
	cwdPath        string
	listFormat     int
	format         string
	userFormat     string
	listCols       int
	briefCols      int
	sortInfo       int
	sortField      string
	marked         int
	dirsMarked     int
	total          uint64
	selected       int
	statusFormat   string
	userMiniStatus bool
	contentShift   int
}

func NewPanelWidget() *PanelWidget {
	pw := &PanelWidget{
		dirty:   true,
		active:  true,
		cwdPath: "/",
	}
	return pw
}

func (self *PanelWidget) GetPanelItems() int {
	return 0
}

func (self *PanelWidget) processCallback(msg widget.WidgetMsg) {
	switch msg {
	case widget.MsgDraw:
		self.Draw()
		//default:
		//	self.Widget.callback(msg)
	}
}

func (self *PanelWidget) Draw() {
	self.Widget.Erase()
	self.showDir()
	self.printHeader()
	self.adjustTopFile()
	self.paintDir()
	self.miniInfoSeparator()
	self.displayMiniInfo()
	self.dirty = false
}

func (self *PanelWidget) GetCorrectPathToShow() string {
	return self.cwdPath
}

func (self *PanelWidget) showDir() {

	mainTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()

	self.SetColors()
	mainTerm.DrawBox(self.Y, self.X, self.Lines, self.Cols, true)

	var show_mini_info bool = false // TODO - panels_options.show_mini_info
	if show_mini_info {
		var drawRune rune
		posY := self.GetLines() //+ 2

		self.Widget.GotoYX(posY, 0)
		drawRune = mainTerm.GetAltChar(tty.FRM_LTEE, true) // tty_print_alt_char(ACS_LTEE, FALSE)
		mainTerm.Print(fmt.Sprintf("%c", drawRune))

		self.Widget.GotoYX(posY, self.Cols-1)
		drawRune = mainTerm.GetAltChar(tty.FRM_RTEE, true) // tty_print_alt_char(ACS_LTEE, FALSE)
		mainTerm.Print(fmt.Sprintf("%c", drawRune))
	}

	reverseColorIndex := mainSkin.GetColor("core", "reverse")
	if self.active {
		mainTerm.ColorOn(reverseColorIndex)
	}

	/* Draw path */
	var newPath string = self.GetCorrectPathToShow()
	pathWeight := util.MIN(util.MAX(self.Cols-12, 0), self.Cols)
	newPath = fmt.Sprintf(" %s ", newPath)
	newPath = strutil.FitToTerm(newPath, uint(pathWeight), 0, false)
	self.Widget.GotoYX(0, 3)
	mainTerm.Print(newPath)

	//if !panels_options.show_mini_info {
	//	if panel- > marked == 0 {
	//		/* Show size of curret file in the bottom of panel */
	//		if S_ISREG(panel- > dir.list[panel- > selected].st.st_mode) {
	//			char
	//			buffer[BUF_SMALL]
	//
	//			g_snprintf(buffer, sizeof(buffer), " %s ",
	//				size_trunc_sep(panel- > dir.list[panel- > selected].st.st_size,
	//					panels_options.kilobyte_si))
	//			tty_setcolor(NORMAL_COLOR)
	//			widget_gotoyx(w, w- > lines-1, 4)
	//			tty_print_string(buffer)
	//		}
	//	} else
	//	{
	//		/* Show total size of marked files
	//		 * In the bottom of panel, display size only. */
	//		display_total_marked_size(panel, w- > lines-1, 2, TRUE)
	//	}
	//}

	self.showFreeSpace()

	if self.active {
		mainTerm.ColorOff(reverseColorIndex)
	}

}

func (self *PanelWidget) printHeader() {

	mainTerm := ctx.GetTerm()

	/* Erase */
	self.Widget.GotoYX(1, 1)
	curY, curX := mainTerm.GetYX()
	//tty_setcolor (NORMAL_COLOR);
	mainTerm.DrawHLine(curY, curX, ' ', self.Cols-2)
}

func (self *PanelWidget) adjustTopFile() {

}

func (self *PanelWidget) paintDir() {

}

func (self *PanelWidget) miniInfoSeparator() {

}

func (self *PanelWidget) displayMiniInfo() {

}

func (self *PanelWidget) SetColors() {
	mainSkin := ctx.GetSkin()
	mainTerm := ctx.GetTerm()
	var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
	mainTerm.ColorOn(normalColorIndex)
}

func (self *PanelWidget) showFreeSpace() {

	mainTerm := ctx.GetTerm()

	/* TODO - make status on remote system ... */

	/* Determine disk space */
	var myStatFs util.MyStatFs
	err1 := util.GetStatus(&myStatFs, self.cwdPath)
	if err1 != nil {
		panic(err1)
	}

	/* Show */
	newAvail := humanize.Bytes(myStatFs.Avail)
	newTotal := humanize.Bytes(myStatFs.Total)
	var percent uint64 = 0
	if myStatFs.Total != 0 {
		percent = 100 * myStatFs.Avail / myStatFs.Total
	}
	status := fmt.Sprintf(" %s/%s (%d%%) ", newAvail, newTotal, percent)
	statusLen := len(status)
	self.Widget.GotoYX(self.Lines-1, self.Cols-2-statusLen)
	// TODO - tty_setcolor (NORMAL_COLOR);
	mainTerm.Print(status)

}

// GetLines extract the number of available lines in a panel
func (self *PanelWidget) GetLines() int {
	/* 3 lines are: top frame, column header, bottom frame */
	// TODO - return self.lines - 3 - (panels_options.show_mini_info ? 2 : 0))
	return self.Lines - 3
}

// GetItemCount returns the number of items in the given panel
func (self *PanelWidget) GetItemCount() int {
	return self.GetLines() * self.listCols
}
