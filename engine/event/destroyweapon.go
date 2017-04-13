package event

import (
	"github.com/kcnm/rocky/engine"
)

type destroyWeapon struct {
	player engine.Player
}

func DestroyWeapon(player engine.Player) engine.Event {
	return &destroyWeapon{player}
}

func (ev *destroyWeapon) Subject() interface{} {
	return ev.player
}

func (ev *destroyWeapon) Verb() engine.Verb {
	return engine.DestroyWeapon
}

func (ev *destroyWeapon) Object() interface{} {
	return nil
}

func (ev *destroyWeapon) Trigger(q engine.EventQueue) {
	ev.player.DestroyWeapon()
}
