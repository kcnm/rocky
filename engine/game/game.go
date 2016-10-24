package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type game struct {
	*engine.EventBus
	players []engine.Player
	turn    int
	over    bool
	winner  engine.Player
	idGen   engine.CharID
	rng     *rand.Rand
}

func New(player1, player2 engine.Player, rng *rand.Rand) engine.Game {
	g := Resume(player1, player2, 0, rng).(*game)
	player1.Deck().Shuffle(g.rng)
	player2.Deck().Shuffle(g.rng)
	g.PostAndTrigger(engine.StartTurn)
	return g
}

func Resume(
	player1 engine.Player,
	player2 engine.Player,
	turn int,
	rng *rand.Rand) engine.Game {
	if player1 == nil {
		panic("nil player1")
	}
	if player2 == nil {
		panic("nil player2")
	}
	if player1.ID() != 1 {
		panic(fmt.Errorf("invalid player1 ID %d", player1.ID()))
	}
	if player2.ID() != 2 {
		panic(fmt.Errorf("invalid player2 ID %d", player2.ID()))
	}
	if player1.Health() <= 0 {
		panic(fmt.Errorf("non-positive player1 health %d", player1.Health()))
	}
	if player2.Health() <= 0 {
		panic(fmt.Errorf("non-positive player2 health %d", player2.Health()))
	}
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().Unix()))
	}
	// Validates character IDs, and initialize ID generator.
	ids, idGen := make(map[engine.CharID]bool), engine.CharID(2)
	for _, p := range []engine.Player{player1, player2} {
		ids[p.ID()] = true
		for _, m := range p.Board().Minions() {
			if ids[m.ID()] {
				panic(fmt.Errorf("duplicated character ID %d", m.ID()))
			}
			ids[m.ID()] = true
			if m.ID() > idGen {
				idGen = m.ID()
			}
		}
	}
	g := &game{
		engine.NewEventBus(),
		[]engine.Player{player2, player1},
		turn,
		false, // over
		nil,   // winner
		idGen,
		rng,
	}
	g.AddListener(g)
	g.AddListener(player1.Board())
	g.AddListener(player2.Board())
	return g
}

func (g *game) Handle(ev engine.Event) {
	switch ev.Verb() {
	case engine.StartTurn:
		g.handleStartTurn(ev)
	case engine.Destroy:
		g.handleDestroy(ev)
	}
}

func (g *game) Events() *engine.EventBus {
	return g.EventBus
}

func (g *game) Turn() int {
	return g.turn
}

func (g *game) CurrentPlayer() engine.Player {
	return g.players[g.turn%2]
}

func (g *game) Opponent(player engine.Player) engine.Player {
	if player == g.players[0] {
		return g.players[1]
	}
	if player == g.players[1] {
		return g.players[0]
	}
	panic("player is not in the game")
}

func (g *game) IsOver() (over bool, winner engine.Player) {
	return g.over, g.winner
}

func (g *game) Summon(
	card engine.MinionCard,
	board engine.Board,
	position int) engine.Minion {
	if board.IsFull() {
		return nil
	}
	minion := newMinion(g.nextCharID(), card)
	return board.Put(minion, position)
}

func (g *game) nextCharID() engine.CharID {
	g.idGen++
	return g.idGen
}

func (g *game) handleStartTurn(ev engine.Event) {
	g.turn++
	player := g.CurrentPlayer()
	g.Post(event.Draw(g, player), ev)
	if !player.HasMaxCrystal() {
		player.GainCrystal(1)
	}
	player.Refresh()
}

func (g *game) handleDestroy(ev engine.Event) {
	p0, p1 := false, false
	subject := ev.Subject()
	if subjects, ok := subject.([]interface{}); ok {
		for _, sub := range subjects {
			if sub == g.players[0] {
				p0 = true
			} else if sub == g.players[1] {
				p1 = true
			}
		}
	} else if subject == g.players[0] {
		p0 = true
	} else if subject == g.players[1] {
		p1 = true
	}
	if p0 || p1 {
		g.over = true
		if !p0 && p1 {
			g.winner = g.players[0]
		} else if p0 && !p1 {
			g.winner = g.players[1]
		}
		g.Post(engine.GameOver, ev)
	}
}

func (g *game) String() string {
	player1, player2 := g.players[0], g.players[1]
	if player1 == g.CurrentPlayer() {
		return fmt.Sprintf("Turn %d\n*%v\n %v\n", g.turn, player1, player2)
	} else {
		return fmt.Sprintf("Turn %d\n %v\n*%v\n", g.turn, player1, player2)
	}
}
