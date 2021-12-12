package widget

import (
	"fmt"
	"github.com/gbin/goncurses"
	"github.com/vit1251/skyline-commander/strutil"
	"github.com/vit1251/skyline-commander/widget/event"
	"log"
)

const (
	BUTTONBAR_LABELS_NUM = 10
)

type ButtonBarLabel struct {
	key      uint
	text     string
	command  uint
	endCoord uint
	receiver IWidget
}

type ButtonBarWidget struct {
	IWidget
	labels []ButtonBarLabel
}

func NewButtonBarWidget() *ButtonBarWidget {

	var labels []ButtonBarLabel

	for _, key := range []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		labels = append(labels, ButtonBarLabel{
			key:      key,
			text:     "",
			command:  0,
			endCoord: 0,
			receiver: nil,
		})
	}

	bbw := &ButtonBarWidget{
		labels: labels,
	}

	return bbw
}

func (self *ButtonBarWidget) SetLabel(key uint, label string) {
	var keyIndex uint = key - 1
	self.labels[keyIndex].text = label
}

func (self *ButtonBarWidget) initButtonPositions(area *Rect) {
	log.Printf("initButtonPositions: area = %v", area)
	var pos uint = 0
	var minWeight uint = BUTTONBAR_LABELS_NUM * 7
	if area.width < minWeight {
		for idx, _ := range []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
			if pos+7 <= area.width {
				pos += 7
			}
			self.labels[idx].endCoord = pos
		}
	} else {
		var dv uint = area.width / BUTTONBAR_LABELS_NUM
		var md uint = area.width % BUTTONBAR_LABELS_NUM

		var partSize uint = BUTTONBAR_LABELS_NUM / 2

		for i := 0; i < int(partSize); i++ {
			pos += dv
			if BUTTONBAR_LABELS_NUM/2-1-uint(i) < md/2 {
				pos += 1
			}
			self.labels[i].endCoord = pos
		}

		for i := partSize; i < BUTTONBAR_LABELS_NUM; i++ {
			pos += dv
			if BUTTONBAR_LABELS_NUM-1-uint(i) < (md+1)/2 {
				pos += 1
			}
			self.labels[i].endCoord = pos
		}

	}
}

func (self *ButtonBarWidget) getButtonWidth(key uint) uint {

	var pervEndCord uint

	for _, label := range self.labels {
		if label.key == 1 && key == 1 {
			return label.endCoord
		}
		if label.key == key {
			size := label.endCoord - pervEndCord
			return size
		}
		pervEndCord = label.endCoord
	}

	return 0
}

func (self *ButtonBarWidget) ProcessEvent(evt *event.Event) {
}

func (self *ButtonBarWidget) drawKey(stdscr *goncurses.Window, area *Rect, key uint) {
	/* Step 1. Draw key name */
	keyName := fmt.Sprintf("%d", key)
	for len(keyName) < 2 {
		keyName = fmt.Sprintf(" %s", keyName)
	}
	//	stdscr.Color() //            print!("{}{}", termion::color::Fg(White), termion::color::Bg(Black));
	stdscr.Print(keyName)

	/* Step 2. Draw key summary */
	width := self.getButtonWidth(key)
	summaryWidth := width - 2
	keyLabel := self.getLabel(key)
	keySummary := fmt.Sprintf("%s", keyLabel.text)
	// Shrink
	if len(keySummary) > int(summaryWidth) {
		keySummary = strutil.FitToTerm(keySummary, summaryWidth, strutil.TextAlignLeft, false)
	}
	// Padding
	for len(keySummary) < int(summaryWidth) {
		keySummary = fmt.Sprintf("%s ", keySummary)
	}
	// Draw
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(Cyan));
	stdscr.Print(keySummary)

	/* Debug message */
	log.Printf("Render: key = %d summary = %s", key, keySummary)

}

func (self *ButtonBarWidget) Render(stdscr *goncurses.Window, area *Rect) {

	/* Step 1. Initialize button positions */
	self.initButtonPositions(area)

	/* Step 0. Detect bottom position */
	stdscr.Move(int(area.height)-1, 1)

	/* Step 1. Draw bar */
	for _, key := range []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		self.drawKey(stdscr, area, key)
	}

}

func (self *ButtonBarWidget) getLabel(key uint) *ButtonBarLabel {
	for _, label := range self.labels {
		if label.key == key {
			return &label
		}
	}
	panic("wrong key")
}
