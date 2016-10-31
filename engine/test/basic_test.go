package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/game"
)

func TestBasicGame(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck(M11, M11, M11, M11))
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck(M11, M11, M11, M11))
	g := game.New(p1, p2, nil /* rng */)

	for _, turn := range []struct {
		current engine.Player
		over    bool
		winner  engine.Player
		players []PlayerStatus
	}{
		{
			// turn 1
			p1, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 0, 0, 0, 4, 0},
			},
		},
		{
			// turn 2
			p2, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 3
			p1, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 4
			p2, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 5
			p1, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 6
			p2, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 7
			p1, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 8
			p2, false, nil,
			[]PlayerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 9
			p1, false, nil,
			[]PlayerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 10
			p2, false, nil,
			[]PlayerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 11
			p1, false, nil,
			[]PlayerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 12
			p2, false, nil,
			[]PlayerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 13
			p1, false, nil,
			[]PlayerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 14
			p2, false, nil,
			[]PlayerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 15
			p1, false, nil,
			[]PlayerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 16
			p2, false, nil,
			[]PlayerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 17
			p1, false, nil,
			[]PlayerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 18
			p2, false, nil,
			[]PlayerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 19
			p1, false, nil,
			[]PlayerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 20
			p2, false, nil,
			[]PlayerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 21
			p1, false, nil,
			[]PlayerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 22
			p2, false, nil,
			[]PlayerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 23
			p1, true, p2,
			[]PlayerStatus{
				{-6, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
	} {
		t.Logf("Turn %d", g.Turn())
		AssertGameStatus(t, g, turn.current, turn.over, turn.winner)
		AssertPlayerStatus(t, p1, turn.players[0])
		AssertPlayerStatus(t, p2, turn.players[1])
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
	}
}

func TestPlayMinion(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck(), M45, M11, M45)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 0, false, 10, 10, 3, 0, 0}
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 plays the first 4/5 minion")
	action.PlayCard(g, p1, 0, 0, nil)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	p1status.BoardSize += 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 1 // fatigue
	p1status.Mana = p1status.Crystal
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, true})

	t.Logf("P1 plays the second 4/5 minion at position 0")
	action.PlayCard(g, p1, 1, 0, nil)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	p1status.BoardSize += 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, true})

	t.Logf("P1 plays the 1/1 minion at position 2")
	action.PlayCard(g, p1, 0, 2, nil)
	p1status.Mana -= 1
	p1status.HandSize -= 1
	p1status.BoardSize += 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M11, 1, 1, false})
}

func TestCastSpell(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck(), S4)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 0, false, 10, 10, 1, 0, 0}
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 casts spell at P2")
	action.PlayCard(g, p1, 0, 0, p2)
	p1status.Mana -= 4
	p1status.HandSize -= 1
	AssertPlayerStatus(t, p1, p1status)
}

func TestEquipWeapon(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck(), W32)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 0, false, 10, 10, 1, 0, 0}
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
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	p2.Equip(W32)
	g := game.Resume(p1, p2, 1, nil /* rng */)
	g.Summon(M11, p1.Board(), 0)
	g.Summon(M11, p1.Board(), 1)
	g.Summon(M45, p1.Board(), 2)
	g.Summon(M45, p1.Board(), 3)
	g.Summon(M11, p2.Board(), 0)
	g.Summon(M45, p2.Board(), 1)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 0, false, 10, 10, 0, 0, 4}
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M11, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M11, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M45, 4, 5, true})
	AssertMinionStatus(t, p1.Board().Get(3), MinionStatus{M45, 4, 5, true})
	p2status := PlayerStatus{30, 0, 3, false, 0, 0, 0, 0, 2}
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 5, false})

	t.Logf("P1's left most 1/1 minion killed itself onto P2's 4/5")
	action.Attack(g, p1.Board().Get(1), p2.Board().Get(1))
	p1status.BoardSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M11, 1, 1, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, true})
	AssertMinionStatus(t, p1.Board().Get(2), MinionStatus{M45, 4, 5, true})
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 4, false})

	t.Logf("Both left most 1/1 minions destroyed each other")
	action.Attack(g, p1.Board().Get(0), p2.Board().Get(0))
	p1status.BoardSize -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 5, true})
	p2status.BoardSize -= 1
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 4, false})

	t.Logf("P1's right 4/5 minion kills P2's 4/4")
	action.Attack(g, p1.Board().Get(1), p2.Board().Get(0))
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, true})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 1, false})
	p2status.BoardSize -= 1
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})

	t.Logf("P1's left 4/5 minion attacks P2")
	action.Attack(g, p1.Board().Get(0), p2)
	AssertPlayerStatus(t, p1, p1status)
	AssertMinionStatus(t, p1.Board().Get(0), MinionStatus{M45, 4, 5, false})
	AssertMinionStatus(t, p1.Board().Get(1), MinionStatus{M45, 4, 1, false})
	p2status.Health -= 4
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
}

func TestPlayerAttack(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p1.Equip(W33)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	p2.GainCrystal(10)
	p2.Equip(W32)
	g := game.Resume(p1, p2, 1, nil /* rng */)
	g.Summon(M11, p2.Board(), 0)
	g.Summon(M45, p2.Board(), 1)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 3, true, 10, 10, 0, 0, 0}
	w1status := WeaponStatus{W33, 3, 3}
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	p2status := PlayerStatus{30, 0, 3, false, 0, 10, 0, 0, 2}
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M11, 1, 1, false})
	AssertMinionStatus(t, p2.Board().Get(1), MinionStatus{M45, 4, 5, false})

	t.Logf("P1 kills P2's 1/1 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.Health -= 1
	p1status.Active = false
	w1status.Durability -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	p2status.BoardSize -= 1
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 1 // fatigue
	p1status.Active = true
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	p2status.Health -= 1 // fatigue
	p2status.Active = true
	p2status.Mana = p2status.Crystal
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 5, true})

	t.Logf("P1 attacks P2's 4/5 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.Health -= 4
	p1status.Active = false
	w1status.Durability -= 1
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, true})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.Health -= 2 // fatigue
	p1status.Active = true
	AssertPlayerStatus(t, p1, p1status)
	AssertWeaponStatus(t, p1, w1status)
	p2status.Health -= 2 // fatigue
	p2status.Active = true
	AssertPlayerStatus(t, p2, p2status)
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, true})

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
	AssertWeaponStatus(t, p2, WeaponStatus{W32, 3, 2})
	AssertMinionStatus(t, p2.Board().Get(0), MinionStatus{M45, 4, 2, true})
}

func TestHeroPower(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, Pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, Pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := PlayerStatus{30, 0, 0, false, 10, 10, 0, 0, 0}
	AssertPlayerStatus(t, p1, p1status)

	t.Logf("P1 uses hero power")
	action.HeroPower(g, p1, p2)
	p1status.Mana -= 2
	AssertPlayerStatus(t, p1, p1status)
}
