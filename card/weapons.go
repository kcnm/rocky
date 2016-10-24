package card

import (
	"github.com/kcnm/rocky/engine"
)

type weaponSpec struct {
	name       string
	class      engine.Class
	mana       int
	attack     int
	durability int
}

func (c weapon) Class() engine.Class {
	return weapons[c].class
}

func (c weapon) Mana() int {
	return weapons[c].mana
}

func (c weapon) Attack() int {
	return weapons[c].attack
}

func (c weapon) Durability() int {
	return weapons[c].durability
}

func (c weapon) String() string {
	return weapons[c].name
}

var weapons = map[weapon]*weaponSpec{
	FieryWarAxe: &weaponSpec{
		name:       "Fiery War Axe",
		class:      engine.Warrior,
		mana:       2,
		attack:     3,
		durability: 2,
	},
}
