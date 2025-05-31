package main

type Scanner struct {
	source string
}

//	func scannerError(line int, msg string) error {
//		return fmt.Errorf("[line %d] Error: %s", line, msg)
//	}
func NewScanner(source string) Scanner {
	return Scanner{source}
}

func (s *Scanner) scanTokens() []Token {
	tokens := make([]Token, 0)
	// tokens = append(tokens, Token{1, "hi"})
	// tokens = append(tokens, Token{2, "i"})
	// tokens = append(tokens, Token{5, "h"})
	return tokens
}
