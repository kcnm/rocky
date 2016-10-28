package engine

type Game interface {
	Listener

	Events() EventBus
	Turn() int
	CurrentPlayer() Player
	Opponent(player Player) Player
	IsOver() (over bool, winner Player)

	Summon(card MinionCard, board Board, position int) Minion
}
