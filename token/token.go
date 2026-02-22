package token

type TokenType string

const (
	LEFT_PAREN      TokenType = "LEFT_PAREN"
	RIGHT_PAREN     TokenType = "RIGHT_PAREN"
	LEFT_BRACE      TokenType = "LEFT_BRACE"
	RIGHT_BRACE     TokenType = "RIGHT_BRACE"
	STAR            TokenType = "STAR"
	DOT             TokenType = "DOT"
	COMMA           TokenType = "COMMA"
	SEMI_COLON      TokenType = "SEMI_COLON"
	PLUS            TokenType = "PLUS"
	MINUS           TokenType = "MINUS"
	EQUAL           TokenType = "EQUAL"
	EQUAL_EQUAL     TokenType = "EQUAL_EQUAL"
	BANG            TokenType = "BANG"
	BANG_EQUAL      TokenType = "BANG_EQUAL"
	LESS_EQUAL      TokenType = "LESS_EQUAL"
	GREATER_EQUAL   TokenType = "GREATER_EQUAL"
	LESS            TokenType = "LESS"
	GREATER         TokenType = "GREATER"
	SLASH           TokenType = "SLASH"
	STRING          TokenType = "STRING"
	NEW_LINE        TokenType = "NEW_LINE"
	TAB             TokenType = "TAB"
	CARRIAGE_RETURN TokenType = "CARRIAGE_RETURN"
	SPACE           TokenType = "SPACE"
	NULL            TokenType = "NULL"
)

type Token struct {
	Kind    TokenType
	Lex     []byte
	Literal string
}
