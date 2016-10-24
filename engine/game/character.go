package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type character struct {
	id      engine.CharacterID
	attack  int
	health  int
	stamina int
}

func newCharacter(
	id engine.CharacterID,
	attack int,
	health int,
	stamina int) engine.Character {
	return &character{id, attack, health, stamina}
}

func (char *character) ID() engine.CharacterID {
	return char.id
}

func (char *character) Attack() int {
	return char.attack
}

func (char *character) Health() int {
	return char.health
}

func (char *character) Stamina() int {
	return char.stamina
}

func (char *character) Active() bool {
	return char.Attack() > 0 && char.Stamina() > 0
}

func (char *character) Assign(id engine.CharacterID) {
	char.id = id
}

func (char *character) Refresh() {
	char.stamina = 1
}

func (char *character) TakeDamage(damage int) (actual int, fatal bool) {
	if damage <= 0 {
		panic(fmt.Errorf("non-positive damage %d", damage))
	}
	char.health -= damage
	return damage, char.health <= 0
}

func (char *character) LoseStamina() {
	if char.stamina <= 0 {
		panic(fmt.Errorf("non-positive stamina %d", char.stamina))
	}
	char.stamina--
}
