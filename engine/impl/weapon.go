package impl

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
)

type weapon struct {
	card       base.WeaponCard
	attack     int
	durability int
}

func newWeapon(card base.WeaponCard) base.Weapon {
	return &weapon{card, card.Attack(), card.Durability()}
}

func (w *weapon) Card() base.WeaponCard {
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
