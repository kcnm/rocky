package engine

type Class string

const (
	Neutral Class = "Neutral"
	Druid         = "Druid"
	Hunter        = "Hunter"
	Mage          = "Mage"
	Paladin       = "Paladin"
	Priest        = "Priest"
	Rogue         = "Rogue"
	Shaman        = "Shaman"
	Warlock       = "Warlock"
	Warrior       = "Warrior"
)

type Effect interface {
	CanHappen(game Game, you Player, target Char) bool
	Happen(game Game, you Player, target Char, cause Event)
}

type Buff interface {
	Apply(game Game, you Player, char Char)
}

type Power interface {
	Class() Class
	Mana() int
	Effect() Effect
}
