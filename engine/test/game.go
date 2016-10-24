package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
)

func assertGameStatus(
	t *testing.T,
	game engine.Game,
	current engine.Player,
	over bool,
	winner engine.Player) {
	assertCurrentPlayer(t, game, current)
	if over {
		assertGameOver(t, game, winner)
	} else {
		assertGameNotOver(t, game)
	}
}

func assertCurrentPlayer(
	t *testing.T,
	game engine.Game,
	player engine.Player) {
	if game.CurrentPlayer() != player {
		t.Errorf("Current player is player%d, expected player%d",
			game.CurrentPlayer().ID(), player.ID())
	}
}

func assertGameNotOver(
	t *testing.T,
	game engine.Game) {
	if over, _ := game.IsOver(); over {
		t.Errorf("Game is over, expected not")
	}
}

func assertGameOver(
	t *testing.T,
	game engine.Game,
	winner engine.Player) {
	o, w := game.IsOver()
	if !o {
		t.Errorf("Game is not over, expected over")
	} else if w != winner {
		if w == nil {
			t.Errorf("Dual, expected winner player%d", winner.ID())
		}
		if winner == nil {
			t.Errorf("Winner is player%d, expected dual", w.ID())
		}
	}
}
