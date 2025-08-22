package core

type Match struct {
	n uint8
}

func (ctx *Match) CheckWin() uint8 {
	return ctx.n
}
