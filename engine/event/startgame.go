package event

import (
	"github.com/kcnm/rocky/engine"
)

// startGame is an implementation of StartGame Event.
type startGame struct {
	game engine.Game
}

// StartGame returns a new Event, where the game starts.
func StartGame(game engine.Game) engine.Event {
	return &startGame{game}
}

func (ev *startGame) Subject() interface{} {
	return ev.game
}

func (ev *startGame) Verb() engine.Verb {
	return engine.StartGame
}

func (ev *startGame) Object() interface{} {
	return nil
}

func (ev *startGame) Trigger(q engine.EventQueue) {
	q.Post(StartTurn(ev.game, ev.game.CurrentPlayer()), ev)
}
