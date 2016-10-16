package event

import (
	"github.com/kcnm/rocky/engine/base"
)

type summon struct {
	game     base.Game
	summoner base.Character
	card     base.MinionCard
	board    base.Board
	toRight  base.Minion
}

func Summon(
	game base.Game,
	summoner base.Character,
	card base.MinionCard,
	board base.Board,
	toRight base.Minion) base.Event {
	return &summon{game, summoner, card, board, toRight}
}

func (ev *summon) Subject() interface{} {
	return ev.summoner
}

func (ev *summon) Verb() base.Verb {
	return base.Summon
}

func (ev *summon) Trigger() {
	ev.game.Summon(ev.card, ev.board, ev.toRight)
}
