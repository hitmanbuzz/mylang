package lexer

import (
	"fmt"
	"mylang/token"
)

type Lexer struct {
	tokens   []token.Token
	CurrIdx  int
	LineNo   int
	Source   string
	ExitCode int
}

func NewLexer() *Lexer {
	return &Lexer{
		tokens:   make([]token.Token, 0),
		CurrIdx:  0,
		LineNo:   1,
		ExitCode: 0,
	}
}

func (l *Lexer) Tokenize() int {
	nt := token.Token{
		Kind:    token.NULL,
		Lex:     []byte{},
		Literal: "",
	}

	nextByte := l.Source[l.CurrIdx+1]
	kind, lex, literal := l.MatchToken(l.CurrIdx, l.peek(), nextByte, l.Source)
	if kind != token.NULL {
		nt.Lex = append(nt.Lex, lex)
		nt.Kind = kind
		nt.Literal = literal
		l.tokens = append(l.tokens, nt)
	}

	return l.CurrIdx
}

// @return
//
// 1st -> TokenType
//
// 2nd  -> Lex
//
// 3rd  -> Literal
//
// The 4th return value is the number of indices to skip from the current char.
func (l *Lexer) MatchToken(currIdx int, currByte byte, nextByte byte, source string) (token.TokenKind, byte, string) {
	switch currByte {
	case ' ':
		l.advance(1)
		return token.SPACE, currByte, "<SPACE>"
	case '\t':
		l.advance(1)
		return token.TAB, currByte, "\\t"
	case '\r':
		l.advance(1)
		return token.CARRIAGE_RETURN, currByte, "\\r"
	case '\n':
		l.advance(1)
		l.LineNo += 1
		return token.NEW_LINE, currByte, "null"
	case '(':
		l.advance(1)
		return token.LEFT_PAREN, currByte, "null"
	case ')':
		l.advance(1)
		return token.RIGHT_PAREN, currByte, "null"
	case '{':
		l.advance(1)
		return token.LEFT_BRACE, currByte, "null"
	case '}':
		l.advance(1)
		return token.RIGHT_BRACE, currByte, "null"
	case ',':
		l.advance(1)
		return token.COMMA, currByte, "null"
	case '.':
		l.advance(1)
		return token.DOT, currByte, "null"
	case '=':
		if nextByte == '=' {
			l.advance(2)
			return token.EQUAL_EQUAL, currByte, "null"
		}
		l.advance(1)
		return token.EQUAL, currByte, "null"
	case '!':
		if nextByte == '!' {
			l.advance(2)
			return token.BANG_EQUAL, currByte, "null"
		}
		l.advance(1)
		return token.BANG, currByte, "null"
	case '>':
		if nextByte == '=' {
			l.advance(2)
			return token.GREATER_EQUAL, currByte, "null"
		}
		l.advance(1)
		return token.GREATER, currByte, "null"
	case '<':
		if nextByte == '=' {
			l.advance(2)
			return token.LESS_EQUAL, currByte, "null"
		}
		l.advance(1)
		return token.LESS, currByte, "null"
	case '+':
		l.advance(1)
		return token.PLUS, currByte, "null"
	case '-':
		l.advance(1)
		return token.MINUS, currByte, "null"
	case ';':
		l.advance(1)
		return token.SEMI_COLON, currByte, "null"
	case '/':
		if nextByte == '/' {
			l.scanComment()
			return token.NULL, 0, "null"
		}
		l.advance(1)
		return token.SLASH, currByte, "null"
	case '*':
		l.advance(1)
		return token.STAR, currByte, "null"
	default:
		fmt.Printf("[line %d] Error: Unexpected character: %b\n", l.LineNo, currByte)
		l.ExitCode = 65
		l.advance(1)
		return token.NULL, currByte, ""
	}
}

func (l *Lexer) scanString() {

}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	return l.Source[l.CurrIdx]
}

func (l *Lexer) advance(num int) {
	if l.CurrIdx+num > len(l.Source) {
		panic(fmt.Sprintf("cannot advance by `%d` due to index out of range", num))
	} else {
		l.CurrIdx += num
	}
}

func (l *Lexer) isAtEnd() bool {
	return l.CurrIdx >= len(l.Source)
}

func (l *Lexer) scanComment() {
	for {
		s := l.peek()
		if l.isAtEnd() {
			break
		}

		if s == '\n' {
			l.LineNo++
			l.advance(1)
			break
		}
		l.advance(1)
	}
}

func (l *Lexer) Display() {
	for _, t := range l.tokens {
		fmt.Printf("%s %s %s\n", t.Kind, t.Lex, t.Literal)
	}

	fmt.Println("EOF  null")
}

// this will return the length of the comment so that we can skip that amount in scanning/lexing
