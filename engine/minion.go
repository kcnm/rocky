package engine

type Minion interface {
	Char

	Card() MinionCard
}
