// Package engine defines the game engine that executes rules, together with
// basic interfaces that abstract in-game entities.
package engine

import (
	"math/rand"
)

type Verb string

const (
	StartTurn     Verb = "StartTurn"
	EndTurn       Verb = "EndTurn"
	Draw          Verb = "Draw"
	TakeCard      Verb = "TakeCard"
	PlayCard      Verb = "PlayCard"
	Summon        Verb = "Summon"
	Cast          Verb = "Cast"
	Equip         Verb = "Equip"
	HeroPower     Verb = "HeroPower"
	Attack        Verb = "Attack"
	Damage        Verb = "Damage"
	Dying         Verb = "Dying"
	Destroy       Verb = "Destroy"
	DestroyWeapon Verb = "DestroyWeapon"
	Impact        Verb = "Impact"
	GameOver      Verb = "GameOver"
	Combined      Verb = "Combined"
	Sequence      Verb = "Sequence"
)

type Event interface {
	Subject() interface{}
	Verb() Verb

	Trigger(EventQueue)
}

type ListenerID int

type Listener interface {
	Handle(ev Event)
}

type Handler func(Event)

type EventQueue interface {
	AddListener(listener Listener) ListenerID
	RemoveListener(id ListenerID) bool
	Fire(ev Event)
	Post(ev Event, cause Event)
	Cache(ev Event, cause Event)
	Drain()
}

type Game interface {
	EventQueue
	Listener

	RNG() *rand.Rand
	Turn() int
	CurrentPlayer() Player
	Opponent(player Player) Player
	AllChars() []Char
	IsOver() (over bool, winner Player)

	Start()
	Summon(card MinionCard, player Player, position int) Minion
}
