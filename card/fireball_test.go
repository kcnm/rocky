package card

import (
	"testing"

	"github.com/kcnm/rocky/engine/test"
)

func TestFireball(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			{test.M11, 1, 1, 1, false},
			{test.M66, 6, 6, 6, false},
			{test.M88, 8, 8, 8, false},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			{test.M11, 1, 1, 1, false},
			{test.M66, 6, 6, 6, false},
			{test.M88, 8, 8, 8, false},
		},
	}

	for _, act := range []struct {
		name   string
		target test.TargetFn
		update test.UpdateFn
	}{
		{
			"Cast on your Hero",
			test.YourHero,
			func(status *test.GameStatus) {
				status.P1.Health -= 6
			},
		},
		{
			"Cast on your 1/1 Minion",
			test.YourMinion(0),
			func(status *test.GameStatus) {
				status.B1 = status.B1[1:]
			},
		},
		{
			"Cast on your 6/6 Minion",
			test.YourMinion(1),
			func(status *test.GameStatus) {
				status.B1 = append(status.B1[:1], status.B1[2:]...)
			},
		},
		{
			"Cast on your 8/8 Minion",
			test.YourMinion(2),
			func(status *test.GameStatus) {
				status.B1[2].Health -= 6
			},
		},
		{
			"Cast on enemy Hero",
			test.EnemyHero,
			func(status *test.GameStatus) {
				status.P2.Health -= 6
			},
		},
		{
			"Cast on enemy 1/1 Minion",
			test.EnemyMinion(0),
			func(status *test.GameStatus) {
				status.B2 = status.B2[1:]
			},
		},
		{
			"Cast on enemy 6/6 Minion",
			test.EnemyMinion(1),
			func(status *test.GameStatus) {
				status.B2 = append(status.B2[:1], status.B2[2:]...)
			},
		},
		{
			"Cast on enemy 8/8 Minion",
			test.EnemyMinion(2),
			func(status *test.GameStatus) {
				status.B2[2].Health -= 6
			},
		},
	} {
		t.Run(act.name, test.PlayCard(
			Fireball(), status, act.target, act.update))
	}
}
