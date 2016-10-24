package engine

type CharacterID int

type Character interface {
	ID() CharacterID
	Attack() int
	Health() int
	Stamina() int
	Active() bool

	Assign(id CharacterID)
	Refresh()
	TakeDamage(damage int) (actual int, fatal bool)
	LoseStamina()
}
