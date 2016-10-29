package power

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect"
)

func (pw standard) Class() engine.Class {
	return standards[pw].class
}

func (pw standard) Mana() int {
	return standards[pw].mana
}

func (pw standard) Effect() engine.Effect {
	return standards[pw].effect
}

func (pw standard) String() string {
	return standards[pw].name
}

var standards = map[standard]*powerSpec{
	Fireblast: &powerSpec{
		name:   "Fireblast",
		class:  engine.Mage,
		mana:   2,
		effect: effect.DealDamage(1),
	},
}
