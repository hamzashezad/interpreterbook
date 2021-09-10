package lexer

import "monkey/token"

type Lexer struct {
	input string // whole input
	currentPosition int // current character
	nextCharPosition int // next character
	// for Unicode/UTF-8 support use `rune` type: an integer value identifying a Unicode code point
	currentChar byte // stores ASCII characters
}

func New(input string) *Lexer {
	// TODO: What does `&` do?
	// XXX: sets initial values for others to default as go spec
	lexer := &Lexer{input: input}

	lexer.readChar()

	return lexer
}

func (lexer *Lexer) readChar() {
	// if reached end of input, set character to ASCII NUL
	if lexer.nextCharPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		// otherwise, the next character
		lexer.currentChar = lexer.input[lexer.nextCharPosition]
	}

	lexer.currentPosition = lexer.nextCharPosition
	lexer.nextCharPosition += 1
}

// returns a relevant `Token` for each character/"word"
func (lexer *Lexer) NextToken() token.Token {
	// TODO: what is `var`? setting uninitialised variables?
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, lexer.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		tok = newToken(token.LPAREN, lexer.currentChar)
	case ')':
		tok = newToken(token.RPAREN, lexer.currentChar)
	case '{':
		tok = newToken(token.LBRACE, lexer.currentChar)
	case '}':
		tok = newToken(token.RBRACE, lexer.currentChar)
	case ',':
		tok = newToken(token.COMMA, lexer.currentChar)
	case '+':
		tok = newToken(token.PLUS, lexer.currentChar)
	// ASCII NUL character
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// if no single character matches, check for whether it is a letter
		// they are either identifiers (in which case return the whole word)) or illegal
		if isLetter(lexer.currentChar) {
			identifier := lexer.readIdentifier()

			tok.Literal = identifier
			tok.Type = token.LookupIdent(identifier)

			return tok
		} else if isDigit(lexer.currentChar) {
			digit := lexer.readDigit()

			tok.Literal = digit
			tok.Type = token.INT

			// early return because readChar in readDigit ends up one after the digit
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.currentChar)
		}
	}

	lexer.readChar()

	return tok
}

// valid identifiers can only begin [a-zA-Z_], no integers allowed
func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.currentPosition

	for isLetter(lexer.currentChar) {
		// XXX: advance to next character as long it is a letter
		lexer.readChar()
	}

	endPosition := lexer.currentPosition

	// XXX: do not miss start : end position
	return lexer.input[startPosition:endPosition]
}

func (lexer *Lexer) readDigit() string {
	startPosition := lexer.currentPosition

	for isDigit(lexer.currentChar) {
		// XXX: advance to next character as long current is a digit
		// XXX: Caveat; when current (last) character is a digit, advances to next (non-digit),
		// so early return is needed
		lexer.readChar()
	}

	endPosition := lexer.currentPosition

	// XXX: do not miss start : end position
	return lexer.input[startPosition:endPosition]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.currentChar == ' ' || lexer.currentChar == '\t' || lexer.currentChar == '\n' || lexer.currentChar == '\r' {
		lexer.readChar()
	}
}

// Token constructor/initializer
func newToken(tokenType token.TokenType, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

// checks whether a character ASCII code is between that of [a-zA-Z_]
func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_';
}
