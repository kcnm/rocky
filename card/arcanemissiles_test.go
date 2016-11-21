package card

import (
	"fmt"
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/test"
)

func TestArcaneMissilesOnEmptyBoard(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
	}
	test.PlayCard(
		ArcaneMissiles,
		status,
		test.NilTarget,
		func(status *test.GameStatus) {
			status.P2.Health -= 3
		},
	)(t)
}

func TestArcaneMissiles(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, false},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, false},
		},
	}

	b2Size := make(map[int]bool)
	satisfy := false
	record := func(g engine.Game, status *test.GameStatus) error {
		op := g.Opponent(g.CurrentPlayer())
		hp := op.Health()
		b2 := op.Board().Minions()
		status.P2.Health = op.Health()
		status.B2 = status.B2[:len(b2)]
		for _, m := range b2 {
			hp += m.Health()
		}
		if hp != 28 {
			return fmt.Errorf("Total health of P2's Chars is %d, expected 28", hp)
		}
		test.AssertGameStatus(t, g, *status)
		b2Size[len(status.B2)] = true
		satisfy = len(b2Size) == 2
		return nil
	}
	test.PlayCardWithRNG(
		t,
		ArcaneMissiles,
		status,
		test.NilTarget,
		record,
		&satisfy)
}
