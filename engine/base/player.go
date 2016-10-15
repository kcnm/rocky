package base

type Player interface {
	Character

	Mana() int
	Crystal() int
	HasMaxCrystal() bool
	Board() Board
	Deck() Deck
	Hand() []Card
	HandIsFull() bool
	IsControlling(char Character) bool

	GainMana(mana int)
	GainCrystal(crystal int)
	Take(card Card) bool
	Play(cardIndex int) Card
}
