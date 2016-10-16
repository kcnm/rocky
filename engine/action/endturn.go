package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/event"
)

func CanEndTurn(game base.Game, player base.Player) (bool, error) {
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
	return true, nil
}

func EndTurn(game base.Game, player base.Player) {
	if ok, err := CanEndTurn(game, player); !ok {
		panic(err)
	}
	game.Events().PostAndTrigger(
		event.EndTurn(game))
}
