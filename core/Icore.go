package core

import "github.com/ppond454/race-to-n-core/player"

type ICore interface {
	// GetWinner() Player
	// IsLegal() bool
	// GetWhoTurn() Player
	GetAllPlayer() []player.Player
	// GetPlayer(id string) Player
	// GetTurnCount() uint
	// CanPlay(player Player) bool
	GetChoice() ([]uint8, error)
	GetN() uint8
	GetState() State

	SetN(n uint8) error
	AddPlayer(player.Player) error
	CanStart() (bool, error)
	Start() error

	isRoomFull() (bool, error)
	changeState(State) error
	// MakeMove(value []uint8, player Player)
}

func New() ICore {
	return &Match{
		players: []player.Player{},
		state:   WAITING,
	}
}
