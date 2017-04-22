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

type Card interface {
	Class() Class
	Mana() int
}

type MinionCard interface {
	Card

	Attack() int
	Health() int
	Battlecry() Effect
	Buff() Buff
}

type SpellCard interface {
	Card

	Effect() Effect
}

type WeaponCard interface {
	Card

	Attack() int
	Durability() int
	Battlecry() Effect
}

type Effect interface {
	CanHappen(game Game, you Player, target Char) bool
	Happen(game Game, you Player, target Char, cause Event)
}

type Buff interface {
	Apply(game Game, you Player, char Char)
}

type Ability interface {
	Class() Class
	Mana() int
	Effect() Effect
}
