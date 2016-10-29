package engine

import (
	"github.com/kcnm/rocky/engine/target"
)

type Card interface {
	Class() Class
	Mana() int
}

type MinionCard interface {
	Card

	Attack() int
	Health() int
}

type SpellCard interface {
	Card
	target.Spec

	Effect() Effect
}

type WeaponCard interface {
	Card

	Attack() int
	Durability() int
}
