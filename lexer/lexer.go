package lexer

import "chepra/token"

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// helper function to get the next character
// and advance our position in the input string
func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func (l *Lexer) readIdentifier() string {
	pos := l.pos

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) SkipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	pos := l.pos

	for isNumber(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.SkipWhitespace()

	switch l.ch {
	case '=':
		tok = l.newToken(token.EQ, l.ch)
	case ';':
		tok = l.newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = l.newToken(token.LPAREN, l.ch)
	case ')':
		tok = l.newToken(token.RPAREN, l.ch)
	case ',':
		tok = l.newToken(token.COMMA, l.ch)
	case '+':
		tok = l.newToken(token.PLUS, l.ch)
	case '-':
		tok = l.newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = l.newToken(token.BANG, l.ch)
		}
	case '*':
		tok = l.newToken(token.ASTERISK, l.ch)
	case '/':
		tok = l.newToken(token.SLASH, l.ch)
	case '<':
		tok = l.newToken(token.LT, l.ch)
	case '>':
		tok = l.newToken(token.GT, l.ch)
	case ':':
		tok = l.newToken(token.COLON, l.ch)
	case '[':
		tok = l.newToken(token.LBRACKET, l.ch)
	case ']':
		tok = l.newToken(token.RBRACKET, l.ch)
	case '{':
		tok = l.newToken(token.LBRACE, l.ch)
	case '}':
		tok = l.newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isNumber(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT

			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
