package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
)

func AssertGameStatus(
	t *testing.T,
	game engine.Game,
	current engine.Player,
	over bool,
	winner engine.Player) {
	AssertCurrentPlayer(t, game, current)
	if over {
		AssertGameOver(t, game, winner)
	} else {
		AssertGameNotOver(t, game)
	}
}

func AssertCurrentPlayer(
	t *testing.T,
	game engine.Game,
	player engine.Player) {
	if game.CurrentPlayer() != player {
		t.Errorf("Current player is player%v, expected player%v",
			game.CurrentPlayer().ID(), player.ID())
	}
}

func AssertGameNotOver(
	t *testing.T,
	game engine.Game) {
	if over, _ := game.IsOver(); over {
		t.Errorf("Game is over, expected not")
	}
}

func AssertGameOver(
	t *testing.T,
	game engine.Game,
	winner engine.Player) {
	o, w := game.IsOver()
	if !o {
		t.Errorf("Game is not over, expected over")
	} else if w != winner {
		if w == nil {
			t.Errorf("Dual, expected winner player%v", winner.ID())
		}
		if winner == nil {
			t.Errorf("Winner is player%v, expected dual", w.ID())
		}
	}
}
