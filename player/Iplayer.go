package player

func NewPlayer(id string) *Player {
	return &Player{
		Id:          id,
		moveHistory: []uint{},
	}
}
