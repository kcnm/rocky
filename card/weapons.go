package card

import (
	"github.com/kcnm/rocky/engine"
)

type weaponSpec struct {
	class      engine.Class
	mana       int
	attack     int
	durability int
}

func (s weaponSpec) Class() engine.Class {
	return s.class
}

func (s weaponSpec) Mana() int {
	return s.mana
}

func (s weaponSpec) Attack() int {
	return s.attack
}

func (s weaponSpec) Durability() int {
	return s.durability
}

var weapons = map[string]*weaponSpec{
	fieryWarAxe: &weaponSpec{
		class:      engine.Warrior,
		mana:       2,
		attack:     3,
		durability: 2,
	},
}
