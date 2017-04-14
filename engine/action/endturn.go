package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

func CanEndTurn(game engine.Game, player engine.Player) (bool, error) {
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

func EndTurn(game engine.Game, player engine.Player) {
	if ok, err := CanEndTurn(game, player); !ok {
		panic(err)
	}
	game.Fire(event.EndTurn(game, player))
}
