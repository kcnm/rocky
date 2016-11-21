package card

import (
	"testing"

	"github.com/kcnm/rocky/engine/test"
)

func TestFlamestrike(t *testing.T) {
	status := test.GameStatus{
		P1: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B1: []test.MinionStatus{},
		P2: test.PlayerStatus{
			Health:    30,
			MaxHealth: 30,
		},
		B2: []test.MinionStatus{},
	}

	for _, act := range []struct {
		name   string
		setup  func()
		update test.UpdateFn
	}{
		{
			"Cast on empty Board",
			func() {},
			func(status *test.GameStatus) {},
		},
		{
			"Cast on 1-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M11, 1, 1, 1, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M11, 1, 1, 1, false})
			},
			func(status *test.GameStatus) {
				status.B2 = status.B2[1:]
			},
		},
		{
			"Cast on 2-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M22, 2, 2, 2, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M22, 2, 2, 2, false})
			},
			func(status *test.GameStatus) {
				status.B2 = status.B2[2:]
			},
		},
		{
			"Cast on 3-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M33, 3, 3, 3, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M33, 3, 3, 3, false})
			},
			func(status *test.GameStatus) {
				status.B2 = status.B2[3:]
			},
		},
		{
			"Cast on 4-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M44, 4, 4, 4, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M44, 4, 4, 4, false})
			},
			func(status *test.GameStatus) {
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 5-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M55, 5, 5, 5, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M55, 5, 5, 5, false})
			},
			func(status *test.GameStatus) {
				status.B2[4].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 6-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M66, 6, 6, 6, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M66, 6, 6, 6, false})
			},
			func(status *test.GameStatus) {
				status.B2[4].Health -= 4
				status.B2[5].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
		{
			"Cast on 7-Minion Board",
			func() {
				status.B1 = append(status.B1, test.MinionStatus{test.M77, 7, 7, 7, false})
				status.B2 = append(status.B2, test.MinionStatus{test.M77, 7, 7, 7, false})
			},
			func(status *test.GameStatus) {
				status.B2[4].Health -= 4
				status.B2[5].Health -= 4
				status.B2[6].Health -= 4
				status.B2 = status.B2[4:]
			},
		},
	} {
		act.setup()
		t.Run(act.name, test.PlayCard(
			Flamestrike, status, test.NilTarget, act.update))
	}
}
