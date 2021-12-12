package strutil

import (
	"testing"
)

func Test_FitToTerm_Ascii(t *testing.T) {
	var source string = "Environment"
	var width uint = 7
	var actual string = FitToTerm(source, width, TextAlignLeft, false)

	var expected string = "Env~ent"
	if actual != expected {
		t.Fatalf("wrong response: actual %+v expected = %+v", actual, expected)
	}
}

func Test_FitToTerm_Unicode(t *testing.T) {
	var source string = "Перемещение"
	var width uint = 7
	var actual string = FitToTerm(source, width, TextAlignLeft, false)

	var expected string = "Пер~ние"
	if actual != expected {
		t.Fatalf("wrong response: actual %+v expected = %+v", actual, expected)
	}

}
