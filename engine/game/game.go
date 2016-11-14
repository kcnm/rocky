package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type game struct {
	events      engine.EventBus
	listenerIDs map[engine.Char]engine.ListenerID
	players     []engine.Player
	turn        int
	over        bool
	winner      engine.Player
	idGen       engine.CharID
	rng         *rand.Rand
}

func New(player1, player2 engine.Player, rng *rand.Rand) engine.Game {
	g := Resume(player1, player2, 0, rng).(*game)
	player1.Deck().Shuffle(g.rng)
	player2.Deck().Shuffle(g.rng)
	g.events.Fire(engine.StartTurn)
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
		event.NewBus(),
		make(map[engine.Char]engine.ListenerID),
		[]engine.Player{player2, player1},
		turn,
		false, // over
		nil,   // winner
		idGen,
		rng,
	}
	g.events.AddListener(g)
	g.events.AddListener(player1.Board())
	g.events.AddListener(player2.Board())
	g.events.AddListener(player1)
	g.events.AddListener(player2)
	// TODO: Add listeners for existing minions.
	return g
}

func (g *game) Handle(ev engine.Event) {
	switch ev.Verb() {
	case engine.StartTurn:
		g.handleStartTurn(ev)
	case engine.Destroy:
		g.handleDestroy(ev)
	case engine.GameOver:
		g.events.Drain()
	case engine.Combined:
		for _, e := range ev.Subject().([]engine.Event) {
			g.Handle(e)
		}
	}
}

func (g *game) Events() engine.EventBus {
	return g.events
}

func (g *game) RNG() *rand.Rand {
	return g.rng
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

func (g *game) AllChars() []engine.Char {
	b1, b2 := g.players[1].Board().Minions(), g.players[0].Board().Minions()
	chars := make([]engine.Char, 2+len(b1)+len(b2))
	chars[0] = g.players[1]
	chars[1] = g.players[0]
	for i := 0; i < len(b1); i++ {
		chars[2+i] = b1[i]
	}
	for i := 0; i < len(b2); i++ {
		chars[2+i] = b2[i]
	}
	return chars
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
	listenerID := g.events.AddListener(minion)
	g.listenerIDs[minion] = listenerID
	return board.Put(minion, position)
}

func (g *game) nextCharID() engine.CharID {
	g.idGen++
	return g.idGen
}

func (g *game) handleStartTurn(ev engine.Event) {
	g.turn++
	player := g.CurrentPlayer()
	g.events.Post(event.Draw(g, player), ev)
	if !player.HasMaxCrystal() {
		player.GainCrystal(1)
	}
	player.Refresh()
}

func (g *game) handleDestroy(ev engine.Event) {
	minion, ok := ev.Subject().(engine.Minion)
	if ok {
		g.events.RemoveListener(g.listenerIDs[minion])
	}

	for i := 0; i < 2; i++ {
		if ev.Subject() == g.players[i] {
			g.over = true
			if g.winner == nil {
				g.winner = g.players[(i+1)%2]
			} else {
				g.winner = nil
			}
			g.events.Post(engine.GameOver, ev)
		}
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
