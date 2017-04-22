package test

import (
	"github.com/kcnm/rocky/engine"
)

func NewAbility(
	class engine.Class,
	mana int,
	effect engine.Effect) engine.Ability {
	return &ability{class, mana, effect}
}

type ability struct {
	class  engine.Class
	mana   int
	effect engine.Effect
}

func (a *ability) Class() engine.Class {
	return a.class
}

func (a *ability) Mana() int {
	return a.mana
}

func (a *ability) Effect() engine.Effect {
	return a.effect
}
