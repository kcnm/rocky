package engine

type Verb string

type Event interface {
	Subject() interface{}
	Verb() Verb

	Trigger()
}

type ListenerID int

type Listener interface {
	Handle(ev Event)
}

type EventBus interface {
	AddListener(listener Listener) ListenerID
	RemoveListener(id ListenerID) bool
	Post(ev Event, cause Event)
	PostAndTrigger(ev Event)
	Drain()
}

const (
	StartTurn     Verb = "StartTurn"
	EndTurn       Verb = "EndTurn"
	Draw          Verb = "Draw"
	TakeCard      Verb = "TakeCard"
	PlayCard      Verb = "PlayCard"
	Summon        Verb = "Summon"
	Cast          Verb = "Cast"
	Equip         Verb = "Equip"
	HeroPower     Verb = "HeroPower"
	Attack        Verb = "Attack"
	Damage        Verb = "Damage"
	Dying         Verb = "Dying"
	Destroy       Verb = "Destroy"
	DestroyWeapon Verb = "DestroyWeapon"
	Impact        Verb = "Impact"
	GameOver      Verb = "GameOver"
	Combined      Verb = "Combined"
)

func (v Verb) Subject() interface{} {
	return nil
}

func (v Verb) Verb() Verb {
	return v
}

func (v Verb) Trigger() {
}
