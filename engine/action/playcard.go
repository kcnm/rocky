package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/event"
)

func CanPlayCard(
	game base.Game,
	player base.Player,
	cardIndex int,
	target base.Character) (bool, error) {
	if game == nil {
		return false, fmt.Errorf("nil game")
	}
	if player == nil {
		return false, fmt.Errorf("nil player")
	}
	if over, _ := game.IsOver(); over {
		return false, fmt.Errorf("game is over")
	}
	if game.CurrentPlayer() != player {
		return false, fmt.Errorf("it is not player%v's turn", player.ID())
	}
	if cardIndex < 0 || len(player.Hand()) <= cardIndex {
		return false, fmt.Errorf("invalid cardIndex %d", cardIndex)
	}

	c := player.Hand()[cardIndex].(base.Card)
	if c.Mana() > player.Mana() {
		return false, fmt.Errorf("player has insufficient mana %d for card cost %d",
			player.Mana(), c.Mana())
	}

	switch c.(type) {
	case base.MinionCard:
		if player.Board().IsFull() {
			return false, fmt.Errorf("board is full")
		}
		if target == nil {
			break
		}
		if m, ok := target.(base.Minion); !ok || !player.IsControlling(m) {
			return false, fmt.Errorf("invalid minion drop location")
		}
	}
	return true, nil
}

func PlayCard(
	game base.Game,
	player base.Player,
	cardIndex int,
	target base.Character) {
	if ok, err := CanPlayCard(game, player, cardIndex, target); !ok {
		panic(err)
	}
	switch card := player.Hand()[cardIndex].(type) {
	case base.MinionCard:
		toRight := (base.Minion)(nil)
		if target != nil {
			toRight = target.(base.Minion)
		}
		game.Events().PostAndTrigger(
			event.PlayCard(player, cardIndex))
		game.Events().PostAndTrigger(
			event.Summon(game, player, card, player.Board(), toRight))
	}
}
