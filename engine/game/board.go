package game

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kcnm/rocky/engine"
)

var maxBoard = flag.Int("max_board", 7, "maximum width of board")

type board struct {
	minions []engine.Minion
}

func newBoard() engine.Board {
	return &board{make([]engine.Minion, 0, *maxBoard)}
}

func (b *board) Minions() []engine.Minion {
	return b.minions
}

func (b *board) IsFull() bool {
	return len(b.minions) >= *maxBoard
}

func (b *board) Find(minion engine.Minion) int {
	for i, m := range b.minions {
		if m == minion {
			return i
		}
	}
	return -1
}

func (b *board) Get(pos int) engine.Minion {
	return b.minions[pos]
}

func (b *board) Put(minion engine.Minion, position int) engine.Minion {
	if b.IsFull() {
		panic("board is full")
	}
	b.minions = append(
		b.minions[:position],
		append([]engine.Minion{minion}, b.minions[position:]...)...)
	return minion
}

func (b *board) Remove(minion engine.Minion) engine.Minion {
	if i := b.Find(minion); i >= 0 {
		b.minions = append(b.minions[:i], b.minions[i+1:]...)
		return minion
	}
	return nil
}

func (b *board) String() string {
	minions := make([]string, len(b.minions))
	for i, m := range b.minions {
		minions[i] = fmt.Sprintf("%v", m)
	}
	return strings.Join(minions, " ")
}
