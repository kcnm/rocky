package test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
)

type TargetFn func(engine.Game) engine.Char

type UpdateFn func(*GameStatus)

type RecordFn func(engine.Game, *GameStatus) error

func NilTarget(game engine.Game) engine.Char {
	return nil
}

func NilUpdate(status *GameStatus) {}

func PlaySingleCard(
	card engine.Card,
	status GameStatus,
	pos int,
	target TargetFn,
	update UpdateFn) func(*testing.T) {
	return func(t *testing.T) {
		g := playSingleCard(t, nil, card, &status, pos, target)
		backupBoard(&status)
		update(&status)
		AssertGameStatus(t, g, status)
	}
}

func PlaySingleCardWithRNG(
	t *testing.T,
	card engine.Card,
	status GameStatus,
	pos int,
	target TargetFn,
	record RecordFn,
	satisfy *bool) {
	bail := false
	for seed := int64(0); bail || !*satisfy; seed++ {
		status := status // creats a local copy
		rng := rand.New(rand.NewSource(seed))
		t.Run(fmt.Sprintf("seed%d", seed), func(t *testing.T) {
			g := playSingleCard(t, rng, card, &status, pos, target)
			backupBoard(&status)
			if err := record(g, &status); err != nil {
				t.Error(err)
				bail = true
			}
		})
	}
}

func newGame(rng *rand.Rand, status *GameStatus) engine.Game {
	p1 := game.NewPlayer(status.P1.MaxHealth, nil, game.NewDeck())
	if status.P1.Health < status.P1.MaxHealth {
		p1.TakeDamage(status.P1.MaxHealth - status.P1.Health)
	}
	if status.P1.Armor > 0 {
		p1.GainArmor(status.P1.Armor)
	}
	p2 := game.NewPlayer(status.P2.MaxHealth, nil, game.NewDeck())
	if status.P2.Health < status.P2.MaxHealth {
		p2.TakeDamage(status.P2.MaxHealth - status.P2.Health)
	}
	if status.P2.Armor > 0 {
		p2.GainArmor(status.P2.Armor)
	}
	g := game.New(p1, p2, rng)
	for i, m := range status.B1 {
		g.Summon(m.Card, p1, i)
	}
	for i, m := range status.B2 {
		g.Summon(m.Card, p2, i)
	}
	status.Current = p1
	return g
}

func playSingleCard(
	t *testing.T,
	rng *rand.Rand,
	card engine.Card,
	status *GameStatus,
	pos int,
	target TargetFn) engine.Game {
	g := newGame(rng, status)
	p1 := status.Current
	p1.GainCrystal(10)
	p1.Deck().PutOnTop(card)
	g.Start()
	AssertGameStatus(t, g, *status)
	action.PlayCard(g, p1, 0, pos, target(g))
	status.P1.Mana -= card.Mana()
	status.P1.HandSize -= 1
	return g
}

func backupBoard(status *GameStatus) {
	b1 := make([]MinionStatus, len(status.B1))
	copy(b1, status.B1)
	b1, status.B1 = status.B1, b1
	b2 := make([]MinionStatus, len(status.B2))
	copy(b2, status.B2)
	b2, status.B2 = status.B2, b2
}
