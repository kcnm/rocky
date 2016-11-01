package card

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/test"
)

func TestFireball(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
			Mana:      10,
			Crystal:   10,
			HandSize:  1,
			BoardSize: 3,
		},
		B1: []test.MinionStatus{
			{test.M11, 1, 1, 1, true},
			{test.M66, 6, 6, 6, true},
			{test.M88, 8, 8, 8, true},
		},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
			BoardSize: 3,
		},
		B2: []test.MinionStatus{
			{test.M11, 1, 1, 1, false},
			{test.M66, 6, 6, 6, false},
			{test.M88, 8, 8, 8, false},
		},
	}

	for _, act := range []struct {
		name   string
		target func(engine.Game) engine.Char
		update func(*test.GameStatus)
	}{
		{
			"Cast on your Hero",
			func(g engine.Game) engine.Char {
				return g.CurrentPlayer()
			},
			func(status *test.GameStatus) {
				status.P1.Health -= 6
			},
		},
		{
			"Cast on your 1/1 Minion",
			func(g engine.Game) engine.Char {
				return g.CurrentPlayer().Board().Get(0)
			},
			func(status *test.GameStatus) {
				status.P1.BoardSize -= 1
				status.B1 = status.B1[1:]
			},
		},
		{
			"Cast on your 6/6 Minion",
			func(g engine.Game) engine.Char {
				return g.CurrentPlayer().Board().Get(1)
			},
			func(status *test.GameStatus) {
				status.P1.BoardSize -= 1
				status.B1 = append(status.B1[:1], status.B1[2:]...)
			},
		},
		{
			"Cast on your 8/8 Minion",
			func(g engine.Game) engine.Char {
				return g.CurrentPlayer().Board().Get(2)
			},
			func(status *test.GameStatus) {
				status.B1[2].Health -= 6
			},
		},
		{
			"Cast on opponent Hero",
			func(g engine.Game) engine.Char {
				return g.Opponent(g.CurrentPlayer())
			},
			func(status *test.GameStatus) {
				status.P2.Health -= 6
			},
		},
		{
			"Cast on opponent 1/1 Minion",
			func(g engine.Game) engine.Char {
				return g.Opponent(g.CurrentPlayer()).Board().Get(0)
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 1
				status.B2 = status.B2[1:]
			},
		},
		{
			"Cast on opponent 6/6 Minion",
			func(g engine.Game) engine.Char {
				return g.Opponent(g.CurrentPlayer()).Board().Get(1)
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 1
				status.B2 = append(status.B2[:1], status.B2[2:]...)
			},
		},
		{
			"Cast on opponent 8/8 Minion",
			func(g engine.Game) engine.Char {
				return g.Opponent(g.CurrentPlayer()).Board().Get(2)
			},
			func(status *test.GameStatus) {
				status.B2[2].Health -= 6
			},
		},
	} {
		t.Run(act.name, test.PlaySingleSpell(
			t, Fireball, status, act.target, act.update))
	}
}
