package engine

import (
	"math/rand"
)

// Entity is an in-game element that participate in event system. Most objects
// can be modeled as Entity to interact with various Events, abstract or
// concrete, such as players, minions and game itself.
type Entity interface {
	// ID returns its EntityID.
	ID() EntityID
	// React performs certain actions on given Event ev. Implementations can have
	// its fixed behavior, while support appending Reactors to modify this.
	React(ev Event)
	// AppendReactor appends Reactor r to the Entity, which is called at React by
	// append order.
	AppendReactor(r Reactor)
	// Reset removes all previously appended Reactors.
	Reset()
}

// EntityID is the ID of an Entity.
type EntityID int

// Reactor is a closure to perform actions on given Event.
type Reactor func(Event)

// Char is a character, either player or minion.
type Char interface {
	Entity

	// Attack returns the current attack.
	Attack() int
	// Health returns the current health.
	Health() int
	// MaxHealth returns the maximum health.
	MaxHealth() int
	// Swings returns the number of swings/attacks performed this turn.
	Swings() int
	// Active returns whether the character can attack now.
	Active() bool

	// Refresh resets character's swings at the start of a turn.
	Refresh()
	// TakeDamage damages character by dmg, returns the actual damage taken and
	// whether it is a fatal damage (resulting in non-positive health).
	TakeDamage(dmg int) (actual int, fatal bool)
	// Swing performs an attack, increase and returns its counter this turn.
	Swing() int
}

// Player is the player's character, or hero.
type Player interface {
	Char

	// Armor returns the current armor.
	Armor() int
	// Mana returns the current available mana.
	Mana() int
	// Crystal returns the current number of crystals.
	Crystal() int
	// HasMaxCrystal returns whether the player has maximum number of crystals.
	HasMaxCrystal() bool
	// Ability returns player's ability, aka hero power.
	Ability() Ability
	// CanHeroPower returns whether the player can use his ability/hero power now.
	CanHeroPower() bool
	// Weapon returns the weapon currently equiped by the player, or nil.
	Weapon() Weapon
	// Board returns the player's board.
	Board() Board
	// Deck returns the player's deck.
	Deck() Deck
	// Hand returns the player's hand of cards.
	Hand() []Card
	// HandIsFull returns whether the player has a full hand.
	HandIsFull() bool
	// IsControlling returns whether the player is controlling given character:
	// true if it is either the player himself, or a minion on his board.
	IsControlling(char Char) bool

	// GainArmor gains given amount of armor (negative for loss), and returns the
	// armor afterwards.
	GainArmor(armor int) int
	// GainMana gains given amount of mana (negative for loss), and returns the
	// mana afterwards.
	GainMana(mana int) int
	// GainCrystal gains given amount of EMPTY crystals (negative for loss), and
	// returns the crystals afterwards. Player's mana will be capped if necessary
	// at losing crystals.
	GainCrystal(crystal int) int
	// Take adds given card to player's hand if not full, and returns whether he
	// has taken it: false when his hand is already full and the card is burnt.
	Take(card Card) bool
	// Play plays out and returns the card at given index in player's hand.
	Play(cardIndex int) Card
	// HeroPower uses player's ability, and returns its effect.
	HeroPower() Effect
	// Equip equips given weapon for player, and returns the weapon. It requires
	// that the player does not have any weapon currently equiped.
	Equip(weapon Weapon) Weapon
}

// Board is a player's board where minions live on.
type Board interface {
	// Minions returns a slice of minions on the board.
	Minions() []Minion
	// IsFull returns whether the board is full: no other minion can be placed.
	IsFull() bool
	// Find returns the index of given minion, starting from 0 as viewed from left
	// of the player. Returns -1 if given minion is not on the board.
	Find(minion Minion) (index int)
	// Get returns the minion at given position index on the board.
	Get(pos int) Minion

	// Put puts given minion at position on the board, and returns it. It requires
	// that the board is not full.
	Put(minion Minion, position int) Minion
	// Remove removes given minion from the board, and returns it. Returns nil if
	// given minion is not on the board.
	Remove(minion Minion) Minion
}

// Deck is a collection of cards drawn from by player.
type Deck interface {
	// Remain returns the number of cards remained in the deck.
	Remain() int

	// PutOnTop puts given card on top of the deck.
	PutOnTop(card Card)
	// Shuffle shuffles the deck with given rng.
	Shuffle(rng *rand.Rand)
	// Draw draws a card from the deck, and returns it. If the deck is empty, the
	// card is returned as nil, together with a fatigue damage.
	Draw() (card Card, fatigue int)
}

// Weapon is equipable by player so that they can attack with it. Each swing
// usually loses 1 durability. The weapon is destroyed when its durability
// reaches 0.
type Weapon interface {
	Entity

	// Card returns the underlying card of the weapon.
	Card() WeaponCard
	// Attack returns its attack.
	Attack() int
	// Durability returns its durability.
	Durability() int

	// LoseDurability loses d durability of the weapon, and returns the durability
	// afterwards. It requires the current durability to be no less than d.
	LoseDurability(d int) int
}

// Minion is a creature/character lives on player's board.
type Minion interface {
	Char

	// Card returns the underlying card of the minion.
	Card() MinionCard
}
