package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
)

type PlayerStatus struct {
	Health    int
	MaxHealth int
	Armor     int
	Attack    int
	Active    bool
	Mana      int
	Crystal   int
	HandSize  int
	DeckSize  int
	BoardSize int
}

func AssertPlayerStatus(
	t *testing.T,
	player engine.Player,
	status PlayerStatus) {
	if player.Health() != status.Health {
		t.Errorf("Player%v has health %d, expected %d",
			player.ID(), player.Health(), status.Health)
	}
	if player.MaxHealth() != status.MaxHealth {
		t.Errorf("Player%v has max health %d, expected %d",
			player.ID(), player.MaxHealth(), status.MaxHealth)
	}
	if player.Armor() != status.Armor {
		t.Errorf("Player%v has armor %d, expected %d",
			player.ID(), player.Armor(), status.Armor)
	}
	if player.Attack() != status.Attack {
		t.Errorf("Player%v has attack %d, expected %d",
			player.ID(), player.Attack(), status.Attack)
	}
	if player.Active() != status.Active {
		if status.Active {
			t.Errorf("Player%v is inactive, expected active", player.ID())
		} else {
			t.Errorf("Player%v is active, expected inactive", player.ID())
		}
	}
	if player.Mana() != status.Mana {
		t.Errorf("Player%v has mana %d, expected %d",
			player.ID(), player.Mana(), status.Mana)
	}
	if player.Crystal() != status.Crystal {
		t.Errorf("Player%v has crystal %d, expected %d",
			player.ID(), player.Crystal(), status.Crystal)
	}
	if len(player.Hand()) != status.HandSize {
		t.Errorf("Player%v has %d cards in hand, expected %d",
			player.ID(), len(player.Hand()), status.HandSize)
	}
	if player.Deck().Remain() != status.DeckSize {
		t.Errorf("Player%v has %d cards in deck, expected %d",
			player.ID(), player.Deck().Remain(), status.DeckSize)
	}
	if len(player.Board().Minions()) != status.BoardSize {
		t.Errorf("Player%v has %d minions on board, expected %d",
			player.ID(), len(player.Board().Minions()), status.BoardSize)
	}
}

type MinionStatus struct {
	Card      engine.MinionCard
	Attack    int
	Health    int
	MaxHealth int
	Active    bool
}

func AssertMinionStatus(
	t *testing.T,
	minion engine.Minion,
	status MinionStatus) {
	if minion.Card() != status.Card {
		t.Errorf("Minion%d is '%v', expected '%v'",
			minion.ID(), minion.Card(), status.Card)
	}
	if minion.Health() != status.Health {
		t.Errorf("Minion%d has health %d, expected %d",
			minion.ID(), minion.Health(), status.Health)
	}
	if minion.MaxHealth() != status.MaxHealth {
		t.Errorf("Minion%d has max health %d, expected %d",
			minion.ID(), minion.MaxHealth(), status.MaxHealth)
	}
	if minion.Attack() != status.Attack {
		t.Errorf("Minion%d has attack %d, expected %d",
			minion.ID(), minion.Attack(), status.Attack)
	}
	if minion.Active() != status.Active {
		if status.Active {
			t.Errorf("Minion%d is inactive, expected active", minion.ID())
		} else {
			t.Errorf("Minion%d is active, expected inactive", minion.ID())
		}
	}
}

type WeaponStatus struct {
	Card       engine.WeaponCard
	Attack     int
	Durability int
}

func AssertWeaponStatus(
	t *testing.T,
	player engine.Player,
	status WeaponStatus) {
	w := player.Weapon()
	if w == nil {
		t.Errorf("Player%v does not equip a weapon, expected so", player.ID())
	}
	if w.Card() != status.Card {
		t.Errorf("Weapon of player%v is '%v', expected '%v'",
			player.ID(), w.Card(), status.Card)
	}
	if w.Attack() != status.Attack {
		t.Errorf("Weapon of player%v has attack %d, expected %d",
			player.ID(), w.Attack(), status.Attack)
	}
	if w.Durability() != status.Durability {
		t.Errorf("Weapon of player%v has durability %d, expected %d",
			player.ID(), w.Durability(), status.Durability)
	}
}
