package widget

const (
	BUTTONBAR_LABELS_NUM = 10
)

type ButtonBarLabel struct {
	text      string
	command   uint
	end_coord uint
	receiver  IWidget
}

type ButtonBarWidget struct {
	labels []ButtonBarLabel
}

func NewButtonBarWidget() *ButtonBarWidget {

	bbw := &ButtonBarWidget{}

	//        let mut labels: Vec<ButtonBarLabel> = Vec::new();

	//        for _idx in 0..BUTTONBAR_LABELS_NUM {
	//            labels.push(ButtonBarLabel{
	//                text: String::new(),
	//                command: 0,
	//                end_coord: 0,
	//                receiver: None,
	//            });
	//        }

	//        ButtonBarWidget {
	//            labels,
	//        }

	return bbw
}

func (self *ButtonBarWidget) SetLabel(key uint, label string) {
	var keyIndex uint = key - 1
	self.labels[keyIndex].text = label
}

func (self *ButtonBarWidget) initButtonPositions(area *Rect) {
	//        debug!("Initialize button positions: area = {:?}", area);
	//        let mut pos: u16 = 0;
	//        let min_weight: u16 = BUTTONBAR_LABELS_NUM as u16 * 7;
	//        if area.width < min_weight {
	//        for i in 0..BUTTONBAR_LABELS_NUM as usize {
	//                if pos + 7 <= area.width {
	//                    pos += 7;
	//                }
	//                self.labels[i].end_coord = pos;
	//            }
	//        } else {
	//            let dv: u16 = area.width / BUTTONBAR_LABELS_NUM as u16;
	//            let md: u16 = area.width % BUTTONBAR_LABELS_NUM as u16;
	//
	//            let part_size: usize = BUTTONBAR_LABELS_NUM as usize / 2;
	//
	//            for i in 0..part_size {
	//                pos += dv;
	//                if BUTTONBAR_LABELS_NUM as u16 / 2 - 1 - (i as u16) < md / 2 {
	//                    pos += 1;
	//                }
	//                self.labels[i].end_coord = pos;
	//            }
	//
	//            for i in part_size..BUTTONBAR_LABELS_NUM as usize {
	//                pos += dv;
	//                if BUTTONBAR_LABELS_NUM as u16 - 1 - (i as u16) < (md + 1) / 2 {
	//                    pos += 1;
	//                }
	//                self.labels[i].end_coord = pos;
	//            }
	//        }
}

func (self *ButtonBarWidget) getButtonWidth(idx uint) uint {
	//        if idx == 0 {
	//            return self.labels[0].end_coord as usize;
	//        } else {
	//            let size = self.labels[idx].end_coord - self.labels[idx - 1].end_coord;
	//            return size as usize;
	//        }
	//    }
	return 0
}

func (self *ButtonBarWidget) ProcessInput(ch rune) {
}

func (self *ButtonBarWidget) Render(area *Rect) {
	//        // Step 1. Initialize button positions
	//        self.init_button_positions(area);
	//        //
	//        debug!("render: area = {:?}", area);
	//        // Step 0. Last screen
	//        print!("{}", termion::cursor::Goto(1, area.height));
	//        // Step 1. Draw bar
	//        for idx in 0..10 {
	//            let key = idx + 1;
	//            // Step 1. Draw index 1, 2, 3, .., 10
	//            // TODO - set black color ...
	//            let mut out = format!("{}", key);
	//            while out.chars().count() < 2 {
	//                out = format!(" {}", out);
	//            }
	//            print!("{}{}", termion::color::Fg(White), termion::color::Bg(Black));
	//            print!("{}", out);
	//            // Step 2. Draw key summary
	//            let width = self.get_button_width(idx);
	//            let summary_width = width - 2;
	//            let summary = &self.labels[idx];
	//            let mut out = format!("{}", summary.text);
	//            // Shrink
	//            if out.chars().count() > summary_width {
	//                out = fit_to_term(&out, summary_width, TextAlignMode::LEFT, false);
	//            }
	//            // Padding
	//            while out.chars().count() < summary_width {
	//                out.push(' ');
	//            }
	//            // Draw
	//            print!("{}{}", termion::color::Fg(Black), termion::color::Bg(Cyan));
	//            print!("{}", out);
	//        }
	//    }
}
