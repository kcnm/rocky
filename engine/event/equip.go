package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type equip struct {
	game   base.Game
	player base.Player
	card   base.WeaponCard
}

func Equip(
	game base.Game,
	player base.Player,
	card base.WeaponCard) base.Event {
	return &equip{game, player, card}
}

func (ev *equip) Subject() interface{} {
	return ev.player
}

func (ev *equip) Verb() base.Verb {
	return base.Equip
}

func (ev *equip) Trigger() {
	ev.player.Equip(ev.card)
}
