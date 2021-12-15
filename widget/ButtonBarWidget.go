package widget

import (
	"fmt"
	"github.com/vit1251/skyline-commander/ctx"
	"github.com/vit1251/skyline-commander/skin"
	"github.com/vit1251/skyline-commander/strutil"
	"github.com/vit1251/skyline-commander/tty"
	"github.com/vit1251/skyline-commander/tty/event"
	"log"
	"unicode/utf8"
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
	Widget
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
	bbw.Widget.callback = bbw.processCallback

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

func (self *ButtonBarWidget) drawKey(pTerm *tty.PTerm, area *Rect, skin *skin.Skin, key uint) {
	/* Step 1. Draw key name */
	keyName := fmt.Sprintf("%d", key)
	for utf8.RuneCountInString(keyName) < 2 {
		keyName = fmt.Sprintf(" %s", keyName)
	}
	hotkeySkinColorIndex := skin.GetColor("buttonbar", "hotkey")
	pTerm.ColorOn(hotkeySkinColorIndex)
	pTerm.Print(keyName)
	pTerm.ColorOff(hotkeySkinColorIndex)

	/* Step 2. Draw key summary */
	width := self.getButtonWidth(key)
	var summaryWidth int = int(width) - 2
	keyLabel := self.getLabel(key)
	keySummary := fmt.Sprintf("%s", keyLabel.text)
	// Shrink
	if utf8.RuneCountInString(keySummary) > summaryWidth {
		keySummary = strutil.FitToTerm(keySummary, uint(summaryWidth), strutil.TextAlignLeft, false)
	}
	// Padding
	for utf8.RuneCountInString(keySummary) < summaryWidth {
		keySummary = fmt.Sprintf("%s ", keySummary)
	}
	// Draw
	buttonSkinColorIndex := skin.GetColor("buttonbar", "button")
	pTerm.ColorOn(buttonSkinColorIndex)
	pTerm.Print(keySummary)
	pTerm.ColorOff(buttonSkinColorIndex)

	/* Debug message */
	log.Printf("Render: key = %d summary = %q", key, keySummary)

}

func (self *ButtonBarWidget) Draw() {

	pTerm := ctx.GetTerm()
	mainSkin := ctx.GetSkin()
	maxY, maxX := pTerm.MaxYX()
	area := NewRect(0, 0, uint(maxX), uint(maxY))

	/* Step 1. Initialize button positions */
	self.initButtonPositions(area)
	self.dumpButtonPositions()

	/* Step 0. Detect bottom position */
	pTerm.Move(int(area.height)-1, 0)

	/* Step 1. Draw bar */
	for _, key := range []uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		self.drawKey(pTerm, area, mainSkin, key)
	}

}

func (self *ButtonBarWidget) processCallback(msg WidgetMsg) {
	switch msg {
	case MsgDraw:
		self.Draw()
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

func (self *ButtonBarWidget) dumpButtonPositions() {
	for index, label := range self.labels {
		log.Printf("Label: index = %d key = %d summary = %q endCoord = %d", index, label.key, label.text, label.endCoord)
	}
}
