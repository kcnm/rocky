package card

import (
	"fmt"
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/test"
)

func TestLightningStormOnEmptyBoard(t *testing.T) {
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
		LightningStorm(),
		status,
		test.NilTarget,
		test.NilUpdate,
	)(t)
}

func TestLightningStorm(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, false},
			test.MinionStatus{test.M22, 2, 2, 2, false},
			test.MinionStatus{test.M33, 3, 3, 3, false},
			test.MinionStatus{test.M44, 4, 4, 4, false},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			test.MinionStatus{test.M11, 1, 1, 1, false},
			test.MinionStatus{test.M22, 2, 2, 2, false},
			test.MinionStatus{test.M33, 3, 3, 3, false},
			test.MinionStatus{test.M44, 4, 4, 4, false},
		},
	}

	b2Size := make(map[int]bool)
	satisfy := false
	record := func(g engine.Game, status *test.GameStatus) error {
		b2 := g.Opponent(g.CurrentPlayer()).Board().Minions()
		status.B2 = status.B2[len(status.B2)-len(b2):]
		for i, m := range b2 {
			dmg := m.MaxHealth() - m.Health()
			if dmg < 2 || 3 < dmg {
				return fmt.Errorf("Minion%d took %d damage, expected 2-3", m.ID(), dmg)
			}
			status.B2[i].Health -= dmg
		}
		test.AssertGameStatus(t, g, *status)
		b2Size[len(status.B2)] = true
		satisfy = len(b2Size) == 2
		return nil
	}
	test.PlayCardWithRNG(
		t,
		LightningStorm(),
		status,
		test.NilTarget,
		record,
		&satisfy)
}
