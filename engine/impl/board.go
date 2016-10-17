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
		b.remove(ev.Subject())
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

func (b *board) Put(minion base.Minion, position int) base.Minion {
	if b.IsFull() {
		return nil
	}
	b.minions = append(
		b.minions[:position],
		append([]base.Minion{minion}, b.minions[position:]...)...)
	return minion
}

func (b *board) remove(subject interface{}) int {
	if subjects, ok := subject.([]interface{}); ok {
		removed := 0
		for _, sub := range subjects {
			removed += b.removeSingle(sub)
		}
		return removed
	}
	return b.removeSingle(subject)
}

func (b *board) removeSingle(subject interface{}) int {
	if minion, ok := subject.(base.Minion); ok {
		if i := b.Find(minion); i >= 0 {
			b.minions = append(b.minions[:i], b.minions[i+1:]...)
			return 1
		}
	}
	return 0
}

func (b *board) String() string {
	minions := make([]string, len(b.minions))
	for i, m := range b.minions {
		minions[i] = fmt.Sprintf("%v", m)
	}
	return strings.Join(minions, " ")
}
