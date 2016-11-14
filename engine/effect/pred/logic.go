package pred

import (
	"github.com/kcnm/rocky/engine"
)

type and struct {
	preds []Pred
}

func And(preds ...Pred) Pred {
	return &and{preds}
}

func (p *and) Eval(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool {
	for _, pred := range p.preds {
		if !pred.Eval(game, you, sub) {
			return false
		}
	}
	return true
}

func (p *and) BindIt(it interface{}) {
	for _, pred := range p.preds {
		pred.BindIt(it)
	}
}

type or struct {
	preds []Pred
}

func Or(preds ...Pred) Pred {
	return &or{preds}
}

func (p *or) Eval(
	game engine.Game,
	you engine.Player,
	sub interface{}) bool {
	for _, pred := range p.preds {
		if pred.Eval(game, you, sub) {
			return true
		}
	}
	return false
}

func (p *or) BindIt(it interface{}) {
	for _, pred := range p.preds {
		pred.BindIt(it)
	}
}
