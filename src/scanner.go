package main

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source               string
	tokens               []Token
	start, current, line int
}

func (Scanner) keywordsMap() map[string]TokenType {
	return map[string]TokenType{
		"and":    AND,
		"class":  CLASS,
		"else":   ELSE,
		"false":  FALSE,
		"for":    FOR,
		"fun":    FUN,
		"if":     IF,
		"nil":    NIL,
		"or":     OR,
		"print":  PRINT,
		"return": RETURN,
		"super":  SUPER,
		"this":   THIS,
		"true":   TRUE,
		"var":    VAR,
		"while":  WHILE,
	}
}

func scannerError(line int, msg string) error {
	return fmt.Errorf("[line %d] Error: %s", line, msg)
}

func NewScanner(source string) Scanner {
	return Scanner{source, []Token{}, 0, 0, 1}
}

func (s *Scanner) scanTokens() ([]Token, []error) {
	errs := []error{}
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			errs = append(errs, err)
		}
	}
	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))
	return s.tokens, errs
}

func (s *Scanner) scanToken() error {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		t := BANG
		if s.match('=') {
			t = BANG_EQUAL
		}
		s.addToken(t)
	case '=':
		t := EQUAL
		if s.match('=') {
			t = EQUAL_EQUAL
		}
		s.addToken(t)
	case '<':
		t := LESS
		if s.match('=') {
			t = LESS_EQUAL
		}
		s.addToken(t)
	case '>':
		t := GREATER
		if s.match('=') {
			t = GREATER_EQUAL
		}
		s.addToken(t)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		break
	case '\n':
		s.line++
	case '"':
		err := s.string()
		if err != nil {
			return err
		}
	default:
		if s.isDigit(c) {
			err := s.number()
			if err != nil {
				return err
			}
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			return scannerError(s.line, "Unexpected character")
		}
	}
	return nil
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func (s *Scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType := s.keywordsMap()[text]
	if tokenType == 0 {
		tokenType = IDENTIFIER
	}
	s.addToken(tokenType)
}

func (s *Scanner) number() error {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	// Look for a frational part
	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume the "."
		s.advance()
		for s.isDigit(s.peek()) {
			s.advance()
		}
	}
	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return scannerError(s.line, "Failed to parse float64")
	}
	s.addTokenLiteral(NUMBER, f)
	return nil
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		return scannerError(s.line, "Unterminated string")
	}

	// The closing "
	s.advance()

	// Trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)
	return nil
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

func (s *Scanner) addTokenLiteral(tokenType TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, NewToken(tokenType, text, literal, s.line))
}
