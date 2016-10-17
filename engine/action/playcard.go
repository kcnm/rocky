package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/base/target"
	"github.com/kcnm/rocky/engine/event"
)

func CanPlayCard(
	game base.Game,
	player base.Player,
	cardIndex int,
	position int,
	tgt base.Character) (bool, error) {
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

	switch card := c.(type) {
	case base.MinionCard:
		if ok, err := canPlayMinion(player, position); !ok {
			return false, err
		}
	case base.SpellCard:
		if ok, err := canPlaySpell(game, player, card, tgt); !ok {
			return false, err
		}
	}
	return true, nil
}

func PlayCard(
	game base.Game,
	player base.Player,
	cardIndex int,
	position int,
	tgt base.Character) {
	if ok, err := CanPlayCard(game, player, cardIndex, position, tgt); !ok {
		panic(err)
	}
	c := player.Hand()[cardIndex]
	game.Events().PostAndTrigger(
		event.PlayCard(player, cardIndex))
	switch card := c.(type) {
	case base.MinionCard:
		game.Events().PostAndTrigger(
			event.Summon(game, player, card, player.Board(), position))
	case base.SpellCard:
		game.Events().PostAndTrigger(
			event.Cast(game, player, card, tgt))
	}
}

func canPlayMinion(player base.Player, position int) (bool, error) {
	if player.Board().IsFull() {
		return false, fmt.Errorf("board is full")
	}
	if position < 0 || len(player.Board().Minions()) < position {
		return false, fmt.Errorf("invalid minion drop location")
	}
	return true, nil
}

func canPlaySpell(
	game base.Game,
	player base.Player,
	card base.SpellCard,
	tgt base.Character) (bool, error) {
	if card.Assign() == target.Manual {
		opponent := game.Opponent(player)
		if card.Side() == target.Friend && !player.IsControlling(tgt) {
			return false, fmt.Errorf("target is not friend")
		}
		if card.Side() == target.Enemy && !opponent.IsControlling(tgt) {
			return false, fmt.Errorf("target is not enemy")
		}
		if _, ok := tgt.(base.Minion); card.Role() == target.Minion && !ok {
			return false, fmt.Errorf("target is not minion")
		}
		if _, ok := tgt.(base.Player); card.Role() == target.Player && !ok {
			return false, fmt.Errorf("target is not player")
		}
	} else if tgt != nil {
		return false, fmt.Errorf("unexpected target")
	}
	return true, nil
}
