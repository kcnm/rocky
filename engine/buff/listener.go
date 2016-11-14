package buff

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/pred"
)

type when struct {
	evPred pred.Pred
	effect engine.Effect
}

func (b *when) Apply(
	game engine.Game,
	you engine.Player,
	char engine.Char) {
	b.evPred.BindIt(char)
	char.AddHandler(func(ev engine.Event) {
		if b.evPred.Eval(game, you, ev) {
			b.effect.Happen(game, you, nil, ev)
		}
	})
}

func When(evPred pred.Pred, effect engine.Effect) engine.Buff {
	return &when{evPred, effect}
}

// TODO: Change to After.
func Deathrattle(effect engine.Effect) engine.Buff {
	return When(pred.Destroy(pred.It()), effect)
}
