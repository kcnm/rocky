package effect

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/param"
)

type splitDamage struct {
	damage param.Param
	target param.Param
}

func SplitDamage(damage param.Param, target param.Param) engine.Effect {
	return &splitDamage{damage, target}
}

func (e splitDamage) CanHappen(
	game engine.Game,
	you engine.Player,
	target engine.Char) bool {
	dmg := e.damage.Eval(game, you, target)
	t := e.target.Eval(game, you, target)
	return dmg.(int) > 0 && t != nil
}

func (e splitDamage) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	damage := e.damage.Eval(game, you, target).(int)
	effects := make([]engine.Effect, damage)
	for i := 0; i < damage; i++ {
		effects[i] = DealDamage(param.Const(1), e.target)
	}
	Sequence(effects...).Happen(game, you, target, cause)
}
