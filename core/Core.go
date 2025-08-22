package core

type Player struct {
	Id          string
	moveHistory []uint
}

type Match struct {
	n       uint8
	players []Player
	count   uint
	current uint8
	turn    Player
	winner  Player
}
