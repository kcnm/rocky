package engine

type Card interface {
	Class() Class
	Mana() int
}

type MinionCard interface {
	Card

	Attack() int
	Health() int
	Buff() Buff
}

type SpellCard interface {
	Card

	Effect() Effect
}

type WeaponCard interface {
	Card

	Attack() int
	Durability() int
}
