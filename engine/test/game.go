package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
)

type GameStatus struct {
	Current engine.Player
	Over    bool
	Winner  engine.Player
	P1, P2  PlayerStatus
	B1, B2  []MinionStatus
}

func AssertGameStatus(
	t *testing.T,
	game engine.Game,
	status GameStatus) {
	// Checks general game status.
	if status.Over {
		AssertGameOver(t, game, status.Winner)
		return // No need for rest of checks.
	} else {
		AssertGameNotOver(t, game)
	}
	AssertCurrentPlayer(t, game, status.Current)

	// Checks each player's status.
	p1 := game.CurrentPlayer()
	if game.Turn()%2 == 0 {
		p1 = game.Opponent(p1)
	}
	p2 := game.Opponent(p1)
	AssertPlayerStatus(t, p1, status.P1)
	AssertPlayerStatus(t, p2, status.P2)

	// Checks board status.
	for i, m := range status.B1 {
		AssertMinionStatus(t, p1.Board().Get(i), m)
	}
	for i, m := range status.B2 {
		AssertMinionStatus(t, p2.Board().Get(i), m)
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
