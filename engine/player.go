package engine

type Player interface {
	Char

	Armor() int
	Mana() int
	Crystal() int
	HasMaxCrystal() bool
	Power() Power
	Powered() bool
	Weapon() Weapon
	Board() Board
	Deck() Deck
	Hand() []Card
	HandIsFull() bool
	IsControlling(char Char) bool

	GainMana(mana int)
	GainCrystal(crystal int)
	Take(card Card) bool
	Play(cardIndex int) Card
	HeroPower() []Effect
	Equip(card WeaponCard)
	DestroyWeapon()
}
