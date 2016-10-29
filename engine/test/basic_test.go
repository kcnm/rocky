package test

import (
	"testing"

	"github.com/kcnm/rocky/engine"
	"github.com/kcnm/rocky/engine/action"
	"github.com/kcnm/rocky/engine/effect"
	"github.com/kcnm/rocky/engine/game"
)

var (
	m11 = NewMinionCard(engine.Neutral, 1, 1, 1)
	m45 = NewMinionCard(engine.Neutral, 4, 4, 5)
	s4  = NewSpellCard(engine.Neutral, 4, effect.None)
	w32 = NewWeaponCard(engine.Neutral, 2, 3, 2)
	w33 = NewWeaponCard(engine.Neutral, 4, 3, 3)
	pw2 = NewPower(engine.Neutral, 2, effect.None)
)

func TestBasicGame(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck(m11, m11, m11, m11))
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck(m11, m11, m11, m11))
	g := game.New(p1, p2, nil /* rng */)

	for _, turn := range []struct {
		current engine.Player
		over    bool
		winner  engine.Player
		players []playerStatus
	}{
		{
			// turn 1
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 0, 0, 0, 4, 0},
			},
		},
		{
			// turn 2
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 1, 1, 1, 3, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 3
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 1, 1, 1, 3, 0},
			},
		},
		{
			// turn 4
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 2, 2, 2, 2, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 5
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 2, 2, 2, 2, 0},
			},
		},
		{
			// turn 6
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 3, 3, 3, 1, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 7
			p1, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 3, 3, 3, 1, 0},
			},
		},
		{
			// turn 8
			p2, false, nil,
			[]playerStatus{
				{30, 0, 0, false, 4, 4, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 9
			p1, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{30, 0, 0, false, 4, 4, 4, 0, 0},
			},
		},
		{
			// turn 10
			p2, false, nil,
			[]playerStatus{
				{29, 0, 0, false, 5, 5, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 11
			p1, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{29, 0, 0, false, 5, 5, 4, 0, 0},
			},
		},
		{
			// turn 12
			p2, false, nil,
			[]playerStatus{
				{27, 0, 0, false, 6, 6, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 13
			p1, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{27, 0, 0, false, 6, 6, 4, 0, 0},
			},
		},
		{
			// turn 14
			p2, false, nil,
			[]playerStatus{
				{24, 0, 0, false, 7, 7, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 15
			p1, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{24, 0, 0, false, 7, 7, 4, 0, 0},
			},
		},
		{
			// turn 16
			p2, false, nil,
			[]playerStatus{
				{20, 0, 0, false, 8, 8, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 17
			p1, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{20, 0, 0, false, 8, 8, 4, 0, 0},
			},
		},
		{
			// turn 18
			p2, false, nil,
			[]playerStatus{
				{15, 0, 0, false, 9, 9, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 19
			p1, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{15, 0, 0, false, 9, 9, 4, 0, 0},
			},
		},
		{
			// turn 20
			p2, false, nil,
			[]playerStatus{
				{9, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 21
			p1, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{9, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 22
			p2, false, nil,
			[]playerStatus{
				{2, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
		{
			// turn 23
			p1, true, p2,
			[]playerStatus{
				{-6, 0, 0, false, 10, 10, 4, 0, 0},
				{2, 0, 0, false, 10, 10, 4, 0, 0},
			},
		},
	} {
		t.Logf("Turn %d", g.Turn())
		assertGameStatus(t, g, turn.current, turn.over, turn.winner)
		assertPlayerStatus(t, p1, turn.players[0])
		assertPlayerStatus(t, p2, turn.players[1])
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
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck(), m45, m11, m45)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 0, false, 10, 10, 3, 0, 0}
	assertPlayerStatus(t, p1, p1status)

	t.Logf("P1 plays the first 4/5 minion")
	action.PlayCard(g, p1, 0, 0, nil)
	p1status.mana -= 4
	p1status.handSize -= 1
	p1status.boardSize += 1
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.health -= 1 // fatigue
	p1status.mana = p1status.crystal
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, true})

	t.Logf("P1 plays the second 4/5 minion at position 0")
	action.PlayCard(g, p1, 1, 0, nil)
	p1status.mana -= 4
	p1status.handSize -= 1
	p1status.boardSize += 1
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})

	t.Logf("P1 plays the 1/1 minion at position 2")
	action.PlayCard(g, p1, 0, 2, nil)
	p1status.mana -= 1
	p1status.handSize -= 1
	p1status.boardSize += 1
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(2), minionStatus{m11, 1, 1, false})
}

func TestCastSpell(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck(), s4)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 0, false, 10, 10, 1, 0, 0}
	assertPlayerStatus(t, p1, p1status)

	t.Logf("P1 casts spell at P2")
	action.PlayCard(g, p1, 0, 0, p2)
	p1status.mana -= 4
	p1status.handSize -= 1
	assertPlayerStatus(t, p1, p1status)
}

func TestEquipWeapon(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck(), w32)
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 0, false, 10, 10, 1, 0, 0}
	assertPlayerStatus(t, p1, p1status)

	t.Logf("P1 equips weapon")
	action.PlayCard(g, p1, 0, 0, nil)
	p1status.attack += 3
	p1status.active = true
	p1status.mana -= 2
	p1status.handSize -= 1
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, weaponStatus{w32, 3, 2})
}

func TestMinionAttack(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	p2.Equip(w32.(engine.WeaponCard))
	g := game.Resume(p1, p2, 1, nil /* rng */)
	g.Summon(m11.(engine.MinionCard), p1.Board(), 0)
	g.Summon(m11.(engine.MinionCard), p1.Board(), 1)
	g.Summon(m45.(engine.MinionCard), p1.Board(), 2)
	g.Summon(m45.(engine.MinionCard), p1.Board(), 3)
	g.Summon(m11.(engine.MinionCard), p2.Board(), 0)
	g.Summon(m45.(engine.MinionCard), p2.Board(), 1)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 0, false, 10, 10, 0, 0, 4}
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m11, 1, 1, true})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m11, 1, 1, true})
	assertMinionStatus(t, p1.Board().Get(2), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(3), minionStatus{m45, 4, 5, true})
	p2status := playerStatus{30, 0, 3, false, 0, 0, 0, 0, 2}
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m11, 1, 1, false})
	assertMinionStatus(t, p2.Board().Get(1), minionStatus{m45, 4, 5, false})

	t.Logf("P1's left most 1/1 minion killed itself onto P2's 4/5")
	action.Attack(g, p1.Board().Get(1), p2.Board().Get(1))
	p1status.boardSize -= 1
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m11, 1, 1, true})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(2), minionStatus{m45, 4, 5, true})
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m11, 1, 1, false})
	assertMinionStatus(t, p2.Board().Get(1), minionStatus{m45, 4, 4, false})

	t.Logf("Both left most 1/1 minions destroyed each other")
	action.Attack(g, p1.Board().Get(0), p2.Board().Get(0))
	p1status.boardSize -= 1
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 5, true})
	p2status.boardSize -= 1
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 4, false})

	t.Logf("P1's right 4/5 minion kills P2's 4/4")
	action.Attack(g, p1.Board().Get(1), p2.Board().Get(0))
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, true})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 1, false})
	p2status.boardSize -= 1
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})

	t.Logf("P1's left 4/5 minion attacks P2")
	action.Attack(g, p1.Board().Get(0), p2)
	assertPlayerStatus(t, p1, p1status)
	assertMinionStatus(t, p1.Board().Get(0), minionStatus{m45, 4, 5, false})
	assertMinionStatus(t, p1.Board().Get(1), minionStatus{m45, 4, 1, false})
	p2status.health -= 4
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
}

func TestPlayerAttack(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck())
	p1.GainCrystal(10)
	p1.Equip(w33.(engine.WeaponCard))
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	p2.GainCrystal(10)
	p2.Equip(w32.(engine.WeaponCard))
	g := game.Resume(p1, p2, 1, nil /* rng */)
	g.Summon(m11.(engine.MinionCard), p2.Board(), 0)
	g.Summon(m45.(engine.MinionCard), p2.Board(), 1)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 3, true, 10, 10, 0, 0, 0}
	w1status := weaponStatus{w33, 3, 3}
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, w1status)
	p2status := playerStatus{30, 0, 3, false, 0, 10, 0, 0, 2}
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m11, 1, 1, false})
	assertMinionStatus(t, p2.Board().Get(1), minionStatus{m45, 4, 5, false})

	t.Logf("P1 kills P2's 1/1 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.health -= 1
	p1status.active = false
	w1status.durability -= 1
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, w1status)
	p2status.boardSize -= 1
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 5, false})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.health -= 1 // fatigue
	p1status.active = true
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, w1status)
	p2status.health -= 1 // fatigue
	p2status.active = true
	p2status.mana = p2status.crystal
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 5, true})

	t.Logf("P1 attacks P2's 4/5 minion")
	action.Attack(g, p1, p2.Board().Get(0))
	p1status.health -= 4
	p1status.active = false
	w1status.durability -= 1
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, w1status)
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 2, true})

	t.Logf("P1 ends turn %d", g.Turn())
	action.EndTurn(g, p1)

	t.Logf("P2 ends turn %d", g.Turn())
	action.EndTurn(g, p2)
	p1status.health -= 2 // fatigue
	p1status.active = true
	assertPlayerStatus(t, p1, p1status)
	assertWeaponStatus(t, p1, w1status)
	p2status.health -= 2 // fatigue
	p2status.active = true
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 2, true})

	t.Logf("P1 attacks P2")
	action.Attack(g, p1, p2)
	p1status.attack = 0
	p1status.active = false
	assertPlayerStatus(t, p1, p1status)
	if p1.Weapon() != nil {
		t.Errorf("Player1 still equips a weapon, expected not")
	}
	p2status.health -= 3
	assertPlayerStatus(t, p2, p2status)
	assertWeaponStatus(t, p2, weaponStatus{w32, 3, 2})
	assertMinionStatus(t, p2.Board().Get(0), minionStatus{m45, 4, 2, true})
}

func TestHeroPower(t *testing.T) {
	p1 := game.NewPlayer(1, 30, 0, pw2, game.NewDeck())
	p1.GainCrystal(10)
	p2 := game.NewPlayer(2, 30, 0, pw2, game.NewDeck())
	g := game.Resume(p1, p2, 1, nil /* rng */)
	p1.Refresh()

	t.Logf("Resume game")
	p1status := playerStatus{30, 0, 0, false, 10, 10, 0, 0, 0}
	assertPlayerStatus(t, p1, p1status)

	t.Logf("P1 uses hero power")
	action.HeroPower(g, p1, p2)
	p1status.mana -= 2
	assertPlayerStatus(t, p1, p1status)
}
