package test

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/buff"
	"github.com/kcnm/rocky/engine/effect"
)

var (
	M11 = NewMinionCard(engine.Neutral, 1, 1, 1, buff.None)
	M22 = NewMinionCard(engine.Neutral, 2, 2, 2, buff.None)
	M33 = NewMinionCard(engine.Neutral, 3, 3, 3, buff.None)
	M44 = NewMinionCard(engine.Neutral, 4, 4, 4, buff.None)
	M45 = NewMinionCard(engine.Neutral, 4, 4, 5, buff.None)
	M55 = NewMinionCard(engine.Neutral, 5, 5, 5, buff.None)
	M66 = NewMinionCard(engine.Neutral, 6, 6, 6, buff.None)
	M77 = NewMinionCard(engine.Neutral, 7, 7, 7, buff.None)
	M88 = NewMinionCard(engine.Neutral, 8, 8, 8, buff.None)
	S4  = NewSpellCard(engine.Neutral, 4, effect.None)
	W32 = NewWeaponCard(engine.Neutral, 2, 3, 2)
	W33 = NewWeaponCard(engine.Neutral, 4, 3, 3)
	Pw2 = NewPower(engine.Neutral, 2, effect.None)
)

func NewMinionCard(
	class engine.Class,
	mana int,
	attack int,
	health int,
	buff engine.Buff) engine.MinionCard {
	return &minionCard{class, mana, attack, health, buff}
}

type minionCard struct {
	class  engine.Class
	mana   int
	attack int
	health int
	buff   engine.Buff
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

func (c *minionCard) Buff() engine.Buff {
	return c.buff
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
