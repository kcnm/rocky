package effect

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/param"
	"github.com/kcnm/rocky/engine/event"
)

type dealDamage struct {
	dmg    int
	target param.Param
}

func DealDamage(dmg int, target param.Param) engine.Effect {
	return &dealDamage{dmg, target}
}

func (e dealDamage) CanHappen(
	game engine.Game,
	you engine.Player,
	target engine.Char) bool {
	return true
}

func (e dealDamage) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	if target == nil {
		panic("nil target")
	}
	game.Events().Post(
		event.Damage(game, target, nil, e.dmg), cause)
}
