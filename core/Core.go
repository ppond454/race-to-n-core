package core

import (
	"errors"
)

type State int

const (
	WAITING State = iota
	IDLE
	PLAYING
)

type Player struct {
	Id          string
	moveHistory []uint
}

type Match struct {
	n       *uint8
	players []Player //TODO: Make to Set
	count   uint
	current uint8
	turn    Player
	winner  Player
	state   State
}

func (ctx *Match) SetN(value uint8) error {
	if ctx.state == PLAYING {
		return errors.New("can not set N when playing")
	}
	if ctx.n != nil {
		return errors.New("the n already set")
	}

	if value > 30 {
		return errors.New("the n should not over than 30")
	}

	if value < 10 {
		return errors.New("the n should not less than 10")
	}

	ctx.n = &value
	return nil
}

func (ctx *Match) isRoomFull() (bool, error) {
	if len(ctx.players) == 2 {
		return true, nil
	}
	return false, errors.New("room is not full")
}

func (ctx *Match) AddPlayers(p Player) error {
	if isFull, err := ctx.isRoomFull(); isFull {
		return err
	}
	ctx.players = append(ctx.players, p)
	return nil
}

func (ctx *Match) GetN() uint8 {
	return *ctx.n
}

func (ctx *Match) GetAllPlayer() []Player {
	return ctx.players
}

func (ctx *Match) CanStart() (bool, error) {
	if isFull, err := ctx.isRoomFull(); !isFull {
		return false, err
	}
	if ctx.state != IDLE {
		return false, errors.New("game is not idle state")
	}
	return true, nil
}

func (ctx *Match) Start() error {
	if can, err := ctx.CanStart(); !can {
		return err
	}
	return ctx.changeState(PLAYING)
}

func (ctx *Match) GetState() State {
	return ctx.state
}

func (ctx *Match) changeState(state State) error {
	ctx.state = state
	return nil
}
