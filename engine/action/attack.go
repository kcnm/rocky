package action

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/event"
)

func CanAttack(
	game engine.Game,
	attacker engine.Char,
	defender engine.Char) (bool, error) {
	if game == nil {
		return false, fmt.Errorf("nil game")
	}
	if attacker == nil {
		return false, fmt.Errorf("nil attacker")
	}
	if defender == nil {
		return false, fmt.Errorf("nil defender")
	}
	if over, _ := game.IsOver(); over {
		return false, fmt.Errorf("game is over")
	}
	if !game.CurrentPlayer().IsControlling(attacker) {
		return false, fmt.Errorf("attacker is not controlled by current player")
	}
	if game.CurrentPlayer().IsControlling(defender) {
		return false, fmt.Errorf("defender is controlled by current player")
	}
	if !attacker.Active() {
		return false, fmt.Errorf("attacker is not active")
	}
	return true, nil
}

func Attack(
	game engine.Game,
	attacker engine.Char,
	defender engine.Char) {
	if ok, err := CanAttack(game, attacker, defender); !ok {
		panic(err)
	}
	game.Events().Fire(
		event.Attack(game, attacker, defender))
}
