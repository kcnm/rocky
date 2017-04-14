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
	stamina   int
}

func newChar(
	id engine.EntityID,
	attack int,
	health int,
	maxHealth int,
	stamina int) engine.Char {
	return &char{
		newEntity(id).(*entity),
		attack,
		health,
		maxHealth,
		stamina,
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

func (ch *char) Stamina() int {
	return ch.stamina
}

func (ch *char) Active() bool {
	return ch.Attack() > 0 && ch.Stamina() > 0
}

func (ch *char) Refresh() {
	ch.stamina = 1
}

func (ch *char) TakeDamage(damage int) (actual int, fatal bool) {
	if damage <= 0 {
		panic(fmt.Errorf("non-positive damage %d", damage))
	}
	ch.health -= damage
	return damage, ch.health <= 0
}

func (ch *char) LoseStamina() {
	if ch.stamina <= 0 {
		panic(fmt.Errorf("non-positive stamina %d", ch.stamina))
	}
	ch.stamina--
}
