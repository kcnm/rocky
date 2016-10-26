package target

type Assign string

const (
	None   Assign = "None"
	All    Assign = "All"
	Random Assign = "Random"
	Manual Assign = "Manual"
)

type Side string

const (
	Any    Side = "Any"
	Friend Side = "Friend"
	Enemy  Side = "Enemy"
)

type Role string

const (
	Minion Role = "Minion"
	Player Role = "Player"
	Char   Role = "Char"
)

type Spec interface {
	Assign() Assign
	Side() Side
	Role() Role
}
