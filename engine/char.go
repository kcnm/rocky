package engine

type CharID int

type Char interface {
	ID() CharID
	Attack() int
	Health() int
	Stamina() int
	Active() bool

	Refresh()
	TakeDamage(damage int) (actual int, fatal bool)
	LoseStamina()
}
