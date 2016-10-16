package impl

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kcnm/rocky/engine/base"
)

var maxBoard = flag.Int("max_board", 7, "maximum width of board")

type board struct {
	minions []base.Minion
}

func newBoard() base.Board {
	return &board{make([]base.Minion, 0, *maxBoard)}
}

func (b *board) Handle(ev base.Event) {
	switch ev.Verb() {
	case base.Destroy:
		if minion, ok := ev.Subject().(base.Minion); ok {
			if i := b.Find(minion); i >= 0 {
				b.minions = append(b.minions[:i], b.minions[i+1:]...)
			}
		}
	}
}

func (b *board) Minions() []base.Minion {
	return b.minions
}

func (b *board) IsFull() bool {
	return len(b.minions) >= *maxBoard
}

func (b *board) Find(minion base.Minion) int {
	for i, m := range b.minions {
		if m == minion {
			return i
		}
	}
	return -1
}

func (b *board) Get(pos int) base.Minion {
	if pos < 0 || len(b.minions) <= pos {
		return nil
	}
	return b.minions[pos]
}

func (b *board) Put(minion base.Minion, toRight base.Minion) base.Minion {
	if b.IsFull() {
		return nil
	}
	if toRight == nil {
		b.minions = append([]base.Minion{minion}, b.minions...)
		return minion
	}
	for i, m := range b.minions {
		if m == toRight {
			b.minions = append(
				b.minions[:i+1],
				append([]base.Minion{minion}, b.minions[i+1:]...)...)
			return minion
		}
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
