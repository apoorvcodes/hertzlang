package lexer

import (
	"bufio"
	"fmt"
	// "fmt"
	"io"
	"unicode"
)

type lexer struct {
	pos position
	buf *bufio.Reader
}

func NewLexer(reader io.Reader) *lexer {
	return &lexer{
		pos: NewPos(),
		buf: bufio.NewReader(reader),
	}
}

func (l *lexer) Lex() (position, Token, string) {
	for {
		r, _, err := l.buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.pos, Iota, ""
			}
			panic(err)
		}
		l.pos.Column ++

		switch r {
		 case '\n' :
			l.reset()
		 case ';':
            return l.pos, SEMI, ";"
         case '+':
            return l.pos, ADD, "+"
         case '-':
            return l.pos, SUB, "-"
         case '*':
            return l.pos, MUL, "*"
         case '/':
            return l.pos, DIV, "/"
         case '=':
            return l.pos, ASSIGN, "=" 
		 default:
			if unicode.IsSpace(r) {
				continue
			}else if  unicode.IsDigit(r) {
				l.backup()
				lit := l.lexInt()
				return l.pos, INT, lit
				
			}else if unicode.IsLetter(r) {
				l.backup()
				lit := l.lexIdent()
				switch lit {
				case "let":
				 return l.pos, VARIABLE, "let" 
				case "fn": 
				return l.pos, FUNCTION, "fn"
				default:
					return l.pos, IDENT, lit
				} 
				}else  {
				strg := string(r)
				fmt.Print(strg + "\n")
			}

		}
	}
}


func (l *lexer) reset() {
	l.pos.Line++
	l.pos.Column = 0
}


func (l *lexer) backup() {
	if err := l.buf.UnreadRune(); err != nil {
		panic(err)
	}
	
	l.pos.Column--
}

func (l *lexer) lexInt() string {
	var lit string
	for {
		r, _, err := l.buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}

		l.pos.Column++
		if unicode.IsDigit(r) {
			lit = lit + string(r)
		} else {
			// scanned something not in the integer
			l.backup()
			return lit
		}
	}
}
func (l *lexer) lexIdent() string {
	var lit string
	for {
		r, _, err := l.buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				return lit
			}
		}
			
        l.pos.Column++
		if unicode.IsLetter(r) {
			lit = lit + string(r)
		} else {
			l.backup()
			return lit
		}
	}
}