package card

type minion int
type spell int
type weapon int

const (
	SilverHandRecruit minion = iota
	Fireball          spell  = iota
	FieryWarAxe       weapon = iota
)
