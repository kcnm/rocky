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
		g.turn++
		player := g.CurrentPlayer()
		g.Post(event.Draw(g, player))
		if !player.HasMaxCrystal() {
			player.GainCrystal(1)
		}
		player.Refresh()
	case base.Destroy:
		if ev.Subject() == g.players[0] {
			g.over = true
			g.winner = g.players[1]
		}
		if ev.Subject() == g.players[1] {
			g.over = true
			g.winner = g.players[0]
		}
		if g.over {
			g.Post(base.GameOver)
		}
	}
}

func (g *game) Events() *base.EventBus {
	return g.EventBus
}

func (g *game) CurrentPlayer() base.Player {
	return g.players[g.turn%2]
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

func (g *game) String() string {
	player1, player2 := g.players[0], g.players[1]
	if player1 == g.CurrentPlayer() {
		return fmt.Sprintf("Turn %d\n*%v\n %v\n", g.turn, player1, player2)
	} else {
		return fmt.Sprintf("Turn %d\n %v\n*%v\n", g.turn, player1, player2)
	}
}
