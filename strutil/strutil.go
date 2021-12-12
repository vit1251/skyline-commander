package strutil

type TextAlignMode int

const (
    TEXT_ALIGN_LEFT    = 0x01,
    TEXT_ALIGN_CENTER  = 0x02,
    TEXT_ALIGN_RIGHT   = 0x03,
}

func FitToTerm(src string, width uint, just_mode TextAlignMode, fit bool) string {
//    let mut result = String::new();
//    let input_size = text.chars().count();
//    // Input so bigger
//    if input_size > width {
//        let prefix_count = width / 2;
//        let suffix_count = width / 2;
//        let suffix_count = if prefix_count + suffix_count == width {
//            suffix_count - 1
//        } else {
//            suffix_count
//        };
//        let prefix_pos = prefix_count;
//        let suffix_pos = input_size - suffix_count;
//        let mut prefix = String::new();
//        let mut suffix = String::new();
//        for (idx, ch) in text.chars().enumerate() {
            // Prefix
//            if idx < prefix_pos {
//                prefix.push(ch);
//            }
            // Suffix
//            if idx >= suffix_pos {
//                suffix.push(ch);
//            }
//        }
//        result.push_str(&prefix);
//        result.push('~');
//        result.push_str(&suffix);
//        return result;
//    }

    // Input is small

//    match just_mode {
//        TextAlignMode::LEFT => {
//            let out = format!("{}", text);
//            result.push_str(&out);
//        },
//        TextAlignMode::CENTER => {
//        },
//        TextAlignMode::RIGHT => {
//        }
//    }
//    result
}
