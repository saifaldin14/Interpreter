package lexer

import (
	"testing"
	"saif/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{ token.ASSIGN, "=" }
		{ token.ASSIGN, "=" }
	}
}