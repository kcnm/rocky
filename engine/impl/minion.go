package impl

import (
	"fmt"

	"github.com/kcnm/rocky/engine/base"
)

type minion struct {
	*character
	card base.MinionCard
}

func newMinion(id base.CharacterID, card base.MinionCard) base.Minion {
	return &minion{
		newCharacter(
			id,
			card.InitialAttack(),
			card.InitialHealth(),
			0, // stamina
		).(*character),
		card,
	}
}

func (m *minion) Card() base.MinionCard {
	return m.card
}

func (m *minion) String() string {
	return fmt.Sprintf("%d/%d", m.attack, m.health)
}
