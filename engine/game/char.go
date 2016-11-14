package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type char struct {
	id        engine.CharID
	attack    int
	health    int
	maxHealth int
	stamina   int
	handlers  []engine.Handler
}

func newChar(
	id engine.CharID,
	attack int,
	health int,
	maxHealth int,
	stamina int) engine.Char {
	return &char{
		id,
		attack,
		health,
		maxHealth,
		stamina,
		make([]engine.Handler, 0),
	}
}

func (ch *char) Handle(ev engine.Event) {
	for _, h := range ch.handlers {
		h(ev)
	}
}

func (ch *char) ID() engine.CharID {
	return ch.id
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

func (ch *char) AddHandler(handler engine.Handler) {
	ch.handlers = append(ch.handlers, handler)
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
