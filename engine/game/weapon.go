package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type weapon struct {
	*entity
	card       engine.WeaponCard
	attack     int
	durability int
}

func newWeapon(
	id engine.EntityID,
	card engine.WeaponCard) engine.Weapon {
	return &weapon{
		newEntity(id).(*entity),
		card,
		card.Attack(),
		card.Durability(),
	}
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

func (w *weapon) LoseDurability(d int) int {
	if w.durability < d {
		panic(fmt.Errorf("cannot lose durability: %d/%d", d, w.durability))
	}
	w.durability -= d
	return w.durability
}

func (w *weapon) String() string {
	return fmt.Sprintf("%v %d/%d", w.card, w.attack, w.durability)
}
