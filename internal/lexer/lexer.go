package lexer

import "github.com/joevtap/monkey/internal/token"

type Lexer struct {
	input        string
	position     int
	peekPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		peekPosition: 0,
		ch:           0,
	}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.peekPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.peekPosition]
	}

	l.position = l.peekPosition
	l.peekPosition++
}

func (l *Lexer) peekChar(ch byte) bool {
	if l.peekPosition >= len(l.input) {
		return false
	}

	if l.input[l.peekPosition] == ch {
		l.readChar()
		return true
	}

	return false
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == '\t' || l.ch == '\n' || l.ch == '\r' || l.ch == ' ' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar('=') {
			tok.Literal = "=="
			tok.Type = token.EQ
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '!':
		if l.peekChar('=') {
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		if l.peekChar('=') {
			tok.Literal = "<="
			tok.Type = token.LT_OR_EQ
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar('=') {
			tok.Literal = ">="
			tok.Type = token.GT_OR_EQ
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
