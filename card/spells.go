package card

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect"
	"github.com/kcnm/rocky/engine/effect/choose"
	"github.com/kcnm/rocky/engine/effect/param"
	"github.com/kcnm/rocky/engine/effect/pred"
)

type spellSpec struct {
	class  engine.Class
	mana   int
	effect engine.Effect
}

func (c spell) Class() engine.Class {
	return spells[c].class
}

func (c spell) Mana() int {
	return spells[c].mana
}

func (c spell) Effect() engine.Effect {
	return spells[c].effect
}

var spells = map[spell]*spellSpec{
	Fireball: &spellSpec{
		class:  engine.Mage,
		mana:   4,
		effect: effect.DealDamage(6, param.Char(choose.Manual, pred.Char)),
	},
	Flamestrike: &spellSpec{
		class:  engine.Mage,
		mana:   7,
		effect: effect.DealDamage(4, param.Char(choose.All, pred.And(pred.Enemy, pred.Minion))),
	},
}
