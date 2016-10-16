package base

type Board interface {
	Listener

	Minions() []Minion
	IsFull() bool
	Find(minion Minion) (index int)
	Get(pos int) Minion

	Put(minion Minion, toRight Minion) Minion
}
