package ctx

import "github.com/vit1251/skyline-commander/tty"

var mainTerm *tty.PTerm

func GetTerm() *tty.PTerm {
	return mainTerm
}

func SetTerm(pTerm2 *tty.PTerm) {
	mainTerm = pTerm2
}
