package test

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect"
)

var (
	M11 = NewMinionCard(engine.Neutral, 1, 1, 1)
	M45 = NewMinionCard(engine.Neutral, 4, 4, 5)
	S4  = NewSpellCard(engine.Neutral, 4, effect.None)
	W32 = NewWeaponCard(engine.Neutral, 2, 3, 2)
	W33 = NewWeaponCard(engine.Neutral, 4, 3, 3)
	Pw2 = NewPower(engine.Neutral, 2, effect.None)
)

func NewMinionCard(
	class engine.Class,
	mana int,
	attack int,
	health int) engine.MinionCard {
	return &minionCard{class, mana, attack, health}
}

type minionCard struct {
	class  engine.Class
	mana   int
	attack int
	health int
}

func (c *minionCard) Class() engine.Class {
	return c.class
}

func (c *minionCard) Mana() int {
	return c.mana
}

func (c *minionCard) Attack() int {
	return c.attack
}

func (c *minionCard) Health() int {
	return c.health
}

func NewSpellCard(
	class engine.Class,
	mana int,
	effect engine.Effect) engine.SpellCard {
	return &spellCard{class, mana, effect}
}

type spellCard struct {
	class  engine.Class
	mana   int
	effect engine.Effect
}

func (c *spellCard) Class() engine.Class {
	return c.class
}

func (c *spellCard) Mana() int {
	return c.mana
}

func (c *spellCard) Effect() engine.Effect {
	return c.effect
}

func NewWeaponCard(
	class engine.Class,
	mana int,
	attack int,
	durability int) engine.WeaponCard {
	return &weaponCard{class, mana, attack, durability}
}

type weaponCard struct {
	class      engine.Class
	mana       int
	attack     int
	durability int
}

func (c *weaponCard) Class() engine.Class {
	return c.class
}

func (c *weaponCard) Mana() int {
	return c.mana
}

func (c *weaponCard) Attack() int {
	return c.attack
}

func (c *weaponCard) Durability() int {
	return c.durability
}
