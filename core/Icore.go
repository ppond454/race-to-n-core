package core

type ICore interface {
	// GetWinner() Player
	// IsLegal() bool
	// GetWhoTurn() Player
	GetAllPlayer() []Player
	// GetPlayer(id string) Player
	// GetTurnCount() uint
	// CanPlay(player Player) bool
	// GetChoice() []uint8
	GetN() uint8
	GetState() State

	SetN(n uint8) error
	AddPlayer(Player) error
	CanStart() (bool, error)
	Start() error

	isRoomFull() (bool, error)
	changeState(State) error
	// MakeMove(value []uint8, player Player)
}

func New() *Match {
	return &Match{
		players: []Player{},
		state:   WAITING,
	}
}

func NewPlayer(id string) *Player {
	return &Player{
		Id:          id,
		moveHistory: []uint{},
	}
}
