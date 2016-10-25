package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
	"github.com/kcnm/rocky/engine/target"
)

var (
	m11 = NewMinionCard(engine.Neutral, 1, 1, 1)
	m45 = NewMinionCard(engine.Neutral, 4, 4, 5)
	s4  = NewSpellCard(engine.Neutral, 4, target.Manual, target.Any, target.Char, nil)
	w32 = NewWeaponCard(engine.Neutral, 2, 3, 2)
)

func TestBasicGame(t *testing.T) {
	p1 := game.NewPlayer(
		1,  // id
		30, // health
		0,  // armor
		game.NewDeck(m11, m11, m11, m11),
	)
	p2 := game.NewPlayer(
		2,  // id
		30, // health
		0,  // armor
		game.NewDeck(m11, m11, m11, m11),
	)
	g := game.New(p1, p2, nil /* rng */)

	for _, turn := range []struct {
		current engine.Player
		over    bool
		winner  engine.Player
		players []playerStatus
	}{
		{
			// turn 1
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 0, 0, 0, 4, 0},
			},
		},
		{
			// turn 2
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 3
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 4
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 5
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 6
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 7
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 8
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 9
			p1, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 10
			p2, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 11
			p1, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 12
			p2, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 13
			p1, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 14
			p2, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 15
			p1, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 16
			p2, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 17
			p1, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 18
			p2, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 19
			p1, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 20
			p2, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 21
			p1, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 22
			p2, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 23
			p1, true, p2,
			[]playerStatus{
				{-6, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
	} {
		t.Logf("Turn %d", g.Turn())
		assertGameStatus(t, g, turn.current, turn.over, turn.winner)
		assertPlayerStatus(t, p1, turn.players[0])
		assertPlayerStatus(t, p2, turn.players[1])
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
	p1 := game.NewPlayer(1, 30, 0, game.NewDeck(), m45, m11, m45)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)

	// turn 1
	assertPlayerStatus(t, p1, playerStatus{29, 0, 0, false, 10, 10, 3, 0, 0})
	// Plays the first 4/5 minion
	action.PlayCard(g, p1, 0, 0, nil)
	assertPlayerStatus(t, p1, playerStatus{29, 0, 0, false, 6, 10, 2, 0, 1})
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	action.EndTurn(g, p1)

	// turn 2
	action.EndTurn(g, p2)

	// turn 3
	assertPlayerStatus(t, p1, playerStatus{27, 0, 0, false, 10, 10, 2, 0, 1})
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, true})
	// Plays the second 4/5 minion at position 0
	action.PlayCard(g, p1, 1, 0, nil)
	assertPlayerStatus(t, p1, playerStatus{27, 0, 0, false, 6, 10, 1, 0, 2})
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})
	// Plays the 1/1 minion at position 2
	action.PlayCard(g, p1, 0, 2, nil)
	assertPlayerStatus(t, p1, playerStatus{27, 0, 0, false, 5, 10, 0, 0, 3})
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(2), minionStatus{m11, 1, 1, false})
}

func TestCastSpell(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, game.NewDeck(), s4)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	assertPlayerStatus(t, p1, playerStatus{30, 0, 0, false, 10, 10, 1, 0, 0})
	action.PlayCard(g, p1, 0, 0, p2)
	assertPlayerStatus(t, p1, playerStatus{30, 0, 0, false, 6, 10, 0, 0, 0})
}

func TestEquipWeapon(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, game.NewDeck(), w32)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	assertPlayerStatus(t, p1, playerStatus{30, 0, 0, false, 10, 10, 1, 0, 0})
	action.PlayCard(g, p1, 0, 0, nil)
	assertPlayerStatus(t, p1, playerStatus{30, 0, 3, true, 8, 10, 0, 0, 0})
	assertWeaponStatus(t, p1, weaponStatus{w32, 3, 2})
}
