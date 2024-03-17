package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string // the original input string
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination i.e. input[position]
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // init the Lexer by reading the first char
	return l
}

// Reads the next character in the input and puts in the ch variable
func (l *Lexer) readChar() {
	// we are never returning the char.
	// we are just updating the pointer to which its pointing to.

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// we are just looking at the next character.
// we are not updating the position and readposition pointers because we dont want to move the lexer.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Returns the appropriate token for the entire input
// Switches on the current input character and processes the next characters accordingly
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// read an identifier (which is a word/string of characters)
func (l *Lexer) readIdentifier() string {
	// readIdentifier and readNumber follows the same pattern
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// read a multi-digit number (i.e.234)
// If the current character is a digit, then read the next character
func (l *Lexer) readNumber() string {
	position := l.position
	// as long as as the current character is a digit keep reading the next character
	for isDigit(l.ch) {
		l.readChar()
	}
	// if the current character is not a digit
	// then the number is from position to (l.position - 1)
	// recall: how slice-indexig works
	return l.input[position:l.position]
}

// checks if the character being examined    is a digit or not
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Is the current character under inspection is a Letter ?
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// If the current character is whitespace, it reads the next one
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
