package card

import (
	"github.com/kcnm/rocky/engine"
)

type minionSpec struct {
	class  engine.Class
	mana   int
	attack int
	health int
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
	SilverHandRecruit: &minionSpec{
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
	},
	ChillwindYeti: &minionSpec{
		class:  engine.Neutral,
		mana:   4,
		attack: 4,
		health: 5,
	},
}
