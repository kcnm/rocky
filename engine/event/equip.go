package event

import (
	"github.com/kcnm/rocky/engine"
)

// equip is an implementation of Equip Event.
type equip struct {
	game   engine.Game
	player engine.Player
	card   engine.WeaponCard
}

// Equip returns a new Event, where a player equips a weapon.
func Equip(
	game engine.Game,
	player engine.Player,
	card engine.WeaponCard) engine.Event {
	return &equip{game, player, card}
}

func (ev *equip) Subject() interface{} {
	return ev.player
}

func (ev *equip) Verb() engine.Verb {
	return engine.Equip
}

func (ev *equip) Object() interface{} {
	return ev.card
}

func (ev *equip) Trigger(q engine.EventQueue) {
	ev.game.Equip(ev.player, ev.card)
}
