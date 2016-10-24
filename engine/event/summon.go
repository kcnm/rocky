package event

import (
	"github.com/kcnm/rocky/engine"
)

type summon struct {
	game     engine.Game
	summoner engine.Char
	card     engine.MinionCard
	board    engine.Board
	position int
}

func Summon(
	game engine.Game,
	summoner engine.Char,
	card engine.MinionCard,
	board engine.Board,
	position int) engine.Event {
	return &summon{game, summoner, card, board, position}
}

func (ev *summon) Subject() interface{} {
	return ev.summoner
}

func (ev *summon) Verb() engine.Verb {
	return engine.Summon
}

func (ev *summon) Trigger() {
	ev.game.Summon(ev.card, ev.board, ev.position)
}
