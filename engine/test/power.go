package test

import (
	"github.com/kcnm/rocky/engine"
)

func NewPower(
	class engine.Class,
	mana int,
	effects ...engine.Effect) engine.Power {
	return &power{class, mana, effects}
}

type power struct {
	class   engine.Class
	mana    int
	effects []engine.Effect
}

func (pw *power) Class() engine.Class {
	return pw.class
}

func (pw *power) Mana() int {
	return pw.mana
}

func (pw *power) Effects() []engine.Effect {
	return pw.effects
}
