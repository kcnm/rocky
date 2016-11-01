package effect

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/param"
	"github.com/kcnm/rocky/engine/event"
)

type dealDamage struct {
	damage param.Param
	target param.Param
}

func DealDamage(damage param.Param, target param.Param) engine.Effect {
	return &dealDamage{damage, target}
}

func (e dealDamage) CanHappen(
	game engine.Game,
	you engine.Player,
	target engine.Char) bool {
	dmg := e.damage.Eval(game, you, target)
	t := e.target.Eval(game, you, target)
	return dmg.(int) > 0 && t != nil
}

func (e dealDamage) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	dmg := e.damage.Eval(game, you, target).(int)
	t := e.target.Eval(game, you, target)
	switch t := t.(type) {
	case engine.Char:
		game.Events().Post(
			event.Damage(game, t, nil, dmg), cause)
	case []engine.Char:
		events := make([]engine.Event, len(t))
		for i, ch := range t {
			events[i] = event.Damage(game, ch, nil, dmg)
		}
		game.Events().Post(
			event.Combined(events...), cause)
	default:
		panic("invalid evaluated target")
	}
}
