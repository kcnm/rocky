package card

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/buff"
	"github.com/kcnm/rocky/engine/effect"
	"github.com/kcnm/rocky/engine/effect/choose"
	"github.com/kcnm/rocky/engine/effect/param"
	"github.com/kcnm/rocky/engine/effect/pred"
)

type minionSpec struct {
	class     engine.Class
	mana      int
	attack    int
	health    int
	battlecry engine.Effect
	buff      engine.Buff
}

func (s minionSpec) Class() engine.Class {
	return s.class
}

func (s minionSpec) Mana() int {
	return s.mana
}

func (s minionSpec) Attack() int {
	return s.attack
}

func (s minionSpec) Health() int {
	return s.health
}

func (s minionSpec) Battlecry() engine.Effect {
	e := s.battlecry
	if e == nil {
		return effect.None
	} else {
		return e
	}
}

func (s minionSpec) Buff() engine.Buff {
	b := s.buff
	if b == nil {
		return buff.None
	} else {
		return b
	}
}

var minions = map[string]*minionSpec{
	chillwindYeti: &minionSpec{
		class:  engine.Neutral,
		mana:   4,
		attack: 4,
		health: 5,
	},
	elvenArcher: &minionSpec{
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
		battlecry: effect.DealDamage(
			param.Const(1),
			param.Char(choose.Manual, pred.Char),
		),
	},
	leperGnome: &minionSpec{
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
		buff: buff.Deathrattle(
			effect.DealDamage(
				param.Const(2),
				param.Char(choose.All, pred.And(pred.Enemy, pred.Hero)),
			),
		),
	},
	silverHandRecruit: &minionSpec{
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
	},
}
