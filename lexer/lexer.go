package lexer

import (
	"fmt"
	"mylang/token"
	"mylang/utils"
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
	nextByte := l.Source[l.CurrIdx+1]
	l.MatchToken(l.CurrIdx, l.peek(), nextByte, l.Source)
	return l.CurrIdx
}

// @return
//
// 1st -> TokenKind
//
// 2nd  -> Lex
//
// 3rd  -> Literal
//
// The 4th return value is the number of indices to skip from the current char.
func (l *Lexer) MatchToken(currIdx int, currByte byte, nextByte byte, source string) {
	switch currByte {
	case ' ':
		l.AddToken(token.SPACE, utils.ToByteArr(currByte), "<SPACE>")
		l.advance(1)
	case '\t':
		l.AddToken(token.SPACE, utils.ToByteArr(currByte), "\\t")
		l.advance(1)
	case '\r':
		l.AddToken(token.CARRIAGE_RETURN, utils.ToByteArr(currByte), "\\r")
		l.advance(1)
	case '\n':
		l.AddToken(token.NEW_LINE, utils.ToByteArr(currByte), "null")
		l.advance(1)
		l.incrementLine()
	case '(':
		l.AddToken(token.LEFT_PAREN, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ')':
		l.AddToken(token.RIGHT_PAREN, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '{':
		l.AddToken(token.LEFT_BRACE, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '}':
		l.AddToken(token.RIGHT_BRACE, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ',':
		l.AddToken(token.COMMA, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '.':
		l.AddToken(token.DOT, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '=':
		if nextByte == '=' {
			l.advance(2)
			l.AddToken(token.EQUAL_EQUAL, utils.ToByteArr(currByte), "null")
		} else {
			l.AddToken(token.EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '!':
		if nextByte == '=' {
			l.AddToken(token.BANG_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.AddToken(token.BANG, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '>':
		if nextByte == '=' {
			l.AddToken(token.GREATER_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.AddToken(token.GREATER, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '<':
		if nextByte == '=' {
			l.AddToken(token.LESS_EQUAL, utils.ToByteArr(currByte), "null")
			l.advance(2)
		} else {
			l.AddToken(token.LESS, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '+':
		l.AddToken(token.PLUS, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '-':
		l.AddToken(token.MINUS, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case ';':
		l.AddToken(token.SEMI_COLON, utils.ToByteArr(currByte), "null")
		l.advance(1)
	case '/':
		if nextByte == '/' {
			l.scanComment()
		} else {
			l.AddToken(token.SLASH, utils.ToByteArr(currByte), "null")
			l.advance(1)
		}
	case '"':
		isString := l.scanString()
		if !isString {
			// todo
		}
	case '*':
		l.AddToken(token.STAR, utils.ToByteArr(currByte), "null")
		l.advance(1)
	default:
		fmt.Printf("[line %d] Error: Unexpected character: %b\n", l.LineNo, currByte)
		l.ExitCode = 65
		l.advance(1)
	}
}

func (l *Lexer) scanString() bool {
	isString := false

	return isString
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

// Increment the Line No. Counter
func (l *Lexer) incrementLine() {
	l.LineNo++
}

func (l *Lexer) scanComment() {
	for {
		s := l.peek()
		if l.isAtEnd() {
			break
		}

		if s == '\n' {
			l.incrementLine()
			l.advance(1)
			break
		}
		l.advance(1)
	}
}

func (l *Lexer) AddToken(kind token.TokenKind, lex []byte, literal string) {
	l.tokens = append(l.tokens, token.Token{
		Kind:    kind,
		Lex:     lex,
		Literal: literal,
	})
}

func (l *Lexer) Display() {
	for _, t := range l.tokens {
		fmt.Printf("%s %s %s\n", t.Kind, t.Lex, t.Literal)
	}

	fmt.Println("EOF  null")
}
