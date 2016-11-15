package pred

import (
	"github.com/kcnm/rocky/engine"
)

type event struct {
	subPred Pred
	verb    engine.Verb
}

func (p *event) Eval(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool {
	ev, ok := sub.(engine.Event)
	if !ok {
		return false
	}
	if ev.Verb() != p.verb {
		return false
	}
	if !p.subPred.Eval(game, you, ev.Subject()) {
		return false
	}
	return true
}

func (p *event) BindIt(x interface{}) Pred {
	return &event{
		p.subPred.BindIt(x),
		p.verb,
	}
}

func Destroy(subPred Pred) Pred {
	return &event{subPred, engine.Destroy}
}
