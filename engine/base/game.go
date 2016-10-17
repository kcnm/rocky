package base

type Game interface {
	Listener

	Events() *EventBus
	CurrentPlayer() Player
	Opponent(player Player) Player
	IsOver() (over bool, winner Player)

	Summon(card MinionCard, board Board, toRight Minion) Minion
}
