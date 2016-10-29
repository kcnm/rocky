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
	return e.target.Eval(game, you, target) != nil
}

func (e dealDamage) Happen(
	game engine.Game,
	you engine.Player,
	target engine.Char,
	cause engine.Event) {
	t := e.target.Eval(game, you, target)
	switch t := t.(type) {
	case engine.Char:
		game.Events().Post(
			event.Damage(game, t, nil, e.dmg), cause)
	case []engine.Char:
		events := make([]engine.Event, len(t))
		for i, ch := range t {
			events[i] = event.Damage(game, ch, nil, e.dmg)
		}
		game.Events().Post(
			event.Combined(events...), cause)
	default:
		panic("invalid evaluated target")
	}
}
