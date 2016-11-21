package card

import (
	"testing"

	"github.com/kcnm/rocky/engine/test"
)

func TestLeperGnome(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			{LeperGnome, 1, 1, 1, false},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			{test.M01, 0, 1, 1, false},
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
			"Attack enemy Hero",
			test.EnemyHero,
			func(status *test.GameStatus) {
				status.B1[0].Active = false
				status.P2.Health -= 1
			},
		},
		{
			"Attack enemy 0/1 Minion",
			test.EnemyMinion(0),
			func(status *test.GameStatus) {
				status.B1[0].Active = false
				status.B2 = status.B2[1:]
			},
		},
		{
			"Attack enemy 1/1 Minion",
			test.EnemyMinion(1),
			func(status *test.GameStatus) {
				status.P2.Health -= 2
				status.B1 = status.B1[1:]
				status.B2 = append(status.B2[:1], status.B2[2:]...)
			},
		},
		{
			"Attack enemy 2/2 Minion",
			test.EnemyMinion(2),
			func(status *test.GameStatus) {
				status.P2.Health -= 2
				status.B1 = status.B1[1:]
				status.B2[2].Health -= 1
			},
		},
	} {
		t.Run(act.name, test.MinionAttack(
			status, act.target, act.update))
	}
}

func TestLeperGnomeDual(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    3, // +1 fatigue
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{
			{LeperGnome, 1, 1, 1, false},
		},
		P2: test.PlayerStatus{
			Health:    2,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{
			{LeperGnome, 1, 1, 1, false},
		},
	}
	test.MinionAttack(
		status,
		test.EnemyMinion(0),
		func(status *test.GameStatus) {
			status.Over = true
			status.Winner = nil
		},
	)(t)
}
