package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type char struct {
	*entity
	attack    int
	health    int
	maxHealth int
	swings    int
	maxSwings int
}

func newChar(
	id engine.EntityID,
	attack int,
	health int,
	maxHealth int) engine.Char {
	return &char{
		newEntity(id).(*entity),
		attack,
		health,
		maxHealth,
		0, // swings
		0, // maxSwings
	}
}

func (ch *char) Attack() int {
	return ch.attack
}

func (ch *char) Health() int {
	return ch.health
}

func (ch *char) MaxHealth() int {
	return ch.maxHealth
}

func (ch *char) Swings() int {
	return ch.swings
}

func (ch *char) Active() bool {
	return ch.Attack() > 0 && ch.swings < ch.maxSwings
}

func (ch *char) Refresh() {
	ch.swings = 0
	ch.maxSwings = 1
}

func (ch *char) TakeDamage(damage int) (actual int, fatal bool) {
	if damage <= 0 {
		panic(fmt.Errorf("non-positive damage %d", damage))
	}
	ch.health -= damage
	return damage, ch.health <= 0
}

func (ch *char) Swing() int {
	if ch.swings >= ch.maxSwings {
		panic(fmt.Errorf("cannot swing: %d/%d", ch.swings, ch.maxSwings))
	}
	ch.swings++
	return ch.swings
}
