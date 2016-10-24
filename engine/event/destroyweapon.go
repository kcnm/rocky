package event

import (
	"github.com/kcnm/rocky/engine"
)

type destroyWeapon struct {
	game   engine.Game
	player engine.Player
}

func DestroyWeapon(game engine.Game, player engine.Player) engine.Event {
	return &destroyWeapon{game, player}
}

func (ev *destroyWeapon) Subject() interface{} {
	return ev.player
}

func (ev *destroyWeapon) Verb() engine.Verb {
	return engine.DestroyWeapon
}

func (ev *destroyWeapon) Trigger() {
	ev.player.DestroyWeapon()
}
