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

func (s spellSpec) Class() engine.Class {
	return s.class
}

func (s spellSpec) Mana() int {
	return s.mana
}

func (s spellSpec) Effect() engine.Effect {
	return s.effect
}

var spells = map[string]*spellSpec{
	arcaneMissiles: &spellSpec{
		class: engine.Mage,
		mana:  1,
		effect: effect.SplitDamage(
			param.Const(3),
			param.Char(choose.Random, pred.Enemy),
		),
	},
	fireball: &spellSpec{
		class: engine.Mage,
		mana:  4,
		effect: effect.DealDamage(
			param.Const(6),
			param.Char(choose.Manual, pred.Char),
		),
	},
	flamestrike: &spellSpec{
		class: engine.Mage,
		mana:  7,
		effect: effect.DealDamage(
			param.Const(4),
			param.Char(choose.All, pred.And(pred.Enemy, pred.Minion)),
		),
	},
	lightningStorm: &spellSpec{
		class: engine.Shaman,
		mana:  3,
		// TODO: Overload
		effect: effect.DealDamage(
			param.Span(2, 3),
			param.Char(choose.All, pred.And(pred.Enemy, pred.Minion)),
		),
	},
}
