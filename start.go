package main

import (
	"fmt"
	"os"

	"github.com/hertzlng/hertz/lib/lexer"
)

func main() {
	file, err := os.Open("./example/main.hz")
	if err != nil {
		panic(err)
	}
	lex := lexer.NewLexer(file)

	for {
		pos, tok, lit := lex.Lex()
		if tok == lexer.Iota {
			break
		}

		fmt.Printf("%d:%d\t%s\t%s\n", pos.Line, pos.Column, tok, lit)
	}
}