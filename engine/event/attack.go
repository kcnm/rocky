package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type attack struct {
	game     base.Game
	attacker base.Character
	defender base.Character
}

func Attack(
	game base.Game,
	attacker base.Character,
	defender base.Character) base.Event {
	return &attack{game, attacker, defender}
}

func (ev *attack) Subject() interface{} {
	return ev.attacker
}

func (ev *attack) Verb() base.Verb {
	return base.Attack
}

func (ev *attack) Trigger() {
	ev.attacker.LoseStamina()
	active := Damage(ev.game, ev.defender, ev.attacker, ev.attacker.Attack())
	passive := Damage(ev.game, ev.attacker, ev.defender, ev.defender.Attack())
	if ev.defender.Attack() > 0 {
		ev.game.Events().Post(base.Combined(active, passive), ev)
	} else {
		ev.game.Events().Post(active, ev)
	}
}
