package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type hit struct {
	game     base.Game
	attacker base.Character
	defender base.Character
}

func Hit(
	game base.Game,
	attacker base.Character,
	defender base.Character) base.Event {
	return &hit{game, attacker, defender}
}

func (ev *hit) Subject() interface{} {
	return ev.attacker
}

func (ev *hit) Verb() base.Verb {
	return base.Hit
}

func (ev *hit) Trigger() {
	ev.attacker.LoseStamina()
	active := Damage(ev.game, ev.defender, ev.attacker, ev.attacker.Attack())
	passive := Damage(ev.game, ev.attacker, ev.defender, ev.defender.Attack())
	if ev.defender.Attack() > 0 {
		ev.game.Events().Post(base.Combined(active, passive), ev)
	} else {
		ev.game.Events().Post(active, ev)
	}
}
