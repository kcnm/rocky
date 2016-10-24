package engine

type Minion interface {
	Character

	Card() MinionCard
}
