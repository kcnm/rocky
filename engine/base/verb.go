package base

type Verb int

const (
	StartTurn Verb = iota
	EndTurn
	Draw
	TakeCard
	PlayCard
	Summon
	Hit
	Damage
	Dying
	Destroy
	GameOver
)

func (v Verb) Subject() interface{} {
	return nil
}

func (v Verb) Verb() Verb {
	return v
}

func (v Verb) Trigger() {
}
