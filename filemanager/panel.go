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
	"log"
)

type FormatItem struct {
	requestedFieldLen int /* Requested Field len */
	fieldLen          int
	expand            bool
	stringFn          func() string /* Format FileEntry */
	title             string
	id                string
}

type FileEntry struct {
	Name string
}

type PanelWidget struct {
	widget.Widget
	active         bool         /* If panel currently selected    */
	dirty          bool         /* Should we redisplay the panel? */
	dir            []*FileEntry /**/
	dirStat        int          /**/
	cwdPath        string
	listFormat     int
	format         []*FormatItem
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
	showMiniInfo   bool
	topFile        int
}

func NewPanelWidget() *PanelWidget {
	pw := &PanelWidget{
		dirty:        true,
		active:       true,
		cwdPath:      "/",
		showMiniInfo: true,
	}

	/* Set up views */
	pw.setupBriefFormat()

	return pw
}

func NewFormatItem(id string, title string, requestedFieldLen int, expand bool) *FormatItem {
	fi := &FormatItem{
		id:                id,
		title:             title,
		requestedFieldLen: requestedFieldLen,
		expand:            expand,
	}
	return fi
}

func (self *PanelWidget) setupBriefFormat() {

	self.format = nil

	/* Name column */
	nameItem := NewFormatItem("name", "Name", 0, true)
	self.format = append(self.format, nameItem)

	sizeItem := NewFormatItem("size", "Size", 7, false)
	self.format = append(self.format, sizeItem)

	modItem := NewFormatItem("mtime", "Modify time", 12, false)
	self.format = append(self.format, modItem)

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

	/* Draw box */
	var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
	mainTerm.ColorOn(normalColorIndex)
	mainTerm.DrawBox(self.Y, self.X, self.Lines, self.Cols, true)
	if self.showMiniInfo {
		var drawRune rune
		posY := self.Lines - 3

		self.Widget.GotoYX(posY, 0)
		drawRune = mainTerm.GetAltChar(tty.FRM_LTEE, true) // tty_print_alt_char(ACS_LTEE, FALSE)
		mainTerm.Print(fmt.Sprintf("%c", drawRune))

		self.Widget.GotoYX(posY, self.Cols-1)
		drawRune = mainTerm.GetAltChar(tty.FRM_RTEE, true) // tty_print_alt_char(ACS_LTEE, FALSE)
		mainTerm.Print(fmt.Sprintf("%c", drawRune))
	}
	mainTerm.ColorOff(normalColorIndex)

	/* Draw path */
	var reverseColorIndex skin.ColorPair = mainSkin.GetColor("core", "reverse")
	if self.active {
		mainTerm.ColorOn(reverseColorIndex)
	}
	var newPath string = self.GetCorrectPathToShow()
	pathWeight := util.MIN(util.MAX(self.Cols-12, 0), self.Cols)
	newPath = fmt.Sprintf(" %s ", newPath)
	newPath = strutil.FitToTerm(newPath, uint(pathWeight), 0, false)
	self.Widget.GotoYX(0, 3)
	mainTerm.Print(newPath)
	if self.active {
		mainTerm.ColorOff(reverseColorIndex)
	}

	if self.showMiniInfo {
		/* Show total size of marked files In the bottom of panel, display size only. */
		self.displayTotalMarkedSize(self.Lines-1, 2, true)
	}

	/* Disk summary */
	self.showFreeSpace()

}

func (self *PanelWidget) printHeader() {

	mainTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()

	/* Erase */
	var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
	mainTerm.ColorOn(normalColorIndex)
	mainTerm.DrawHLine(self.Y+1, self.X+1, ' ', self.Cols-2)
	mainTerm.ColorOff(normalColorIndex)

	/* Draw calculate size */
	var summaryRequestedFieldsLen int = 0
	for _, row := range self.format {
		summaryRequestedFieldsLen = summaryRequestedFieldsLen + row.requestedFieldLen
	}
	expandFieldsLen := self.Cols - summaryRequestedFieldsLen
	for _, row := range self.format {
		if row.expand {
			row.fieldLen = expandFieldsLen
		} else {
			row.fieldLen = row.requestedFieldLen
		}
	}

	/* Draw headers */
	var headerColorIndex skin.ColorPair = mainSkin.GetColor("core", "header")
	mainTerm.ColorOn(headerColorIndex)
	var startX int = 1
	var stopX int = 1
	for _, row := range self.format {
		/* Next */
		startX = stopX
		stopX = startX + row.fieldLen - 1
		/* Render */
		log.Printf("row = %+v size = %d", row, row.fieldLen)
		newTitle := strutil.FitToTerm(row.title, uint(stopX-startX), 0, true)

		mainTerm.GotoYX(self.Y+1, self.X+startX)
		mainTerm.Print(newTitle)
	}
	mainTerm.ColorOff(headerColorIndex)

}

// adjustTopFile is update panel->selected value to avoid out of range
func (self *PanelWidget) adjustTopFile() {
	var panelItemCount int = len(self.dir)
	if self.selected > panelItemCount {
		self.selected = panelItemCount
	}
}

func (self *PanelWidget) paintDir() {

	mainTerm := ctx.GetTerm()

	var itemCount int = self.Lines
	itemCount = itemCount - 1 // use by box
	itemCount = itemCount - 1 // use by header row
	if self.showMiniInfo {
		itemCount = itemCount - 1 // use by box
		itemCount = itemCount - 1 // use by mini info panel
	}
	itemCount = itemCount - 1 // use by box
	log.Printf("paintDir: itemCount = %d", itemCount)

	/* Draw directory entries items */
	var itemIndex int
	for itemIndex = 0; itemIndex < itemCount; itemIndex++ {
		/* Draw background */
		mainTerm.GotoYX(self.Y+2+itemIndex, self.X+1)
		mainTerm.Print(fmt.Sprintf("%d", itemIndex))

		/* Draw meta */
		var colorIndex skin.ColorPair = 0
		self.repaintFile(itemIndex+self.topFile, colorIndex, false)
	}
}

func (self *PanelWidget) miniInfoSeparator() {

	mainTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()

	if self.showMiniInfo {
		var posY int = self.Lines - 3
		var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
		mainTerm.ColorOn(normalColorIndex)
		fillRune := mainTerm.GetAltChar(tty.FRM_HORIZ, true)
		mainTerm.DrawHLine(self.Y+posY, self.X+1, fillRune, self.Cols-2)
		mainTerm.ColorOff(normalColorIndex)

		/* Status displays total marked size. Centered in panel, full format. */
		//self.displayTotalMarkedSize (posY, -1, false)
	}
}

func (self *PanelWidget) displayMiniInfo() {

	mainTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()

	if self.showMiniInfo {
		/* Fill space */
		var posY int = self.Lines - 2
		var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
		mainTerm.ColorOn(normalColorIndex)
		mainTerm.DrawHLine(self.Y+posY, self.X+1, ' ', self.Cols-2)
		mainTerm.ColorOff(normalColorIndex)
		/* Link source */
		/* Dot directory */
		/* File name */
	}
}

func (self *PanelWidget) showFreeSpace() {

	mainTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()

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
	var normalColorIndex skin.ColorPair = mainSkin.GetColor("core", "_default_")
	mainTerm.ColorOn(normalColorIndex)
	mainTerm.Print(status)
	mainTerm.ColorOff(normalColorIndex)

}

func (self *PanelWidget) displayTotalMarkedSize(i int, i2 int, b bool) {
	// TODO - implement it later...
}

func (self *PanelWidget) repaintFile(itemIndex int, colorIndex skin.ColorPair, showStatus bool) {

}
