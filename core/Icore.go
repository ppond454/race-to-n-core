package core

type ICore interface {
	GetWinner() Player
	MakeMove(value []uint8, player Player)
	IsLegal() bool
	GetWhoTurn() Player
	GetAllPlayer() []Player
	GetPlayer(id string) Player
	GetTurnCount() uint
	CanPlay(player Player) bool
	GetChoice() []uint8

	SetN(n uint8) ICore
	AddPlayer(Player)
	CanStart() bool
	Start()
}
