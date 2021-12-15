package widget

import (
	"fmt"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/strutil"
	"github.com/vit1251/skyline-commander/tty"
)

type PanelWidget struct {
	Widget
	active         bool /* If panel currently selected    */
	dirty          bool /* Should we redisplay the panel? */
	dir            int  /**/
	dirStat        int  /**/
	cwdPath        string
	lwdPath        string
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
		dirty: true,
	}
	return pw
}

func (self *PanelWidget) GetPanelItems() int {
	return 0
}

func (self *PanelWidget) processCallback(msg WidgetMsg) {
	switch msg {
	case MsgDraw:
		self.render(nil, nil, nil)
	default:
		self.defaultCallback(msg)
	}
}

func (self *PanelWidget) render(pTerm *tty.PTerm, area *Rect, skin *skin.Skin) {
	self.Widget.Erase()
	self.showDir()
	self.printHeader()
	self.adjustTopFile()
	self.paintDir()
	self.miniInfoSeparator()
	self.displayMiniInfo()
	self.dirty = false
}

func (self *PanelWidget) showDir() {

	pTerm := ctx.GetTerm()
	skin := ctx.GetSkin()

	//set_colors (panel)
	pTerm.DrawBox(self.y, self.x, self.lines, self.cols, false)

	//if panels_options.show_mini_info {
	//	int y
	//
	//	y = panel_lines(panel) + 2
	//
	//	widget_gotoyx(w, y, 0)
	//	tty_print_alt_char(ACS_LTEE, FALSE)
	//	widget_gotoyx(w, y, w- > cols-1)
	//	tty_print_alt_char(ACS_RTEE, FALSE)
	//}

	/* Show ??? */
	//self.Widget.GotoYX(0, 1)
	//pTerm.Print(panel_history_prev_item_char)

	//	tmp = panels_options.show_dot_files ? panel_hiddenfiles_show_char :	panel_hiddenfiles_hide_char
	//	tmp = g_strdup_printf("%s[%s]%s", tmp, panel_history_show_list_char, panel_history_next_item_char)

	//	widget_gotoyx(w, 0, w- > cols-6)
	//	tty_print_string(tmp)

	self.Widget.GotoYX(0, 3)

	reverseColorIndex := skin.GetColor("", "") // REVERSE_COLOR
	if self.active {
		pTerm.ColorOn(reverseColorIndex)
	}

	/* Show path */
	// TODO - tmp = panel_correct_path_to_show(panel)
	var newPath string = "/var/lib"
	//pathWeight := MIN(MAX(w- > cols-12, 0), w- > cols))
	var pathWeight uint = 20
	newPath = fmt.Sprintf(" %s ", newPath)
	newPath = strutil.FitToTerm(newPath, pathWeight, 0, false)
	pTerm.Print(newPath)

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

	//show_free_space(panel)

	if self.active {
		pTerm.ColorOff(reverseColorIndex)
	}

}

func (self *PanelWidget) printHeader() {

}

func (self *PanelWidget) adjustTopFile() {

}

func (self *PanelWidget) paintDir() {

}

func (self *PanelWidget) miniInfoSeparator() {

}

func (self *PanelWidget) displayMiniInfo() {

}
