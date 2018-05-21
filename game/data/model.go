package data


type Position struct {
	X, Y int
}

type Token int

const (
	Unset Token = iota
	X
	O
)

func (t Token) IsSet() bool {
	return t != Unset
}