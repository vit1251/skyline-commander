package widget

type Scoreboard struct {
	menu		int
	bar		int
	groups		int
	keysym		int
}

func NewScoreboard() *Scoreboard {
	return new(Scoreboard)
}

