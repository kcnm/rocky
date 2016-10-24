package test

import (
	"testing"

	"github.com/kcnm/rocky/card"
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
)

func TestBasicGame(t *testing.T) {
	player1 := game.NewPlayer(
		30, // health
		0,  // armor
		0,  // crystal
		game.NewDeck(
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
		),
	)
	player2 := game.NewPlayer(
		30, // health
		0,  // armor
		0,  // crystal
		game.NewDeck(
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
		),
	)
	g := game.New(player1, player2, nil /* rng */)

	for _, turn := range []struct {
		current engine.Player
		over    bool
		winner  engine.Player
		players []playerStatus
	}{
		{
			// turn 1
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 0, 0, 0, 4, 0},
			},
		},
		{
			// turn 2
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 3
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 4
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 5
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 6
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 7
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 8
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 9
			player1, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 10
			player2, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 11
			player1, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 12
			player2, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 13
			player1, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 14
			player2, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 15
			player1, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 16
			player2, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 17
			player1, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 18
			player2, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 19
			player1, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 20
			player2, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 21
			player1, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 22
			player2, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 23
			player1, true, player2,
			[]playerStatus{
				{-6, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
	} {
		t.Logf("Turn %d", g.Turn())
		assertGameStatus(t, g, turn.current, turn.over, turn.winner)
		assertPlayerStatus(t, player1, turn.players[0])
		assertPlayerStatus(t, player2, turn.players[1])
		p := g.CurrentPlayer()
		if over, _ := g.IsOver(); !over {
			if ok, _ := action.CanEndTurn(g, g.Opponent(p)); ok {
				t.Errorf("Opponent can end turn, expected not")
			}
			action.EndTurn(g, p)
		} else {
			if ok, _ := action.CanEndTurn(g, p); ok {
				t.Errorf("Player can end turn after game over, expected not")
			}
			if ok, _ := action.CanEndTurn(g, g.Opponent(p)); ok {
				t.Errorf("Opponent can end turn after game over, expected not")
			}
		}
	}
}

func TestPlayMinion(t *testing.T) {
	player1 := game.NewPlayer(30, 0, 10, game.NewDeck(),
		card.ChillwindYeti, card.SilverHandRecruit, card.ChillwindYeti)
	player2 := game.NewPlayer(30, 0, 10, game.NewDeck())
	g := game.New(player1, player2, nil /* rng */)
	// turn 1
	assertPlayerStatus(t, player1, playerStatus{
		29, 0, 0, false, 10, 10, 3, 0, 0})
	action.PlayCard(g, player1, 0, 0, nil)
	assertPlayerStatus(t, player1, playerStatus{
		29, 0, 0, false, 6, 10, 2, 0, 1})
	assertMinionStatus(t, player1.Board().Get(0), minionStatus{
		4, 5, false, card.ChillwindYeti})
	action.EndTurn(g, player1)
	// turn 2
	action.EndTurn(g, player2)
	// turn 3
	assertPlayerStatus(t, player1, playerStatus{
		27, 0, 0, false, 10, 10, 2, 0, 1})
	assertMinionStatus(t, player1.Board().Get(0), minionStatus{
		4, 5, true, card.ChillwindYeti})
	action.PlayCard(g, player1, 1, 0, nil)
	assertPlayerStatus(t, player1, playerStatus{
		27, 0, 0, false, 6, 10, 1, 0, 2})
	assertMinionStatus(t, player1.Board().Get(0), minionStatus{
		4, 5, false, card.ChillwindYeti})
	assertMinionStatus(t, player1.Board().Get(1), minionStatus{
		4, 5, true, card.ChillwindYeti})
	action.PlayCard(g, player1, 0, 2, nil)
	assertPlayerStatus(t, player1, playerStatus{
		27, 0, 0, false, 5, 10, 0, 0, 3})
	assertMinionStatus(t, player1.Board().Get(0), minionStatus{
		4, 5, false, card.ChillwindYeti})
	assertMinionStatus(t, player1.Board().Get(1), minionStatus{
		4, 5, true, card.ChillwindYeti})
	assertMinionStatus(t, player1.Board().Get(2), minionStatus{
		1, 1, false, card.SilverHandRecruit})
}
