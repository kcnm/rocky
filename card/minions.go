package card

import (
	"github.com/kcnm/rocky/engine/base"
)

type minionSpec struct {
	name   string
	class  base.Class
	mana   int
	attack int
	health int
}

func (c minion) Class() base.Class {
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

func (c minion) String() string {
	return minions[c].name
}

var minions = map[minion]*minionSpec{
	SilverHandRecruit: &minionSpec{
		name:   "Silver Hand Recruit",
		class:  base.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
	},
}
