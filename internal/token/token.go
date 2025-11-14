package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "ASSIGN"
	PLUS = "PLUS"
	BANG = "BANG"
	SLASH = "SLASH"
	MINUS = "MINUS"
	ASTERISK = "ASTERISK"

	LT = "LT"
	LT_OR_EQ = "LT_OR_EQ"
	GT = "GT"
	GT_OR_EQ = "GT_OR_EQ"
	EQ = "EQ"
	NOT_EQ = "NOT_EQ"

	COMMA = "COMMA"
	SEMICOLON = "SEMICOLON"

	LPAREN = "LPAREN"
	RPAREN = "RPAREN"
	LBRACE = "LBRACE"
	RBRACE = "RBRACE"

	FUNCTION = "FUNCTION"

	LET = "LET"

	IF = "IF"
	ELSE = "ELSE"

	TRUE = "TRUE"
	FALSE = "FALSE"

	RETURN = "RETURN"
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
