package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

func CanPlayCard(
	game engine.Game,
	player engine.Player,
	cardIndex int,
	position int,
	target engine.Char) (bool, error) {
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

	card := player.Hand()[cardIndex].(engine.Card)
	if card.Mana() > player.Mana() {
		return false, fmt.Errorf("player has insufficient mana %d for card cost %d",
			player.Mana(), card.Mana())
	}

	switch card := card.(type) {
	case engine.MinionCard:
		if ok, err := canPlayMinion(player, position); !ok {
			return false, err
		}
	case engine.SpellCard:
		if ok, err := canPlaySpell(game, player, card, target); !ok {
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
	target engine.Char) {
	if ok, err := CanPlayCard(game, player, cardIndex, position, target); !ok {
		panic(err)
	}
	card := player.Hand()[cardIndex]
	game.Events().Fire(
		event.PlayCard(player, cardIndex))
	switch card := card.(type) {
	case engine.MinionCard:
		game.Events().Fire(
			event.Impact(game, player, target, card.Battlecry()))
		game.Events().Fire(
			event.Summon(game, player, card, position))
	case engine.SpellCard:
		game.Events().Fire(
			event.Cast(game, player, card, target))
	case engine.WeaponCard:
		game.Events().Fire(
			event.Impact(game, player, target, card.Battlecry()))
		if player.Weapon() != nil {
			game.Events().Fire(
				event.DestroyWeapon(game, player))
		}
		game.Events().Fire(
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
	target engine.Char) (bool, error) {
	if !card.Effect().CanHappen(game, player, target) {
		return false, fmt.Errorf("Spell effect cannot happen")
	}
	return true, nil
}
