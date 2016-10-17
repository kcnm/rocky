package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type destroyWeapon struct {
	game   base.Game
	player base.Player
}

func DestroyWeapon(game base.Game, player base.Player) base.Event {
	return &destroyWeapon{game, player}
}

func (ev *destroyWeapon) Subject() interface{} {
	return ev.player
}

func (ev *destroyWeapon) Verb() base.Verb {
	return base.DestroyWeapon
}

func (ev *destroyWeapon) Trigger() {
	ev.player.DestroyWeapon()
}
