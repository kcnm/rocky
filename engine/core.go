// Package engine defines the game engine that executes rules, together with
// basic interfaces that abstract in-game entities.
package engine

import (
	"math/rand"
)

// Game is a single game instance of two players.
type Game interface {
	Entity
	EventQueue

	// RNG returns the random number generator used to create randomized events.
	RNG() *rand.Rand
	// Turn returns the current number of turns since game start, from 1.
	Turn() int
	// CurrentPlayer returns the current player in play.
	CurrentPlayer() Player
	// Opponent returns the opponent of given player.
	Opponent(player Player) Player
	// AllChars returns a slice of all characters currently in game.
	AllChars() []Char
	// IsOver returns whether the game is over, and the winner if so. The winner
	// is nil if the game is not over, or it is a dual.
	IsOver() (over bool, winner Player)

	// Start kicks off the game and automatically starts the first turn.
	Start()
	// Summon results in a minion of given card summoned for the player at its
	// given board position. Returns the summoned minion instance.
	Summon(player Player, card MinionCard, position int) Minion
	// Equip results in a weapon of given card equiped for the player. Returns the
	// equiped weapon instance.
	Equip(player Player, card WeaponCard) Weapon
}

// EventQueue is the core component of the game engine. Any in-game status
// changes are modeled as an Event, which may trigger a chain of events as its
// consequences. This queue is responsible for managing events in a game, and
// resolve the causality chains when actions are taken by player.
//
// There are participants modeled as Entity, which can subscribe to EventQueue
// and react to events at its resolving process, possibly chaining more events.
type EventQueue interface {
	// Join allows entity e to participate in the queue and react to its events.
	Join(e Entity)
	// Exit removes entity e from queue's participants, returns true if so, or
	// false if e is not participating.
	Exit(e Entity) bool
	// Fire kicks off event ev and resolves all of its consequences.
	Fire(ev Event)
	// Post enqueues event ev with its cause, to be resolved in current run. This
	// is the normal case when chaining events for causality.
	Post(ev Event, cause Event)
	// Cache enqueues event ev with its cause, to be resolved in future run. This
	// is used for modeling sequential effects where multiple runs of resolving
	// are required.
	Cache(ev Event, cause Event)
	// Drain dumps all the contents in the queue, usually happens at game over.
	Drain()
}

// Event is a incident modelling in-game status change as a natural language
// phrase. In example "minion1 attacks player2", where "minion1" is the subject,
// "attack" is the verb, and "player2" is the object.
type Event interface {
	// Subject returns the subject of this event.
	Subject() interface{}
	// Verb returns the verb/type of this event.
	Verb() Verb
	// Object returns the object of this event, or nil if not any.
	Object() interface{}

	// Trigger makes the event actually happen, possibly enqueueing more events as
	// its direct consequence. It is required that new events enqueued through
	// this process are validated before.
	Trigger(q EventQueue)
}

// Verb is the action, or type of event.
type Verb string

const (
	StartGame Verb = "StartGame"
	StartTurn Verb = "StartTurn"
	EndTurn   Verb = "EndTurn"
	Draw      Verb = "Draw"
	TakeCard  Verb = "TakeCard"
	PlayCard  Verb = "PlayCard"
	Summon    Verb = "Summon"
	Cast      Verb = "Cast"
	Equip     Verb = "Equip"
	HeroPower Verb = "HeroPower"
	Attack    Verb = "Attack"
	Damage    Verb = "Damage"
	Dying     Verb = "Dying"
	Destroy   Verb = "Destroy"
	Impact    Verb = "Impact"
	GameOver  Verb = "GameOver"
	Combined  Verb = "Combined"
	Sequence  Verb = "Sequence"
)
