package event

import (
	"github.com/kcnm/rocky/engine"
)

type attack struct {
	game     engine.Game
	attacker engine.Char
	defender engine.Char
}

func Attack(
	game engine.Game,
	attacker engine.Char,
	defender engine.Char) engine.Event {
	return &attack{game, attacker, defender}
}

func (ev *attack) Subject() interface{} {
	return ev.attacker
}

func (ev *attack) Verb() engine.Verb {
	return engine.Attack
}

func (ev *attack) Trigger() {
	ev.attacker.LoseStamina()
	active := Damage(ev.game, ev.defender, ev.attacker, ev.attacker.Attack())
	passive := Damage(ev.game, ev.attacker, ev.defender, ev.defender.Attack())
	_, isPlayer := ev.defender.(engine.Player)
	if !isPlayer && ev.defender.Attack() > 0 {
		ev.game.Events().Post(engine.Combined(active, passive), ev)
	} else {
		ev.game.Events().Post(active, ev)
	}
	if player, isPlayer := ev.attacker.(engine.Player); isPlayer {
		player.Weapon().LoseDurability()
		if player.Weapon().Durability() == 0 {
			ev.game.Events().Post(DestroyWeapon(ev.game, player), ev)
		}
	}
}
