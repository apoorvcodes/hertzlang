package lexer

type Token int

const (
	Iota = iota
	ILLEGAL
	IDENT
	INT
	SEMI 
	ADD 
	FUNCTION
	SUB 
	MUL 
	DIV 
	ASSIGN 
)

var tokens = []string{
	Iota:    "IOTA",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",
    FUNCTION: "fnc",
	// Infix ops
	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	ASSIGN: "=",
}

func (t Token) String() string {
	return tokens[t]
}