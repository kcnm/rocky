package engine

type Power interface {
	Class() Class
	Mana() int
	Effect() Effect
}
