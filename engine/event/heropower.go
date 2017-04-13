package event

import (
	"github.com/kcnm/rocky/engine"
)

type heroPower struct {
	game   engine.Game
	player engine.Player
	target engine.Char
}

func HeroPower(
	game engine.Game,
	player engine.Player,
	target engine.Char) engine.Event {
	return &heroPower{game, player, target}
}

func (ev *heroPower) Subject() interface{} {
	return ev.player
}

func (ev *heroPower) Verb() engine.Verb {
	return engine.HeroPower
}

func (ev *heroPower) Object() interface{} {
	return ev.target
}

func (ev *heroPower) Trigger(q engine.EventQueue) {
	q.Post(Impact(ev.game, ev.player, ev.target, ev.player.HeroPower()), ev)
}
