package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type weapon struct {
	card       engine.WeaponCard
	attack     int
	durability int
}

func newWeapon(card engine.WeaponCard) engine.Weapon {
	return &weapon{card, card.Attack(), card.Durability()}
}

func (w *weapon) Card() engine.WeaponCard {
	return w.card
}

func (w *weapon) Attack() int {
	return w.attack
}

func (w *weapon) Durability() int {
	return w.durability
}

func (w *weapon) LoseDurability() {
	w.durability--
}

func (w *weapon) String() string {
	return fmt.Sprintf("%v %d/%d", w.card, w.attack, w.durability)
}
