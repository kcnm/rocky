package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type game struct {
	engine.EventBus
	players []engine.Player
	turn    int
	over    bool
	winner  engine.Player
	idGen   engine.CharID
	rng     *rand.Rand
}

func New(player1, player2 engine.Player, rng *rand.Rand) engine.Game {
	// Initial checks.
	if player1 == nil {
		panic("nil player1")
	}
	if player2 == nil {
		panic("nil player2")
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
	// Instantiate game.
	g := &game{
		event.NewBus(),
		[]engine.Player{player2, player1},
		0,     // turn
		false, // over
		nil,   // winner
		0,     // idGen
		rng,
	}
	// Assign player character IDs.
	player1.(*player).id = g.nextCharID()
	player2.(*player).id = g.nextCharID()
	// Register event listeners.
	g.AddListener(g)
	g.AddListener(player1)
	g.AddListener(player2)
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
	case engine.GameOver:
		g.Drain()
	case engine.Combined:
		for _, e := range ev.Subject().([]engine.Event) {
			g.Handle(e)
		}
	}
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

func (g *game) Start() {
	if g.turn != 0 {
		panic(fmt.Errorf("non-zero start turn %d", g.turn))
	}
	g.Fire(event.StartTurn())
}

func (g *game) Summon(
	card engine.MinionCard,
	player engine.Player,
	position int) engine.Minion {
	if player.Board().IsFull() {
		return nil
	}
	minion := newMinion(g.nextCharID(), card)
	g.AddListener(minion)
	card.Buff().Apply(g, player, minion)
	return player.Board().Put(minion, position)
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
	for i := 0; i < 2; i++ {
		if ev.Subject() == g.players[i] {
			g.over = true
			if g.winner == nil {
				g.winner = g.players[(i+1)%2]
			} else {
				g.winner = nil
			}
			g.Post(event.GameOver(), ev)
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
