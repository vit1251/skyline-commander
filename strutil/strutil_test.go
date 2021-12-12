package strutil

func test_FitToTerm() {
	var source string = "Перемещение"
	var width int = 7
        var expected string = "Пер~ние"

        var actual string = FitToTerm(source, width, TEXT_ALIGN_LEFT, false)

//        assert_eq!(actual, expected);

}
