package engine

import (
	"math/rand"
)

type Game interface {
	Listener

	Events() EventBus
	RNG() *rand.Rand
	Turn() int
	CurrentPlayer() Player
	Opponent(player Player) Player
	AllChars() []Char
	IsOver() (over bool, winner Player)

	Summon(card MinionCard, board Board, position int) Minion
}
