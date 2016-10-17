package card

import (
	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/base/target"
	"github.com/kcnm/rocky/engine/effect"
)

type spellSpec struct {
	name    string
	class   base.Class
	mana    int
	assign  target.Assign
	side    target.Side
	role    target.Role
	effects []base.Effect
}

func (c spell) Class() base.Class {
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

func (c spell) Effects() []base.Effect {
	return spells[c].effects
}

func (c spell) String() string {
	return spells[c].name
}

var spells = map[spell]*spellSpec{
	Fireball: &spellSpec{
		name:   "Fireball",
		class:  base.Mage,
		mana:   4,
		assign: target.Manual,
		side:   target.Any,
		role:   target.Character,
		effects: []base.Effect{
			effect.DealDamage(6),
		},
	},
}
