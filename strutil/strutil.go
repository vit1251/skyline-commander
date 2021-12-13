package strutil

import (
	"fmt"
	"log"
	"unicode/utf8"
)

type TextAlignMode int

const (
	TextAlignLeft   = 0x01
	TextAlignCenter = 0x02
	TextAlignRight  = 0x03
)

func FitToTerm(src string, width uint, just_mode TextAlignMode, fit bool) string {

	var result string = src

	var inputSize uint = uint(utf8.RuneCountInString(src))
	log.Printf("FitToTerm: string = %s count = %d", src, inputSize)

	// Input so bigger
	if inputSize > width {
		var prefixCount uint = width / 2
		var suffixCount uint = width / 2

		if prefixCount+suffixCount == width {
			suffixCount = suffixCount - 1
		}

		var prefixPos uint = prefixCount
		var suffixPos uint = inputSize - suffixCount

		var prefix string = ""
		var suffix string = ""

		var runes []rune = []rune(src)
		for idx, ch := range runes {

			log.Printf("FitToTerm: idx = %d ch = %c", idx, ch)

			// Prefix
			if uint(idx) < prefixPos {
				prefix = fmt.Sprintf("%s%c", prefix, ch)
			}
			// Suffix
			if uint(idx) >= suffixPos {
				suffix = fmt.Sprintf("%s%c", suffix, ch)
			}
		}
		result = fmt.Sprintf("%s~%s", prefix, suffix)

		return result
	}

	return result
}
