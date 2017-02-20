package card

import (
	"testing"

	"github.com/kcnm/rocky/engine/test"
)

func TestElvenArcher(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			{test.M11, 1, 1, 1, false},
			{test.M22, 2, 2, 2, false},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			{test.M11, 1, 1, 1, false},
			{test.M22, 2, 2, 2, false},
		},
	}

	for _, act := range []struct {
		name   string
		target test.TargetFn
		update test.UpdateFn
	}{
		{
			"Deal 1 damage to your Hero",
			test.YourHero,
			func(status *test.GameStatus) {
				status.P1.Health -= 1
			},
		},
		{
			"Deal 1 damage to your 1/1 Minion",
			test.YourMinion(0),
			func(status *test.GameStatus) {
				status.B1 = append(status.B1[:1], status.B1[2:]...)
			},
		},
		{
			"Deal 1 damage to your 2/2 Minion",
			test.YourMinion(1),
			func(status *test.GameStatus) {
				status.B1[2].Health -= 1
			},
		},
		{
			"Deal 1 damage to enemy Hero",
			test.EnemyHero,
			func(status *test.GameStatus) {
				status.P2.Health -= 1
			},
		},
		{
			"Deal 1 damage to enemy 1/1 Minion",
			test.EnemyMinion(0),
			func(status *test.GameStatus) {
				status.B2 = status.B2[1:]
			},
		},
		{
			"Deal 1 damage to enemy 2/2 Minion",
			test.EnemyMinion(1),
			func(status *test.GameStatus) {
				status.B2[1].Health -= 1
			},
		},
	} {
		t.Run(act.name, test.PlayCard(
			ElvenArcher(), status, act.target, act.update))
	}
}
