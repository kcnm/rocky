package buff

import (
	"github.com/kcnm/rocky/engine"
)

type none int

var None none = 0

func (b none) Apply(
	game engine.Game,
	you engine.Player,
	char engine.Char) {
}
