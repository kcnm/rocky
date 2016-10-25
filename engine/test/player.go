package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
)

type playerStatus struct {
	health    int
	armor     int
	attack    int
	active    bool
	mana      int
	crystal   int
	handSize  int
	deckSize  int
	boardSize int
}

func assertPlayerStatus(
	t *testing.T,
	player engine.Player,
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
	if len(player.Board().Minions()) != status.boardSize {
		t.Errorf("Player%d has %d minions on board, expected %d",
			player.ID(), len(player.Board().Minions()), status.boardSize)
	}
}

type minionStatus struct {
	card   engine.Card
	attack int
	health int
	active bool
}

func assertMinionStatus(
	t *testing.T,
	minion engine.Minion,
	status minionStatus) {
	if minion.Card() != status.card {
		t.Errorf("Minion of ID %d is '%v', expected '%v'",
			minion.ID(), minion.Card(), status.card)
	}
	if minion.Health() != status.health {
		t.Errorf("Minion '%v'(%d) has health %d, expected %d",
			minion.Card(), minion.ID(), minion.Health(), status.health)
	}
	if minion.Attack() != status.attack {
		t.Errorf("Minion '%v'(%d) has attack %d, expected %d",
			minion.Card(), minion.ID(), minion.Attack(), status.attack)
	}
	if minion.Active() != status.active {
		if status.active {
			t.Errorf("Minion '%v'(%d) is inactive, expected active",
				minion.Card(), minion.ID())
		} else {
			t.Errorf("Minion '%v'(%d) is active, expected inactive",
				minion.Card(), minion.ID())
		}
	}
}

type weaponStatus struct {
	card       engine.Card
	attack     int
	durability int
}

func assertWeaponStatus(
	t *testing.T,
	player engine.Player,
	status weaponStatus) {
	w := player.Weapon()
	if w == nil {
		t.Errorf("Player%d does not equip a weapon, expected so", player.ID())
	}
	if w.Card() != status.card {
		t.Errorf("Weapon of player%d is '%v', expected '%v'",
			player.ID(), w.Card(), status.card)
	}
	if w.Attack() != status.attack {
		t.Errorf("Weapon of player%d has attack %d, expected %d",
			player.ID(), w.Attack(), status.attack)
	}
	if w.Durability() != status.durability {
		t.Errorf("Weapon of player%d has durability %d, expected %d",
			player.ID(), w.Durability(), status.durability)
	}
}
