package lexer


type position struct {
	line int
	column int
}

func NewPos() *position {
	return &position{}
}