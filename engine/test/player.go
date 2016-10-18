package test

import (
	"testing"

	"github.com/kcnm/rocky/engine/base"
)

type playerStatus struct {
	health   int
	armor    int
	attack   int
	stamina  int
	active   bool
	mana     int
	crystal  int
	handSize int
	deckSize int
}

func assertPlayerStatus(
	t *testing.T,
	player base.Player,
	status playerStatus) {
	if player.Health() != status.health {
		t.Errorf("Player%d has health %d, expected %d",
			player.ID(), player.Health(), status.health)
	}
	if player.Armor() != status.armor {
		t.Errorf("Player%d has armor %d, expected %d",
			player.ID(), player.Armor(), status.armor)
	}
	if player.Attack() != status.attack {
		t.Errorf("Player%d has attack %d, expected %d",
			player.ID(), player.Attack(), status.attack)
	}
	if player.Stamina() != status.stamina {
		t.Errorf("Player%d has stamina %d, expected %d",
			player.ID(), player.Stamina(), status.stamina)
	}
	if player.Active() != status.active {
		if status.active {
			t.Errorf("Player%d is inactive, expected active", player.ID())
		} else {
			t.Errorf("Player%d is active, expected inactive", player.ID())
		}
	}
	if player.Mana() != status.mana {
		t.Errorf("Player%d has mana %d, expected %d",
			player.ID(), player.Mana(), status.mana)
	}
	if player.Crystal() != status.crystal {
		t.Errorf("Player%d has crystal %d, expected %d",
			player.ID(), player.Crystal(), status.crystal)
	}
	if len(player.Hand()) != status.handSize {
		t.Errorf("Player%d has %d cards in hand, expected %d",
			player.ID(), len(player.Hand()), status.handSize)
	}
	if player.Deck().Remain() != status.deckSize {
		t.Errorf("Player%d has %d cards in deck, expected %d",
			player.ID(), player.Deck().Remain(), status.deckSize)
	}
}
