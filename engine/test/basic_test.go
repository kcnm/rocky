package test

import (
	"testing"

	"github.com/kcnm/rocky/card"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/impl"
)

func TestBasicGame(t *testing.T) {
	player1 := impl.NewPlayer(
		30, // health
		0,  // armor
		0,  // crystal
		impl.NewDeck(
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
		),
		nil, // board
	)
	player2 := impl.NewPlayer(
		30, // health
		0,  // armor
		0,  // crystal
		impl.NewDeck(
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
			card.SilverHandRecruit,
		),
		nil, // board
	)
	game := impl.NewGame(player1, player2, nil /* rng */)

	for _, turn := range []struct {
		current base.Player
		over    bool
		winner  base.Player
		players []playerStatus
	}{
		{
			// turn 1
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 1, 1, 1, 3},
				{30, 0, 0, 0, false, 0, 0, 0, 4},
			},
		},
		{
			// turn 2
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 1, 1, 1, 3},
				{30, 0, 0, 1, false, 1, 1, 1, 3},
			},
		},
		{
			// turn 3
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 2, 2, 2, 2},
				{30, 0, 0, 1, false, 1, 1, 1, 3},
			},
		},
		{
			// turn 4
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 2, 2, 2, 2},
				{30, 0, 0, 1, false, 2, 2, 2, 2},
			},
		},
		{
			// turn 5
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 3, 3, 3, 1},
				{30, 0, 0, 1, false, 2, 2, 2, 2},
			},
		},
		{
			// turn 6
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 3, 3, 3, 1},
				{30, 0, 0, 1, false, 3, 3, 3, 1},
			},
		},
		{
			// turn 7
			player1, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 4, 4, 4, 0},
				{30, 0, 0, 1, false, 3, 3, 3, 1},
			},
		},
		{
			// turn 8
			player2, false, nil,
			[]playerStatus{
				{30, 0, 0, 1, false, 4, 4, 4, 0},
				{30, 0, 0, 1, false, 4, 4, 4, 0},
			},
		},
		{
			// turn 9
			player1, false, nil,
			[]playerStatus{
				{29, 0, 0, 1, false, 5, 5, 4, 0},
				{30, 0, 0, 1, false, 4, 4, 4, 0},
			},
		},
		{
			// turn 10
			player2, false, nil,
			[]playerStatus{
				{29, 0, 0, 1, false, 5, 5, 4, 0},
				{29, 0, 0, 1, false, 5, 5, 4, 0},
			},
		},
		{
			// turn 11
			player1, false, nil,
			[]playerStatus{
				{27, 0, 0, 1, false, 6, 6, 4, 0},
				{29, 0, 0, 1, false, 5, 5, 4, 0},
			},
		},
		{
			// turn 12
			player2, false, nil,
			[]playerStatus{
				{27, 0, 0, 1, false, 6, 6, 4, 0},
				{27, 0, 0, 1, false, 6, 6, 4, 0},
			},
		},
		{
			// turn 13
			player1, false, nil,
			[]playerStatus{
				{24, 0, 0, 1, false, 7, 7, 4, 0},
				{27, 0, 0, 1, false, 6, 6, 4, 0},
			},
		},
		{
			// turn 14
			player2, false, nil,
			[]playerStatus{
				{24, 0, 0, 1, false, 7, 7, 4, 0},
				{24, 0, 0, 1, false, 7, 7, 4, 0},
			},
		},
		{
			// turn 15
			player1, false, nil,
			[]playerStatus{
				{20, 0, 0, 1, false, 8, 8, 4, 0},
				{24, 0, 0, 1, false, 7, 7, 4, 0},
			},
		},
		{
			// turn 16
			player2, false, nil,
			[]playerStatus{
				{20, 0, 0, 1, false, 8, 8, 4, 0},
				{20, 0, 0, 1, false, 8, 8, 4, 0},
			},
		},
		{
			// turn 17
			player1, false, nil,
			[]playerStatus{
				{15, 0, 0, 1, false, 9, 9, 4, 0},
				{20, 0, 0, 1, false, 8, 8, 4, 0},
			},
		},
		{
			// turn 18
			player2, false, nil,
			[]playerStatus{
				{15, 0, 0, 1, false, 9, 9, 4, 0},
				{15, 0, 0, 1, false, 9, 9, 4, 0},
			},
		},
		{
			// turn 19
			player1, false, nil,
			[]playerStatus{
				{9, 0, 0, 1, false, 10, 10, 4, 0},
				{15, 0, 0, 1, false, 9, 9, 4, 0},
			},
		},
		{
			// turn 20
			player2, false, nil,
			[]playerStatus{
				{9, 0, 0, 1, false, 10, 10, 4, 0},
				{9, 0, 0, 1, false, 10, 10, 4, 0},
			},
		},
		{
			// turn 21
			player1, false, nil,
			[]playerStatus{
				{2, 0, 0, 1, false, 10, 10, 4, 0},
				{9, 0, 0, 1, false, 10, 10, 4, 0},
			},
		},
		{
			// turn 22
			player2, false, nil,
			[]playerStatus{
				{2, 0, 0, 1, false, 10, 10, 4, 0},
				{2, 0, 0, 1, false, 10, 10, 4, 0},
			},
		},
		{
			// turn 23
			player1, true, player2,
			[]playerStatus{
				{-6, 0, 0, 1, false, 10, 10, 4, 0},
				{2, 0, 0, 1, false, 10, 10, 4, 0},
			},
		},
	} {
		t.Logf("Turn %d", game.Turn())
		assertGameStatus(t, game, turn.current, turn.over, turn.winner)
		assertPlayerStatus(t, player1, turn.players[0])
		assertPlayerStatus(t, player2, turn.players[1])
		if over, _ := game.IsOver(); !over {
			action.EndTurn(game, game.CurrentPlayer())
		}
	}
}
