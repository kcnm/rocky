package game

import (
	"fmt"

	"github.com/kcnm/rocky/engine"
)

type minion struct {
	*char
	card engine.MinionCard
}

func newMinion(id engine.EntityID, card engine.MinionCard) engine.Minion {
	return &minion{
		newChar(
			id,
			card.Attack(),
			card.Health(),
			card.Health(),
		).(*char),
		card,
	}
}

func (m *minion) Card() engine.MinionCard {
	return m.card
}

func (m *minion) String() string {
	return fmt.Sprintf("%d/%d", m.attack, m.health)
}
