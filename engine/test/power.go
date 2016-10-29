package test

import (
	"github.com/kcnm/rocky/engine"
)

func NewPower(
	class engine.Class,
	mana int,
	effect engine.Effect) engine.Power {
	return &power{class, mana, effect}
}

type power struct {
	class  engine.Class
	mana   int
	effect engine.Effect
}

func (pw *power) Class() engine.Class {
	return pw.class
}

func (pw *power) Mana() int {
	return pw.mana
}

func (pw *power) Effect() engine.Effect {
	return pw.effect
}
