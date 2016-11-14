package engine

type CharID int

type Char interface {
	Listener

	ID() CharID
	Attack() int
	Health() int
	MaxHealth() int
	Stamina() int
	Active() bool

	Refresh()
	AddHandler(handler Handler)
	TakeDamage(damage int) (actual int, fatal bool)
	LoseStamina()
}
