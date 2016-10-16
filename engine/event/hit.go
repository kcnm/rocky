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
	ev.game.Events().Post(
		Damage(ev.game, ev.defender, ev.attacker, ev.attacker.Attack()))
	if ev.defender.Attack() > 0 {
		ev.game.Events().Post(
			Damage(ev.game, ev.attacker, ev.defender, ev.defender.Attack()))
	}
}
