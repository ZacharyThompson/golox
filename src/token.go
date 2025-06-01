package main

import "fmt"

type TokenType int

const (
	EOF TokenType = iota + 1
	// Single-character tokens.
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR
	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	// Literals.
	IDENTIFIER
	STRING
	NUMBER
	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType: tokenType, lexeme: lexeme, literal: literal, line: line}
}

func (t Token) toString() string {
	return fmt.Sprintf("%d %s %v", t.line, t.lexeme, t.literal)
}
