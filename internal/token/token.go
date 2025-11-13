package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF = iota
	ILLEGAL

	IDENT
	INT

	ASSIGN
	PLUS
	BANG
	SLASH
	MINUS
	ASTERISK

	LT
	LT_OR_EQ
	GT
	GT_OR_EQ
	EQ
	NOT_EQ

	COMMA
	SEMICOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	FUNCTION
	LET

	IF
	ELSE

	TRUE
	FALSE

	RETURN
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,

	"if":   IF,
	"else": ELSE,

	"true":  TRUE,
	"false": FALSE,

	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
