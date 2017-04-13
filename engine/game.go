// Package engine defines the game engine that executes rules, together with
// basic interfaces that abstract in-game entities.
package engine

import (
	"math/rand"
)

type Game interface {
	EventBus
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
