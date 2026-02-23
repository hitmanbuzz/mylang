package lexer

import (
	"fmt"
	"mylang/token"
)

// Get the current index value (which is in byte)
func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	return l.Source[l.CurrIdx]
}

// Increment the current index by the passed parameter value
func (l *Lexer) advance(num int) {
	if l.CurrIdx+num > len(l.Source) {
		panic(fmt.Sprintf("cannot advance by `%d` due to index out of range", num))
	} else {
		l.CurrIdx += num
	}
}

// Check if it reach the end of the source code/file
func (l *Lexer) isAtEnd() bool {
	return l.CurrIdx >= len(l.Source)
}

// Add new token based on the given parameters data
func (l *Lexer) addToken(kind token.TokenKind, lex []byte, literal string) {
	l.tokens = append(l.tokens, token.Token{
		Kind:    kind,
		Lex:     lex,
		Literal: literal,
	})
}

// Printing out all tokens with ther kind, lex & literal
func (l *Lexer) Display() {
	for _, t := range l.tokens {
		switch t.Kind {
		case token.NEW_LINE:
			continue
		case token.TAB:
			continue
		case token.CARRIAGE_RETURN:
			continue
		case token.STRING, token.NUMBER:
			fmt.Printf("%s \"%s\" %s\n", t.Kind, t.Lex, t.Literal)
		default:
			fmt.Printf("%s %s %s\n", t.Kind, t.Lex, t.Literal)
		}
	}

	fmt.Println("EOF  null")
}
