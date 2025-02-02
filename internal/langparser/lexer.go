package langparser

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	TokenSelect   TokenType = "SELECT"
	TokenFrom     TokenType = "FROM"
	TokenWhere    TokenType = "WHERE"
	TokenIdent    TokenType = "IDENTIFIER"
	TokenString   TokenType = "STRING"
	TokenNumber   TokenType = "NUMBER"
	TokenOperator TokenType = "OPERATOR"
	TokenComma    TokenType = "COMMA"
	TokenAsterix  TokenType = "ASTERIX"
	TokenEOF      TokenType = "EOF"
)

// Token represents single token in input
type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input   string
	idx     int
	readIdx int
	ch      byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readIdx >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readIdx]
	}
	l.idx = l.readIdx
	l.readIdx++
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.ch)) {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	l.readChar() // skips reading starting \'
	start := l.idx
	for l.ch != '\'' && l.ch != 0 {
		l.readChar()
	}
	end := l.idx
	l.readChar() // skips reading last \'
	return l.input[start:end]
}

func (l *Lexer) readNumber() string {
	start := l.idx
	for unicode.IsDigit(rune(l.ch)) {
		l.readChar()
	}
	return l.input[start:l.idx]
}

func (l *Lexer) readIdentifier() string {
	start := l.idx
	for unicode.IsLetter(rune(l.ch)) || unicode.IsDigit(rune(l.ch)) || l.ch == '_' {
		l.readChar()
	}

	return l.input[start:l.idx]
}

func lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"SELECT": TokenSelect,
		"FROM":   TokenFrom,
		"WHERE":  TokenWhere,
	}
	if tok, ok := keywords[strings.ToUpper(ident)]; ok {
		return tok
	}
	return TokenIdent
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	switch {
	case l.ch == ',':
		l.readChar()
		return Token{Type: TokenComma, Value: ","}
	case l.ch == '*':
		l.readChar()
		return Token{Type: TokenAsterix, Value: "*"}
	case l.ch == '\'':
		return Token{Type: TokenString, Value: l.readString()}
	case unicode.IsDigit(rune(l.ch)):
		return Token{Type: TokenNumber, Value: l.readNumber()}
	case l.ch == '=' || l.ch == '>' || l.ch == '<':
		op := string(l.ch)
		l.readChar()
		return Token{Type: TokenOperator, Value: op}
	case unicode.IsLetter(rune(l.ch)):
		ident := l.readIdentifier()
		return Token{Type: lookupIdent(ident), Value: ident}
	case l.ch == 0:
		return Token{Type: TokenEOF, Value: ""}
	default:
		l.readChar()
		return Token{Type: TokenEOF, Value: ""}
	}
}

func (l *Lexer) Tokens() []Token {
	var tokens []Token
	for {
		token := l.NextToken()
		if token.Type == TokenEOF {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}
