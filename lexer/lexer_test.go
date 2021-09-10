package lexer

import (
	"testing"

	"monkey/token"
)

type GeneralTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

type DefaultValueTest struct {
	currentPosition int // current character
	nextCharPosition int // next character
	currentChar byte // stores ASCII characters
}

func TestBasicToken(t *testing.T) {
	input := "={}5;)"

	expectedResult := []GeneralTest {
		{ token.ASSIGN, "=" },
		{ token.LBRACE, "{" },
		{ token.RBRACE, "}" },
		{ token.INT, "5" },
		{ token.SEMICOLON, ";" },
		{ token.RPAREN, ")" },
	}

	lexer := New(input)

	for i, result := range expectedResult {
		got := lexer.NextToken()

		if got.Type != result.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got=%q", i, result.expectedType, got.Type)
		}

		if got.Literal != result.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, result.expectedLiteral, got.Literal)
		}
	}
}

func TestDefaultValue(t *testing.T) {
	input := `let five = 5;`

	expectedResult := DefaultValueTest {
		0,
		1,
		'l',
	}

	lexer := New(input)

	if lexer.currentPosition != expectedResult.currentPosition {
		t.Fatalf("wrong position. expected=%q, got=%q", lexer.currentPosition, expectedResult.currentPosition)
	}

	if lexer.nextCharPosition != expectedResult.nextCharPosition {
		t.Fatalf("wrong readPosition. expected=%q, got=%q", lexer.nextCharPosition, expectedResult.nextCharPosition)
	}

	if lexer.currentChar != expectedResult.currentChar {
		t.Fatalf("wrong character. expected=%q, got=%q", lexer.currentChar, expectedResult.currentChar)
	}
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;`

	expectedResult := []GeneralTest {
		{ token.LET, "let" },
		{ token.IDENT, "five" },
		{ token.ASSIGN, "=" },
		{ token.INT, "5" },
		{ token.SEMICOLON, ";" },
	}

	lexer := New(input)

	for i, result := range expectedResult {
		got := lexer.NextToken()

		if got.Type != result.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got=%q", i, result.expectedType, got.Type)
		}

		if got.Literal != result.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, result.expectedLiteral, got.Literal)
		}
	}
}

func TestBasicSyntax(t *testing.T) {
	input := `let x = 5;
	let y = 10;
	let add = fn(x, y) {
		x + y;
	};

	let result = add(x, y);`

	expectedResult := []GeneralTest {
		{ token.LET, "let" },
		{ token.IDENT, "x" },
		{ token.ASSIGN, "=" },
		{ token.INT, "5" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "y" },
		{ token.ASSIGN, "=" },
		{ token.INT, "10" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "add" },
		{ token.ASSIGN, "=" },
		{ token.FUNCTION, "fn" },
		{ token.LPAREN, "(" },
		{ token.IDENT, "x" },
		{ token.COMMA, "," },
		{ token.IDENT, "y" },
		{ token.RPAREN, ")" },
		{ token.LBRACE, "{" },
		{ token.IDENT, "x" },
		{ token.PLUS, "+" },
		{ token.IDENT, "y" },
		{ token.SEMICOLON, ";" },
		{ token.RBRACE, "}" },
		{ token.SEMICOLON, ";" },
		{ token.LET, "let" },
		{ token.IDENT, "result" },
		{ token.ASSIGN, "=" },
		{ token.IDENT, "add" },
		{ token.LPAREN, "(" },
		{ token.IDENT, "x" },
		{ token.COMMA, "," },
		{ token.IDENT, "y" },
		{ token.RPAREN, ")" },
		{ token.SEMICOLON, ";" },
	}

	lexer := New(input)

	for i, result := range expectedResult {
		got := lexer.NextToken()

		if got.Type != result.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected %q, got=%q", i, result.expectedType, got.Type)
		}

		if got.Literal != result.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, result.expectedLiteral, got.Literal)
		}
	}
}
