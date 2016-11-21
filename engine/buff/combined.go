package buff

import (
	"github.com/kcnm/rocky/engine"
)

type combined struct {
	buffs []engine.Buff
}

func Combined(buffs ...engine.Buff) engine.Buff {
	flattened := make([]engine.Buff, 0, len(buffs))
	for _, b := range buffs {
		flattened = append(flattened, flatten(b)...)
	}
	return &combined{flattened}
}

func (b *combined) Apply(
	game engine.Game,
	you engine.Player,
	char engine.Char) {
	for _, buff := range b.buffs {
		buff.Apply(game, you, char)
	}
}

func flatten(buff engine.Buff) []engine.Buff {
	if b, ok := buff.(*combined); ok {
		flattened := make([]engine.Buff, 0, len(b.buffs))
		for _, b := range b.buffs {
			flattened = append(flattened, flatten(b)...)
		}
		return flattened
	} else {
		return []engine.Buff{buff}
	}
}
