package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
)

func PlaySingleCard(
	t *testing.T,
	card engine.Card,
	status GameStatus,
	pos int,
	target func(engine.Game) engine.Char,
	update func(*GameStatus)) func(*testing.T) {
	return func(t *testing.T) {
		p1 := game.NewPlayer(
			1,
			status.P1.Health,
			status.P1.MaxHealth,
			status.P1.Armor,
			nil,
			game.NewDeck(),
			card)
		p2 := game.NewPlayer(
			2,
			status.P2.Health,
			status.P2.MaxHealth,
			status.P2.Armor,
			nil,
			game.NewDeck())
		p1.GainCrystal(10)
		g := game.Resume(p1, p2, 1, nil /* rng */)
		for i, m := range status.B1 {
			g.Summon(m.Card, p1.Board(), i)
		}
		for i, m := range status.B2 {
			g.Summon(m.Card, p2.Board(), i)
		}
		p1.Refresh()
		status.Current = p1

		t.Logf("Resume game")
		AssertGameStatus(t, g, status)
		action.PlayCard(g, p1, 0, pos, target(g))
		status.P1.Mana -= card.Mana()
		status.P1.HandSize -= 1
		copyBoard(&status)
		update(&status)
		AssertGameStatus(t, g, status)
	}
}

func PlaySingleSpell(
	t *testing.T,
	card engine.SpellCard,
	status GameStatus,
	target func(engine.Game) engine.Char,
	update func(*GameStatus)) func(*testing.T) {
	return PlaySingleCard(t, card, status, 0, target, update)
}

func copyBoard(status *GameStatus) {
	b1 := make([]MinionStatus, len(status.B1))
	copy(b1, status.B1)
	b1, status.B1 = status.B1, b1
	b2 := make([]MinionStatus, len(status.B2))
	copy(b2, status.B2)
	b2, status.B2 = status.B2, b2
}
