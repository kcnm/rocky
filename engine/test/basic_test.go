package test

import (
	"testing"

	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
)

func TestBasicGame(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck(M11, M11, M11, M11))
	p2 := game.NewPlayer(30, Pw2, game.NewDeck(M11, M11, M11, M11))
	g := game.New(p1, p2, nil /* rng */)
	g.Start()

	for _, turn := range []struct {
		name   string
		status GameStatus
	}{
		{
			"Turn 1",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{30, 30, 0, 0, false, 1, 1, 1, 3},
				P2:      PlayerStatus{30, 30, 0, 0, false, 0, 0, 0, 4},
			},
		},
		{
			"Turn 2",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{30, 30, 0, 0, false, 1, 1, 1, 3},
				P2:      PlayerStatus{30, 30, 0, 0, false, 1, 1, 1, 3},
			},
		},
		{
			"Turn 3",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{30, 30, 0, 0, false, 2, 2, 2, 2},
				P2:      PlayerStatus{30, 30, 0, 0, false, 1, 1, 1, 3},
			},
		},
		{
			"Turn 4",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{30, 30, 0, 0, false, 2, 2, 2, 2},
				P2:      PlayerStatus{30, 30, 0, 0, false, 2, 2, 2, 2},
			},
		},
		{
			"Turn 5",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{30, 30, 0, 0, false, 3, 3, 3, 1},
				P2:      PlayerStatus{30, 30, 0, 0, false, 2, 2, 2, 2},
			},
		},
		{
			"Turn 6",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{30, 30, 0, 0, false, 3, 3, 3, 1},
				P2:      PlayerStatus{30, 30, 0, 0, false, 3, 3, 3, 1},
			},
		},
		{
			"Turn 7",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{30, 30, 0, 0, false, 4, 4, 4, 0},
				P2:      PlayerStatus{30, 30, 0, 0, false, 3, 3, 3, 1},
			},
		},
		{
			"Turn 8",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{30, 30, 0, 0, false, 4, 4, 4, 0},
				P2:      PlayerStatus{30, 30, 0, 0, false, 4, 4, 4, 0},
			},
		},
		{
			"Turn 9",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{29, 30, 0, 0, false, 5, 5, 4, 0},
				P2:      PlayerStatus{30, 30, 0, 0, false, 4, 4, 4, 0},
			},
		},
		{
			"Turn 10",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{29, 30, 0, 0, false, 5, 5, 4, 0},
				P2:      PlayerStatus{29, 30, 0, 0, false, 5, 5, 4, 0},
			},
		},
		{
			"Turn 11",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{27, 30, 0, 0, false, 6, 6, 4, 0},
				P2:      PlayerStatus{29, 30, 0, 0, false, 5, 5, 4, 0},
			},
		},
		{
			"Turn 12",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{27, 30, 0, 0, false, 6, 6, 4, 0},
				P2:      PlayerStatus{27, 30, 0, 0, false, 6, 6, 4, 0},
			},
		},
		{
			"Turn 13",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{24, 30, 0, 0, false, 7, 7, 4, 0},
				P2:      PlayerStatus{27, 30, 0, 0, false, 6, 6, 4, 0},
			},
		},
		{
			"Turn 14",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{24, 30, 0, 0, false, 7, 7, 4, 0},
				P2:      PlayerStatus{24, 30, 0, 0, false, 7, 7, 4, 0},
			},
		},
		{
			"Turn 15",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{20, 30, 0, 0, false, 8, 8, 4, 0},
				P2:      PlayerStatus{24, 30, 0, 0, false, 7, 7, 4, 0},
			},
		},
		{
			"Turn 16",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{20, 30, 0, 0, false, 8, 8, 4, 0},
				P2:      PlayerStatus{20, 30, 0, 0, false, 8, 8, 4, 0},
			},
		},
		{
			"Turn 17",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{15, 30, 0, 0, false, 9, 9, 4, 0},
				P2:      PlayerStatus{20, 30, 0, 0, false, 8, 8, 4, 0},
			},
		},
		{
			"Turn 18",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{15, 30, 0, 0, false, 9, 9, 4, 0},
				P2:      PlayerStatus{15, 30, 0, 0, false, 9, 9, 4, 0},
			},
		},
		{
			"Turn 19",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{9, 30, 0, 0, false, 10, 10, 4, 0},
				P2:      PlayerStatus{15, 30, 0, 0, false, 9, 9, 4, 0},
			},
		},
		{
			"Turn 20",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{9, 30, 0, 0, false, 10, 10, 4, 0},
				P2:      PlayerStatus{9, 30, 0, 0, false, 10, 10, 4, 0},
			},
		},
		{
			"Turn 21",
			GameStatus{
				Current: p1,
				P1:      PlayerStatus{2, 30, 0, 0, false, 10, 10, 4, 0},
				P2:      PlayerStatus{9, 30, 0, 0, false, 10, 10, 4, 0},
			},
		},
		{
			"Turn 22",
			GameStatus{
				Current: p2,
				P1:      PlayerStatus{2, 30, 0, 0, false, 10, 10, 4, 0},
				P2:      PlayerStatus{2, 30, 0, 0, false, 10, 10, 4, 0},
			},
		},
		{
			"Turn 23",
			GameStatus{
				Current: p1,
				Over:    true,
				Winner:  p2,
				P1:      PlayerStatus{-6, 30, 0, 0, false, 10, 10, 4, 0},
				P2:      PlayerStatus{2, 30, 0, 0, false, 10, 10, 4, 0},
			},
		},
	} {
		t.Run(turn.name, func(t *testing.T) {
			t.Logf("Turn %d", g.Turn())
			AssertGameStatus(t, g, turn.status)
			p := g.CurrentPlayer()
			if over, _ := g.IsOver(); !over {
				if ok, _ := action.CanEndTurn(g, g.Opponent(p)); ok {
					t.Errorf("Opponent can end turn, expected not")
				}
				action.EndTurn(g, p)
			} else {
				if ok, _ := action.CanEndTurn(g, p); ok {
					t.Errorf("Player can end turn after game over, expected not")
				}
				if ok, _ := action.CanEndTurn(g, g.Opponent(p)); ok {
					t.Errorf("Opponent can end turn after game over, expected not")
				}
			}
		})
	}
}

func TestPlayMinion(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck(), M45, M11, M45)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)
	p1status := PlayerStatus{30, 30, 0, 0, false, 10, 10, 3, 0}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 plays the first 4/5 minion")
	action.PlayCard(g, p1, 0, 0, nil)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 2 // fatigue
	p1status.Mana = p1status.Crystal
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, true})

	t.Logf("P1 plays the second 4/5 minion at position 0")
	action.PlayCard(g, p1, 1, 0, nil)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, 5, true})

	t.Logf("P1 plays the 1/1 minion at position 2")
	action.PlayCard(g, p1, 0, 2, nil)
	p1status.Mana -= 1
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, 5, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M11, 1, 1, 1, false})
}

func TestCastSpell(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck(), S4)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)
	p1status := PlayerStatus{30, 30, 0, 0, false, 10, 10, 1, 0}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 casts spell at P2")
	action.PlayCard(g, p1, 0, 0, p2)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
}

func TestEquipWeapon(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck(), W32)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)
	p1status := PlayerStatus{30, 30, 0, 0, false, 10, 10, 1, 0}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 equips weapon")
	action.PlayCard(g, p1, 0, 0, nil)
	p1status.Attack += 3
	p1status.Active = true
	p1status.Mana -= 2
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, WeaponStatus{W32, 3, 2})
}

func TestMinionAttack(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)
	g.Equip(p2, W32)
	g.Summon(p1, M11, 0)
	g.Summon(p1, M11, 1)
	g.Summon(p1, M45, 2)
	g.Summon(p1, M45, 3)
	g.Summon(p2, M11, 0)
	g.Summon(p2, M45, 1)
	p1status := PlayerStatus{30, 30, 0, 0, false, 10, 10, 0, 0}
	p2status := PlayerStatus{30, 30, 0, 3, false, 0, 0, 0, 0}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M11, 1, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M11, 1, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M45, 4, 5, 5, true})
	AssertMinionStatus(t, p1.Board().Get(3), MinionStatus{M45, 4, 5, 5, true})
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 5, 5, false})

	t.Logf("P1's left most 1/1 minion killed itself onto P2's 4/5")
	action.Attack(g, p1.Board().Get(0), p2.Board().Get(1))
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M11, 1, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, 5, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M45, 4, 5, 5, true})
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 4, 5, false})

	t.Logf("Both left most 1/1 minions destroyed each other")
	action.Attack(g, p1.Board().Get(0), p2.Board().Get(0))
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, 5, true})
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 4, 5, false})

	t.Logf("P1's right 4/5 minion kills P2's 4/4")
	action.Attack(g, p1.Board().Get(1), p2.Board().Get(0))
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 1, 5, false})
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})

	t.Logf("P1's left 4/5 minion attacks P2")
	action.Attack(g, p1.Board().Get(0), p2)
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 1, 5, false})
	p2status.Health -= 4
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
}

func TestPlayerAttack(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	p2.GainCrystal(10)
	g := game.New(p1, p2, nil /* rng */)
	g.Equip(p1, W33)
	g.Equip(p2, W32)
	g.Summon(p2, M11, 0)
	g.Summon(p2, M45, 1)
	p1status := PlayerStatus{30, 30, 0, 3, true, 10, 10, 0, 0}
	p2status := PlayerStatus{30, 30, 0, 3, false, 0, 10, 0, 0}
	w1status := WeaponStatus{W33, 3, 3}
	w2status := WeaponStatus{W32, 3, 2}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 5, 5, false})

	t.Logf("P1 kills P2's 1/1 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.Health -= 1
	p1status.Active = false
	w1status.Durability -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 5, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)
	p2status.Health -= 1 // fatigue
	p2status.Active = true
	p2status.Mana = p2status.Crystal
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 5, 5, true})

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 2 // fatigue
	p1status.Active = true
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 5, 5, true})

	t.Logf("P1 attacks P2's 4/5 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.Health -= 4
	p1status.Active = false
	w1status.Durability -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, 5, true})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)
	p2status.Health -= 2 // fatigue
	p2status.Active = true
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, 5, true})

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 3 // fatigue
	p1status.Active = true
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, 5, true})

	t.Logf("P1 attacks P2")
	action.Attack(g, p1, p2)
	p1status.Attack = 0
	p1status.Active = false
	AssertPlayerStatus(t, p1, p1status)
	if p1.Weapon() != nil {
		t.Errorf("Player1 still equips a weapon, expected not")
	}
	p2status.Health -= 3
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, w2status)
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, 5, true})
}

func TestHeroPower(t *testing.T) {
	p1 := game.NewPlayer(30, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(30, Pw2, game.NewDeck())
	g := game.New(p1, p2, nil /* rng */)
	p1status := PlayerStatus{30, 30, 0, 0, false, 10, 10, 0, 0}

	t.Logf("Start game")
	g.Start()
	p1status.Health -= 1 // fatigue
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 uses hero power")
	action.HeroPower(g, p1, p2)
	p1status.Mana -= 2
	AssertPlayerStatus(t, p1, p1status)
}
