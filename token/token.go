package token

type TokenKind string

const (
	LEFT_PAREN      TokenKind = "LEFT_PAREN"
	RIGHT_PAREN     TokenKind = "RIGHT_PAREN"
	LEFT_BRACE      TokenKind = "LEFT_BRACE"
	RIGHT_BRACE     TokenKind = "RIGHT_BRACE"
	STAR            TokenKind = "STAR"
	DOT             TokenKind = "DOT"
	COMMA           TokenKind = "COMMA"
	SEMI_COLON      TokenKind = "SEMI_COLON"
	PLUS            TokenKind = "PLUS"
	MINUS           TokenKind = "MINUS"
	EQUAL           TokenKind = "EQUAL"
	EQUAL_EQUAL     TokenKind = "EQUAL_EQUAL"
	BANG            TokenKind = "BANG"
	BANG_EQUAL      TokenKind = "BANG_EQUAL"
	LESS_EQUAL      TokenKind = "LESS_EQUAL"
	GREATER_EQUAL   TokenKind = "GREATER_EQUAL"
	LESS            TokenKind = "LESS"
	GREATER         TokenKind = "GREATER"
	SLASH           TokenKind = "SLASH"
	STRING          TokenKind = "STRING"
	NEW_LINE        TokenKind = "NEW_LINE"
	TAB             TokenKind = "TAB"
	CARRIAGE_RETURN TokenKind = "CARRIAGE_RETURN"
	SPACE           TokenKind = "SPACE"
	NULL            TokenKind = "NULL"
)

type Token struct {
	Kind    TokenKind
	Lex     []byte
	Literal string
}
