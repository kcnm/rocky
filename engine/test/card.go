package test

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/target"
)

func NewMinionCard(
	class engine.Class,
	mana int,
	attack int,
	health int) engine.Card {
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
	assign target.Assign,
	side target.Side,
	role target.Role,
	effects []engine.Effect) engine.Card {
	return &spellCard{class, mana, assign, side, role, effects}
}

type spellCard struct {
	class   engine.Class
	mana    int
	assign  target.Assign
	side    target.Side
	role    target.Role
	effects []engine.Effect
}

func (c *spellCard) Class() engine.Class {
	return c.class
}

func (c *spellCard) Mana() int {
	return c.mana
}

func (c *spellCard) Assign() target.Assign {
	return c.assign
}

func (c *spellCard) Side() target.Side {
	return c.side
}

func (c *spellCard) Role() target.Role {
	return c.role
}

func (c *spellCard) Effects() []engine.Effect {
	return c.effects
}

func NewWeaponCard(
	class engine.Class,
	mana int,
	attack int,
	durability int) engine.Card {
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
