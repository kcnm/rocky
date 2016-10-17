package base

type Verb string

const (
	StartTurn     Verb = "StartTurn"
	EndTurn       Verb = "EndTurn"
	Draw          Verb = "Draw"
	TakeCard      Verb = "TakeCard"
	PlayCard      Verb = "PlayCard"
	Summon        Verb = "Summon"
	Cast          Verb = "Cast"
	Equip         Verb = "Equip"
	Attack        Verb = "Attack"
	Damage        Verb = "Damage"
	Dying         Verb = "Dying"
	Destroy       Verb = "Destroy"
	DestroyWeapon Verb = "DestroyWeapon"
	GameOver      Verb = "GameOver"
)

func (v Verb) Subject() interface{} {
	return nil
}

func (v Verb) Verb() Verb {
	return v
}

func (v Verb) Trigger() {
}
