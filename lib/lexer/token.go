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
	VARIABLE
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
    FUNCTION: "FUNCTION",
	VARIABLE: "VARIABLE",
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