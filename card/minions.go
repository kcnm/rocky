package card

import (
	"github.com/kcnm/rocky/engine"
)

type minionSpec struct {
	name   string
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

func (c minion) String() string {
	return minions[c].name
}

var minions = map[minion]*minionSpec{
	SilverHandRecruit: &minionSpec{
		name:   "Silver Hand Recruit",
		class:  engine.Neutral,
		mana:   1,
		attack: 1,
		health: 1,
	},
	ChillwindYeti: &minionSpec{
		name:   "Chillwind Yeti",
		class:  engine.Neutral,
		mana:   4,
		attack: 4,
		health: 5,
	},
}
