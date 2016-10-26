package power

import (
	"github.com/kcnm/rocky/engine"
)

type standard int

const (
	Fireblast standard = iota
)

type powerSpec struct {
	name    string
	class   engine.Class
	mana    int
	effects []engine.Effect
}
