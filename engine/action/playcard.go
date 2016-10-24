package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
	"github.com/kcnm/rocky/engine/target"
)

func CanPlayCard(
	game engine.Game,
	player engine.Player,
	cardIndex int,
	position int,
	tgt engine.Char) (bool, error) {
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

	c := player.Hand()[cardIndex].(engine.Card)
	if c.Mana() > player.Mana() {
		return false, fmt.Errorf("player has insufficient mana %d for card cost %d",
			player.Mana(), c.Mana())
	}

	switch card := c.(type) {
	case engine.MinionCard:
		if ok, err := canPlayMinion(player, position); !ok {
			return false, err
		}
	case engine.SpellCard:
		if ok, err := canPlaySpell(game, player, card, tgt); !ok {
			return false, err
		}
	}
	return true, nil
}

func PlayCard(
	game engine.Game,
	player engine.Player,
	cardIndex int,
	position int,
	tgt engine.Char) {
	if ok, err := CanPlayCard(game, player, cardIndex, position, tgt); !ok {
		panic(err)
	}
	c := player.Hand()[cardIndex]
	game.Events().PostAndTrigger(
		event.PlayCard(player, cardIndex))
	switch card := c.(type) {
	case engine.MinionCard:
		game.Events().PostAndTrigger(
			event.Summon(game, player, card, player.Board(), position))
	case engine.SpellCard:
		game.Events().PostAndTrigger(
			event.Cast(game, player, card, tgt))
	case engine.WeaponCard:
		if player.Weapon() != nil {
			game.Events().PostAndTrigger(
				event.DestroyWeapon(game, player))
		}
		game.Events().PostAndTrigger(
			event.Equip(game, player, card))
	}
}

func canPlayMinion(player engine.Player, position int) (bool, error) {
	if player.Board().IsFull() {
		return false, fmt.Errorf("board is full")
	}
	if position < 0 || len(player.Board().Minions()) < position {
		return false, fmt.Errorf("invalid minion drop location")
	}
	return true, nil
}

func canPlaySpell(
	game engine.Game,
	player engine.Player,
	card engine.SpellCard,
	tgt engine.Char) (bool, error) {
	if card.Assign() == target.Manual {
		opponent := game.Opponent(player)
		if card.Side() == target.Friend && !player.IsControlling(tgt) {
			return false, fmt.Errorf("target is not friend")
		}
		if card.Side() == target.Enemy && !opponent.IsControlling(tgt) {
			return false, fmt.Errorf("target is not enemy")
		}
		if _, ok := tgt.(engine.Minion); card.Role() == target.Minion && !ok {
			return false, fmt.Errorf("target is not minion")
		}
		if _, ok := tgt.(engine.Player); card.Role() == target.Player && !ok {
			return false, fmt.Errorf("target is not player")
		}
	} else if tgt != nil {
		return false, fmt.Errorf("unexpected target")
	}
	return true, nil
}
