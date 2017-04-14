package event

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type startGame struct {
	game engine.Game
}

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
	if t := ev.game.Turn(); t != 0 {
		panic(fmt.Errorf("non-zero start turn %d", t))
	}
	q.Post(StartTurn(ev.game, ev.game.CurrentPlayer()), ev)
}
