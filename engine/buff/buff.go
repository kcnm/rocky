package buff

import (
	"github.com/kcnm/rocky/engine"
)

func Listeners(buff engine.Buff) engine.Buff {
	switch b := buff.(type) {
	case *combined:
		return extract(b, Listeners)
	case *when:
		return b
	default:
		return None
	}
}

type extractFn func(engine.Buff) engine.Buff

func extract(combined *combined, fn extractFn) engine.Buff {
	extracted := make([]engine.Buff, 0)
	for _, buff := range combined.buffs {
		for _, b := range flatten(buff) {
			if x := fn(b); x != None {
				extracted = append(extracted, x)
			}
		}
	}
	if len(extracted) > 0 {
		return Combined(extracted...)
	} else {
		return None
	}
}
