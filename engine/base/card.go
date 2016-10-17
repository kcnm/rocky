package base

import (
	"github.com/kcnm/rocky/engine/base/target"
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

	Effects() []Effect
}

type WeaponCard interface {
	Card

	Attack() int
	Durability() int
}
