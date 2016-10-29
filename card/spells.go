package card

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect"
	"github.com/kcnm/rocky/engine/target"
)

type spellSpec struct {
	class  engine.Class
	mana   int
	assign target.Assign
	side   target.Side
	role   target.Role
	effect engine.Effect
}

func (c spell) Class() engine.Class {
	return spells[c].class
}

func (c spell) Mana() int {
	return spells[c].mana
}

func (c spell) Assign() target.Assign {
	return spells[c].assign
}

func (c spell) Side() target.Side {
	return spells[c].side
}

func (c spell) Role() target.Role {
	return spells[c].role
}

func (c spell) Effect() engine.Effect {
	return spells[c].effect
}

var spells = map[spell]*spellSpec{
	Fireball: &spellSpec{
		class:  engine.Mage,
		mana:   4,
		assign: target.Manual,
		side:   target.Any,
		role:   target.Char,
		effect: effect.DealDamage(6),
	},
}
