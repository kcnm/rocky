package test

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/buff"
	"github.com/kcnm/rocky/engine/effect"
)

var (
	M01 = NewMinionCard(engine.Neutral, 1, 0, 1, effect.None, buff.None)
	M11 = NewMinionCard(engine.Neutral, 1, 1, 1, effect.None, buff.None)
	M22 = NewMinionCard(engine.Neutral, 2, 2, 2, effect.None, buff.None)
	M33 = NewMinionCard(engine.Neutral, 3, 3, 3, effect.None, buff.None)
	M44 = NewMinionCard(engine.Neutral, 4, 4, 4, effect.None, buff.None)
	M45 = NewMinionCard(engine.Neutral, 4, 4, 5, effect.None, buff.None)
	M55 = NewMinionCard(engine.Neutral, 5, 5, 5, effect.None, buff.None)
	M66 = NewMinionCard(engine.Neutral, 6, 6, 6, effect.None, buff.None)
	M77 = NewMinionCard(engine.Neutral, 7, 7, 7, effect.None, buff.None)
	M88 = NewMinionCard(engine.Neutral, 8, 8, 8, effect.None, buff.None)
	S4  = NewSpellCard(engine.Neutral, 4, effect.None)
	W32 = NewWeaponCard(engine.Neutral, 2, 3, 2, effect.None)
	W33 = NewWeaponCard(engine.Neutral, 4, 3, 3, effect.None)
	A2  = NewAbility(engine.Neutral, 2, effect.None)
)

func NewMinionCard(
	class engine.Class,
	mana int,
	attack int,
	health int,
	battlecry engine.Effect,
	buff engine.Buff) engine.MinionCard {
	return &minionCard{class, mana, attack, health, battlecry, buff}
}

type minionCard struct {
	class     engine.Class
	mana      int
	attack    int
	health    int
	battlecry engine.Effect
	buff      engine.Buff
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

func (c *minionCard) Battlecry() engine.Effect {
	return c.battlecry
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
	durability int,
	battlecry engine.Effect) engine.WeaponCard {
	return &weaponCard{class, mana, attack, durability, battlecry}
}

type weaponCard struct {
	class      engine.Class
	mana       int
	attack     int
	durability int
	battlecry  engine.Effect
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

func (c *weaponCard) Battlecry() engine.Effect {
	return c.battlecry
}
