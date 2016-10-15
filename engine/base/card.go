package base

type Card interface {
	Class() Class
	Mana() int
}

type MinionCard interface {
	Card

	InitialAttack() int
	InitialHealth() int
}

type SpellCard interface {
	Card
}
