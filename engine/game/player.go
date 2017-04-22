package game

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kcnm/rocky/engine"
)

var maxCrystal = flag.Int("max_crystal", 10, "maximum number of crystals")
var maxHand = flag.Int("max_hand", 10, "maximum number of cards in hand")

type player struct {
	*char
	armor         int
	mana          int
	crystal       int
	ability       engine.Ability
	heroPowers    int
	maxHeroPowers int
	weapon        engine.Weapon
	hand          []engine.Card
	deck          engine.Deck
	board         engine.Board
}

func NewPlayer(
	maxHealth int,
	ability engine.Ability,
	deck engine.Deck,
	hand ...engine.Card) engine.Player {
	if len(hand) > *maxHand {
		panic("too many cards in hand")
	}
	return &player{
		newChar(
			0,         // id
			0,         // attack
			maxHealth, // health
			maxHealth, // maxHealth
		).(*char),
		0,       // armor
		0,       // mana
		0,       // crystal
		ability, // ability
		0,       // heroPowers
		1,       // maxHeroPowers
		nil,     // weapon
		hand,
		deck,
		newBoard(),
	}
}

func (p *player) Attack() int {
	attack := p.char.Attack()
	if p.weapon != nil {
		attack += p.weapon.Attack()
	}
	return attack
}

func (p *player) Active() bool {
	return p.Attack() > 0 && p.char.swings < p.char.maxSwings
}

func (p *player) Refresh() {
	p.char.Refresh()
	p.mana = p.crystal
	p.heroPowers = 0
	for _, minion := range p.board.Minions() {
		minion.Refresh()
	}
}

func (p *player) TakeDamage(dmg int) (actual int, fatal bool) {
	if dmg <= 0 {
		panic(fmt.Errorf("non-positive damage %d", dmg))
	}
	p.armor -= dmg
	if p.armor < 0 {
		p.health += p.armor
		p.armor = 0
	}
	return dmg, p.health <= 0
}

func (p *player) Armor() int {
	return p.armor
}

func (p *player) Mana() int {
	return p.mana
}

func (p *player) Crystal() int {
	return p.crystal
}

func (p *player) HasMaxCrystal() bool {
	return p.crystal >= *maxCrystal
}

func (p *player) Ability() engine.Ability {
	return p.ability
}

func (p *player) CanHeroPower() bool {
	return p.heroPowers < p.maxHeroPowers
}

func (p *player) Weapon() engine.Weapon {
	return p.weapon
}

func (p *player) Board() engine.Board {
	return p.board
}

func (p *player) Deck() engine.Deck {
	return p.deck
}

func (p *player) Hand() []engine.Card {
	return p.hand
}

func (p *player) HandIsFull() bool {
	return len(p.hand) >= *maxHand
}

func (p *player) IsControlling(char engine.Char) bool {
	if p == char {
		return true
	}
	if minion, ok := char.(engine.Minion); ok {
		return p.board.Find(minion) >= 0
	}
	return false
}

func (p *player) GainArmor(armor int) int {
	p.armor += armor
	return p.armor
}

func (p *player) GainMana(mana int) int {
	if p.mana+mana < 0 {
		panic(fmt.Errorf("cannot lose mana %d from %d", -mana, p.mana))
	}
	p.mana += mana
	if p.mana > p.crystal {
		p.mana = p.crystal
	}
	return p.mana
}

func (p *player) GainCrystal(crystal int) int {
	if p.crystal+crystal < 0 {
		panic(fmt.Errorf("cannot lose crystal %d from %d", -crystal, p.crystal))
	}
	p.crystal += crystal
	if p.crystal > *maxCrystal {
		p.crystal = *maxCrystal
	}
	if p.mana > p.crystal {
		p.mana = p.crystal
	}
	return p.crystal
}

func (p *player) Take(card engine.Card) bool {
	if p.HandIsFull() {
		return false
	}
	p.hand = append(p.hand, card)
	return true
}

func (p *player) Play(cardIndex int) engine.Card {
	card := p.hand[cardIndex]
	p.GainMana(-card.Mana())
	p.hand = append(p.hand[:cardIndex], p.hand[cardIndex+1:]...)
	return card
}

func (p *player) HeroPower() engine.Effect {
	if p.heroPowers >= p.maxHeroPowers {
		panic(fmt.Errorf("cannot hero power: %d/%d", p.heroPowers, p.maxHeroPowers))
	}
	p.GainMana(-p.ability.Mana())
	p.heroPowers++
	return p.ability.Effect()
}

func (p *player) Equip(weapon engine.Weapon) engine.Weapon {
	if p.weapon != nil {
		panic("player has weapon equiped already")
	}
	p.weapon = weapon
	return p.weapon
}

func (p *player) react(ev engine.Event) {
	switch ev.Verb() {
	case engine.Destroy:
		if p.weapon == ev.Subject() {
			p.weapon = nil
		} else if minion, ok := ev.Subject().(engine.Minion); ok {
			p.board.Remove(minion)
		}
	case engine.Combined:
		for _, ev := range ev.Subject().([]engine.Event) {
			p.react(ev)
		}
	}
}

func (p *player) String() string {
	armor := ""
	if p.armor > 0 {
		armor = fmt.Sprintf("+%d", p.armor)
	}
	mana := fmt.Sprintf("Mana: %d/%d", p.mana, p.crystal)
	deck := fmt.Sprintf("Deck: %d", p.deck.Remain())
	weapon := ""
	if p.Weapon() != nil {
		weapon = fmt.Sprintf("Weapon: %v", p.Weapon())
	}
	cards := make([]string, len(p.hand))
	for i, c := range p.hand {
		cards[i] = fmt.Sprintf("%v", c)
	}
	hand := fmt.Sprintf("Hand: %s", strings.Join(cards, ", "))
	return fmt.Sprintf("Player%v(%d%s) %s %s %s\n%s\n%v",
		p.id, p.health, armor, mana, deck, weapon, hand, p.board)
}
