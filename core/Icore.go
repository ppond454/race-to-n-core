package core

type ICore interface {
	CheckWin() uint8
}

func New(n uint8) ICore {
	return &Match{
		n,
	}
}
