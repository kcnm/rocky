package event

import (
	"github.com/kcnm/rocky/engine"
)

// attack is an implementation of Attack Event.
type attack struct {
	attacker engine.Char
	defender engine.Char
}

// Attack returns a new Event, where one character, attacker, attacks another
// one, defender.
func Attack(
	attacker engine.Char,
	defender engine.Char) engine.Event {
	return &attack{attacker, defender}
}

func (ev *attack) Subject() interface{} {
	return ev.attacker
}

func (ev *attack) Verb() engine.Verb {
	return engine.Attack
}

func (ev *attack) Object() interface{} {
	return ev.defender
}

func (ev *attack) Trigger(q engine.EventQueue) {
	ev.attacker.Swing()
	active := Damage(ev.attacker, ev.defender, ev.attacker.Attack())
	passive := Damage(ev.defender, ev.attacker, ev.defender.Attack())
	_, isPlayer := ev.defender.(engine.Player)
	if !isPlayer && ev.defender.Attack() > 0 {
		q.Post(Combined(active, passive), ev)
	} else {
		q.Post(active, ev)
	}
	if player, isPlayer := ev.attacker.(engine.Player); isPlayer {
		w := player.Weapon()
		if w.LoseDurability(1) == 0 {
			q.Post(Destroy(w), ev)
		}
	}
}
