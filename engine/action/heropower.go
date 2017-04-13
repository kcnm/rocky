package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

func CanHeroPower(
	game engine.Game,
	player engine.Player,
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
	if player.Powered() {
		return false, fmt.Errorf("player has already used hero power this turn")
	}
	if player.Mana() < player.Power().Mana() {
		return false, fmt.Errorf("player does not have enough mana to hero power")
	}
	return true, nil
}

func HeroPower(
	game engine.Game,
	player engine.Player,
	target engine.Char) {
	if ok, err := CanHeroPower(game, player, target); !ok {
		panic(err)
	}
	game.Fire(
		event.HeroPower(game, player, target))
}
