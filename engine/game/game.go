package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

type game struct {
	*entity
	engine.EventQueue
	rng     *rand.Rand
	players []engine.Player
	current engine.Player
	turn    int
	over    bool
	winner  engine.Player
	idGen   engine.EntityID
}

func New(p1, p2 engine.Player, rng *rand.Rand) engine.Game {
	// Initial checks.
	if p1 == nil {
		panic("nil player1")
	}
	if p2 == nil {
		panic("nil player2")
	}
	if p1.Health() <= 0 {
		panic(fmt.Errorf("non-positive player1 health %d", p1.Health()))
	}
	if p2.Health() <= 0 {
		panic(fmt.Errorf("non-positive player2 health %d", p2.Health()))
	}
	if rng == nil {
		rng = rand.New(rand.NewSource(time.Now().Unix()))
	}
	// Instantiate game.
	g := &game{
		newEntity(0).(*entity),
		event.NewQueue(),
		rng,
		[]engine.Player{p2, p1},
		p1,    // current
		0,     // turn
		false, // over
		nil,   // winner
		0,     // idGen
	}
	// Assign player entity IDs.
	p1.(*player).id = g.nextEntityID()
	p2.(*player).id = g.nextEntityID()
	// Join base in-game entities.
	g.AppendReactor(g.react)
	p1.AppendReactor(p1.(*player).react)
	p2.AppendReactor(p2.(*player).react)
	g.Join(g)
	g.Join(p1)
	g.Join(p2)
	return g
}

func (g *game) RNG() *rand.Rand {
	return g.rng
}

func (g *game) Turn() int {
	return g.turn
}

func (g *game) CurrentPlayer() engine.Player {
	return g.current
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
	chars := make([]engine.Char, 0, 2+len(b1)+len(b2))
	chars = append(chars, g.players[1])
	chars = append(chars, g.players[0])
	for _, m := range b1 {
		chars = append(chars, m)
	}
	for _, m := range b2 {
		chars = append(chars, m)
	}
	return chars
}

func (g *game) IsOver() (over bool, winner engine.Player) {
	return g.over, g.winner
}

func (g *game) Start() {
	g.Fire(event.StartGame(g))
}

func (g *game) Summon(
	card engine.MinionCard,
	player engine.Player,
	position int) engine.Minion {
	minion := newMinion(g.nextEntityID(), card)
	g.Join(minion)
	card.Buff().Apply(g, player, minion)
	return player.Board().Put(minion, position)
}

func (g *game) nextEntityID() engine.EntityID {
	g.idGen++
	return g.idGen
}

func (g *game) react(ev engine.Event) {
	switch ev.Verb() {
	case engine.StartTurn:
		g.onStartTurn(ev)
	case engine.Destroy:
		g.onDestroy(ev)
	case engine.GameOver:
		g.Drain()
	case engine.Combined:
		for _, ev := range ev.Subject().([]engine.Event) {
			g.react(ev)
		}
	}
}

func (g *game) onStartTurn(ev engine.Event) {
	if ev.Verb() != engine.StartTurn {
		return
	}
	g.turn++
	player := ev.Subject().(engine.Player)
	g.current = player
	g.Post(event.Draw(player), ev)
	if !player.HasMaxCrystal() {
		player.GainCrystal(1)
	}
	player.Refresh()
}

func (g *game) onDestroy(ev engine.Event) {
	if ev.Verb() != engine.Destroy {
		return
	}
	if player, ok := ev.Subject().(engine.Player); ok {
		g.over = true
		if g.winner == nil {
			g.winner = g.Opponent(player)
		} else {
			g.winner = nil
		}
		g.Post(event.GameOver(g), ev)
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
