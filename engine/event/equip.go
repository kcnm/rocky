package event

import (
	"github.com/kcnm/rocky/engine"
)

type equip struct {
	player engine.Player
	card   engine.WeaponCard
}

func Equip(
	player engine.Player,
	card engine.WeaponCard) engine.Event {
	return &equip{player, card}
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
	ev.player.Equip(ev.card)
}
