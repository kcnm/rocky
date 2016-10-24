package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type minion struct {
	*character
	card engine.MinionCard
}

func newMinion(id engine.CharacterID, card engine.MinionCard) engine.Minion {
	return &minion{
		newCharacter(
			id,
			card.Attack(),
			card.Health(),
			0, // stamina
		).(*character),
		card,
	}
}

func (m *minion) Card() engine.MinionCard {
	return m.card
}

func (m *minion) String() string {
	return fmt.Sprintf("%d/%d", m.attack, m.health)
}
