package card

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect"
)

type weaponSpec struct {
	class      engine.Class
	mana       int
	attack     int
	durability int
	battlecry  engine.Effect
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

func (s weaponSpec) Battlecry() engine.Effect {
	e := s.battlecry
	if e == nil {
		return effect.None
	} else {
		return e
	}
}

var weapons = map[string]*weaponSpec{
	fieryWarAxe: &weaponSpec{
		class:      engine.Warrior,
		mana:       2,
		attack:     3,
		durability: 2,
	},
}
