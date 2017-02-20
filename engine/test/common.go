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

func YourHero(game engine.Game) engine.Char {
	return game.CurrentPlayer()
}

func EnemyHero(game engine.Game) engine.Char {
	return game.Opponent(game.CurrentPlayer())
}

func YourMinion(pos int) TargetFn {
	return func(game engine.Game) engine.Char {
		return game.CurrentPlayer().Board().Get(pos)
	}
}

func EnemyMinion(pos int) TargetFn {
	return func(game engine.Game) engine.Char {
		return game.Opponent(game.CurrentPlayer()).Board().Get(pos)
	}
}

func NilUpdate(status *GameStatus) {}

func PlayCard(
	card engine.Card,
	status GameStatus,
	target TargetFn,
	update UpdateFn) func(*testing.T) {
	return func(t *testing.T) {
		g := playCard(t, nil, &status, card, target)
		backupBoard(&status)
		update(&status)
		AssertGameStatus(t, g, status)
	}
}

func PlayCardWithRNG(
	t *testing.T,
	card engine.Card,
	status GameStatus,
	target TargetFn,
	record RecordFn,
	satisfy *bool) {
	bail := false
	for seed := int64(0); bail || !*satisfy; seed++ {
		status := status // creats a local copy
		rng := rand.New(rand.NewSource(seed))
		t.Run(fmt.Sprintf("seed%d", seed), func(t *testing.T) {
			g := playCard(t, rng, &status, card, target)
			backupBoard(&status)
			if err := record(g, &status); err != nil {
				t.Error(err)
				bail = true
			}
		})
	}
}

func MinionAttack(
	status GameStatus,
	target TargetFn,
	update UpdateFn) func(*testing.T) {
	return func(t *testing.T) {
		g := attack(t, nil, &status, YourMinion(0), target)
		backupBoard(&status)
		update(&status)
		AssertGameStatus(t, g, status)
	}
}

func newGame(rng *rand.Rand, status *GameStatus) engine.Game {
	p1 := game.NewPlayer(status.P1.MaxHealth, nil, game.NewDeck())
	p2 := game.NewPlayer(status.P2.MaxHealth, nil, game.NewDeck())
	for _, p := range []struct {
		player engine.Player
		status PlayerStatus
	}{
		{p1, status.P1},
		{p2, status.P2},
	} {
		if p.status.Crystal > 0 {
			p.player.GainCrystal(p.status.Crystal)
		}
		if p.status.Armor > 0 {
			p.player.GainArmor(p.status.Armor)
		}
		if dmg := p.status.MaxHealth - p.status.Health; dmg > 0 {
			p.player.TakeDamage(dmg)
		}
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

func playCard(
	t *testing.T,
	rng *rand.Rand,
	status *GameStatus,
	card engine.Card,
	target TargetFn) engine.Game {
	g := newGame(rng, status)
	p1 := status.Current
	p1.GainCrystal(10)
	p1.Deck().PutOnTop(card)
	g.Start()
	status.P1.Crystal = p1.Crystal()
	status.P1.Mana = p1.Mana()
	status.P1.HandSize = 1
	for i := range status.B1 {
		status.B1[i].Active = true
	}
	AssertGameStatus(t, g, *status)
	action.PlayCard(g, p1, 0, 0, target(g))
	status.P1.Mana -= card.Mana()
	status.P1.HandSize -= 1
	if c, ok := card.(engine.MinionCard); ok {
		status.B1 = append(
			[]MinionStatus{{c, c.Attack(), c.Health(), c.Health(), false}},
			status.B1...)
	}
	return g
}

func attack(
	t *testing.T,
	rng *rand.Rand,
	status *GameStatus,
	attacker TargetFn,
	defender TargetFn) engine.Game {
	g := newGame(rng, status)
	p1 := status.Current
	g.Start()
	status.P1.Health -= 1 // fatigue
	status.P1.Crystal = p1.Crystal()
	status.P1.Mana = p1.Mana()
	for i := range status.B1 {
		status.B1[i].Active = true
	}
	AssertGameStatus(t, g, *status)
	action.Attack(g, attacker(g), defender(g))
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
