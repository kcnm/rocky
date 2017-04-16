package engine

import (
	"math/rand"
)

type EntityID int

type Entity interface {
	ID() EntityID
	React(ev Event)
	AppendReactor(r Reactor)
	Reset()
}

type Reactor func(Event)

type Char interface {
	Entity

	Attack() int
	Health() int
	MaxHealth() int
	Swings() int
	Active() bool

	Refresh()
	TakeDamage(damage int) (actual int, fatal bool)
	Swing()
}

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
	GainArmor(armor int)
	Take(card Card) bool
	Play(cardIndex int) Card
	HeroPower() Effect
	Equip(card WeaponCard)
	DestroyWeapon()
}

type Board interface {
	Minions() []Minion
	IsFull() bool
	Find(minion Minion) (index int)
	Get(pos int) Minion

	Put(minion Minion, position int) Minion
}

type Deck interface {
	Remain() int

	PutOnTop(card Card)
	Shuffle(rng *rand.Rand)
	Draw() (card Card, fatigue int)
}

type Weapon interface {
	Card() WeaponCard
	Attack() int
	Durability() int

	LoseDurability()
}

type Minion interface {
	Char

	Card() MinionCard
}
