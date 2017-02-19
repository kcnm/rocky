package card

import (
	"github.com/kcnm/rocky/engine"
)

const (
	arcaneMissiles    = "Arcane Missiles"
	chillwindYeti     = "Chillwind Yeti"
	fieryWarAxe       = "Fiery War Axe"
	fireball          = "Fireball"
	flamestrike       = "Flamestrike"
	leperGnome        = "Leper Gnome"
	lightningStorm    = "Lightning Storm"
	silverHandRecruit = "Silver Hand Recruit"
)

func ArcaneMissiles() engine.SpellCard     { return *spells[arcaneMissiles] }
func ChillwindYeti() engine.MinionCard     { return *minions[chillwindYeti] }
func FieryWarAxe() engine.WeaponCard       { return *weapons[fieryWarAxe] }
func Fireball() engine.SpellCard           { return *spells[fireball] }
func Flamestrike() engine.SpellCard        { return *spells[flamestrike] }
func LeperGnome() engine.MinionCard        { return *minions[leperGnome] }
func LightningStorm() engine.SpellCard     { return *spells[lightningStorm] }
func SilverHandRecruit() engine.MinionCard { return *minions[silverHandRecruit] }
