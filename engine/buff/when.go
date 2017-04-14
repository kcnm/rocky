package buff

import (
	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/effect/pred"
)

type when struct {
	evPred pred.Pred
	effect engine.Effect
}

func When(evPred pred.Pred, effect engine.Effect) engine.Buff {
	return &when{evPred, effect}
}

func (b *when) Apply(
	game engine.Game,
	you engine.Player,
	char engine.Char) {
	char.AppendReactor(func(ev engine.Event) {
		if b.evPred.BindIt(char).Eval(game, you, ev) {
			b.effect.Happen(game, you, nil, ev)
		}
	})
}

func Deathrattle(effect engine.Effect) engine.Buff {
	return When(pred.Destroy(pred.It()), effect)
}
