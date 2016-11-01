package card

import (
	"math/rand"
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
	"github.com/kcnm/rocky/engine/test"
)

func TestLightningStormOnEmptyBoard(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
			Mana:      10,
			Crystal:   10,
			HandSize:  1,
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
	}
	test.PlaySingleSpell(
		LightningStorm,
		status,
		func(g engine.Game) engine.Char {
			return nil
		},
		func(status *test.GameStatus) {
		},
	)(t)
}

func TestLightningStormOnFullBoard(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
			Mana:      10,
			Crystal:   10,
			HandSize:  1,
			BoardSize: 7,
		},
		B1: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, true},
			test.MinionStatus{test.M22, 2, 2, 2, true},
			test.MinionStatus{test.M33, 3, 3, 3, true},
			test.MinionStatus{test.M44, 4, 4, 4, true},
			test.MinionStatus{test.M55, 5, 5, 5, true},
			test.MinionStatus{test.M66, 6, 6, 6, true},
			test.MinionStatus{test.M77, 7, 7, 7, true},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
			BoardSize: 7,
		},
		B2: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, false},
			test.MinionStatus{test.M22, 2, 2, 2, false},
			test.MinionStatus{test.M33, 3, 3, 3, false},
			test.MinionStatus{test.M44, 4, 4, 4, false},
			test.MinionStatus{test.M55, 5, 5, 5, false},
			test.MinionStatus{test.M66, 6, 6, 6, false},
			test.MinionStatus{test.M77, 7, 7, 7, false},
		},
	}

	p1 := game.NewPlayer(1, 30, 30, 0, nil, game.NewDeck(), LightningStorm)
	p2 := game.NewPlayer(2, 30, 30, 0, nil, game.NewDeck())
	p1.GainCrystal(10)
	g := game.Resume(p1, p2, 1, rand.New(rand.NewSource(0)))
	g.Summon(test.M11, p1.Board(), 0)
	g.Summon(test.M22, p1.Board(), 1)
	g.Summon(test.M33, p1.Board(), 2)
	g.Summon(test.M44, p1.Board(), 3)
	g.Summon(test.M55, p1.Board(), 4)
	g.Summon(test.M66, p1.Board(), 5)
	g.Summon(test.M77, p1.Board(), 6)
	b2 := []engine.Minion{
		g.Summon(test.M11, p2.Board(), 0),
		g.Summon(test.M22, p2.Board(), 1),
		g.Summon(test.M33, p2.Board(), 2),
		g.Summon(test.M44, p2.Board(), 3),
		g.Summon(test.M55, p2.Board(), 4),
		g.Summon(test.M66, p2.Board(), 5),
		g.Summon(test.M77, p2.Board(), 6),
	}
	p1.Refresh()
	status.Current = p1

	t.Logf("Resume game")
	test.AssertGameStatus(t, g, status)
	action.PlayCard(g, p1, 0, 0, nil)
	status.P1.Mana -= 3
	status.P1.HandSize -= 1
	destroyed := 0
	for i, m := range b2 {
		damage := m.MaxHealth() - m.Health()
		if damage < 2 || 3 < damage {
			t.Errorf("Out of range damage %d, expected 2-3", damage)
		}
		if m.Health() <= 0 {
			status.B2 = append(status.B2[:i-destroyed], status.B2[i-destroyed+1:]...)
			destroyed++
		} else {
			status.B2[i-destroyed].Health = m.Health()
		}
	}
	status.P2.BoardSize -= destroyed
	test.AssertGameStatus(t, g, status)
}
