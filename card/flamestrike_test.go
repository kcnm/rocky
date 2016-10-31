package card

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/test"
)

func TestFlamestrike(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:   30,
			Mana:     10,
			Crystal:  10,
			HandSize: 1,
		},
		B1: []test.MinionStatus{},
		P2: test.PlayerStatus{
			Health: 30,
		},
		B2: []test.MinionStatus{},
	}
	target := func(g engine.Game) engine.Char {
		return nil
	}

	for _, act := range []struct {
		name   string
		setup  func()
		update func(*test.GameStatus)
	}{
		{
			"Cast on empty Board",
			func() {},
			func(status *test.GameStatus) {},
		},
		{
			"Cast on 1-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M11, 1, 1, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M11, 1, 1, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 1
				status.B2 = status.B2[1:]
			},
		},
		{
			"Cast on 2-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M22, 2, 2, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M22, 2, 2, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 2
				status.B2 = status.B2[2:]
			},
		},
		{
			"Cast on 3-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M33, 3, 3, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M33, 3, 3, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 3
				status.B2 = status.B2[3:]
			},
		},
		{
			"Cast on 4-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M44, 4, 4, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M44, 4, 4, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 4
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 5-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M55, 5, 5, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M55, 5, 5, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 4
				status.B2[4].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 6-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M66, 6, 6, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M66, 6, 6, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 4
				status.B2[4].Health -= 4
				status.B2[5].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 7-Minion Board",
			func() {
				status.P1.BoardSize += 1
				status.B1 = append(status.B1, test.MinionStatus{test.M77, 7, 7, true})
				status.P2.BoardSize += 1
				status.B2 = append(status.B2, test.MinionStatus{test.M77, 7, 7, false})
			},
			func(status *test.GameStatus) {
				status.P2.BoardSize -= 4
				status.B2[4].Health -= 4
				status.B2[5].Health -= 4
				status.B2[6].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
	} {
		act.setup()
		t.Run(act.name, test.PlaySingleSpell(
			t, Flamestrike, status, target, act.update))
	}
}
