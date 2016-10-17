package effect

import (
	"github.com/kcnm/rocky/engine/base"
	"github.com/kcnm/rocky/engine/event"
)

type dealDamage struct {
	dmg int
}

func DealDamage(dmg int) base.Effect {
	return &dealDamage{dmg}
}

func (e *dealDamage) Happen(
	game base.Game,
	cause base.Event,
	targets []base.Character) {
	if len(targets) == 0 {
		panic("deal damage to 0 targets")
	}
	if len(targets) == 1 {
		game.Events().Post(
			event.Damage(game, targets[0], nil, e.dmg), cause)
	} else {
		events := make([]base.Event, len(targets))
		for i, tgt := range targets {
			events[i] = event.Damage(game, tgt, nil, e.dmg)
		}
		game.Events().Post(
			base.Combined(events...), cause)
	}
}
