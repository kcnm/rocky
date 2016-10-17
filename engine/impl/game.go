package impl

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/event"
)

type game struct {
	*base.EventBus
	players []base.Player
	turn    int
	over    bool
	winner  base.Player
	idGen   base.CharacterID
}

func NewGame(player1, player2 base.Player) base.Game {
	if player1 == nil {
		panic("nil player1")
	}
	if player2 == nil {
		panic("nil player2")
	}
	g := &game{
		base.NewEventBus(),
		[]base.Player{player2, player1},
		0,     // turn
		false, // over
		nil,   // winner
		0,     // idGen
	}
	player1.Assign(g.nextCharacterID())
	player2.Assign(g.nextCharacterID())
	player1.Deck().Shuffle()
	player2.Deck().Shuffle()
	g.AddListener(g)
	g.AddListener(player1.Board())
	g.AddListener(player2.Board())
	g.PostAndTrigger(base.StartTurn)
	return g
}

func (g *game) Handle(ev base.Event) {
	switch ev.Verb() {
	case base.StartTurn:
		g.handleStartTurn(ev)
	case base.Destroy:
		g.handleDestroy(ev)
	}
}

func (g *game) Events() *base.EventBus {
	return g.EventBus
}

func (g *game) CurrentPlayer() base.Player {
	return g.players[g.turn%2]
}

func (g *game) Opponent(player base.Player) base.Player {
	if player == g.players[0] {
		return g.players[1]
	}
	if player == g.players[1] {
		return g.players[0]
	}
	panic("player is not in the game")
}

func (g *game) IsOver() (over bool, winner base.Player) {
	return g.over, g.winner
}

func (g *game) Summon(
	card base.MinionCard,
	board base.Board,
	toRight base.Minion) base.Minion {
	if board.IsFull() {
		return nil
	}
	minion := newMinion(g.nextCharacterID(), card)
	return board.Put(minion, toRight)
}

func (g *game) nextCharacterID() base.CharacterID {
	g.idGen++
	return g.idGen
}

func (g *game) handleStartTurn(ev base.Event) {
	g.turn++
	player := g.CurrentPlayer()
	g.Post(event.Draw(g, player), ev)
	if !player.HasMaxCrystal() {
		player.GainCrystal(1)
	}
	player.Refresh()
}

func (g *game) handleDestroy(ev base.Event) {
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
		g.Post(base.GameOver, ev)
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
