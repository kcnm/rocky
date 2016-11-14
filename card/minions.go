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
	class  engine.Class
	mana   int
	attack int
	health int
	buff   engine.Buff
}

func (c minion) Class() engine.Class {
	return minions[c].class
}

func (c minion) Mana() int {
	return minions[c].mana
}

func (c minion) Attack() int {
	return minions[c].attack
}

func (c minion) Health() int {
	return minions[c].health
}

var minions = map[minion]*minionSpec{
	ChillwindYeti: &minionSpec{
		class:  engine.Neutral,
		mana:   4,
		attack: 4,
		health: 5,
	},
	LeperGnome: &minionSpec{
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
	SilverHandRecruit: &minionSpec{
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
	},
}
