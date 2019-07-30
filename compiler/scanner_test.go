package compiler

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestSingleToken tests for whether the token type, lexeme, and line is set
// Lexeme correctness is tested elsewhere
func TestSingleToken(t *testing.T) {
	tt := []struct {
		Source         string
		ExpectedTokens []Token
		ExpectedError  []error
	}{
		{
			Source: "(",
			ExpectedTokens: []Token{
				{
					Type:    LEFT_PAREN,
					Lexeme:  "(",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ")",
			ExpectedTokens: []Token{
				{
					Type:    RIGHT_PAREN,
					Lexeme:  ")",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "[",
			ExpectedTokens: []Token{
				{
					Type:    LEFT_SQUARE,
					Lexeme:  "[",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "]",
			ExpectedTokens: []Token{
				{
					Type:    RIGHT_SQUARE,
					Lexeme:  "]",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "{",
			ExpectedTokens: []Token{
				{
					Type:    LEFT_BRACE,
					Lexeme:  "{",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "}",
			ExpectedTokens: []Token{
				{
					Type:    RIGHT_BRACE,
					Lexeme:  "}",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ",",
			ExpectedTokens: []Token{
				{
					Type:    COMMA,
					Lexeme:  ",",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ".",
			ExpectedTokens: []Token{
				{
					Type:    DOT,
					Lexeme:  ".",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "-",
			ExpectedTokens: []Token{
				{
					Type:    MINUS,
					Lexeme:  "-",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "+",
			ExpectedTokens: []Token{
				{
					Type:    PLUS,
					Lexeme:  "+",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "*",
			ExpectedTokens: []Token{
				{
					Type:    STAR,
					Lexeme:  "*",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ":",
			ExpectedTokens: []Token{
				{
					Type:    COLON,
					Lexeme:  ":",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "/",
			ExpectedTokens: []Token{
				{
					Type:    SLASH,
					Lexeme:  "/",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source:         "//This is a comment",
			ExpectedTokens: nil,
			ExpectedError:  nil,
		},
		{
			Source: "/* This is \n " +
				"a block comment */",
			ExpectedTokens: nil,
			ExpectedError:  nil,
		},
		{
			Source: "0",
			ExpectedTokens: []Token{
				{
					Type:    INT,
					Lexeme:  "0",
					Literal: 0,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "10",
			ExpectedTokens: []Token{
				{
					Type:    INT,
					Lexeme:  "10",
					Literal: 10,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "0.1",
			ExpectedTokens: []Token{
				{
					Type:    FLOAT,
					Lexeme:  "0.1",
					Literal: 0.1,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "10.1",
			ExpectedTokens: []Token{
				{
					Type:    FLOAT,
					Lexeme:  "10.1",
					Literal: 10.1,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "<",
			ExpectedTokens: []Token{
				{
					Type:    LESS,
					Lexeme:  "<",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ">",
			ExpectedTokens: []Token{
				{
					Type:    GREATER,
					Lexeme:  ">",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "!",
			ExpectedTokens: []Token{
				{
					Type:    BANG,
					Lexeme:  "!",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "=",
			ExpectedTokens: []Token{
				{
					Type:    EQUAL,
					Lexeme:  "=",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "==",
			ExpectedTokens: []Token{
				{
					Type:    EQUAL_EQUAL,
					Lexeme:  "==",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "!=",
			ExpectedTokens: []Token{
				{
					Type:    BANG_EQUAL,
					Lexeme:  "!=",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: ">=",
			ExpectedTokens: []Token{
				{
					Type:    GREATER_EQUAL,
					Lexeme:  ">=",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "<=",
			ExpectedTokens: []Token{
				{
					Type:    LESS_EQUAL,
					Lexeme:  "<=",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "||",
			ExpectedTokens: []Token{
				{
					Type:    PIPE_PIPE,
					Lexeme:  "||",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "&&",
			ExpectedTokens: []Token{
				{
					Type:    AMPER_AMPER,
					Lexeme:  "&&",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "val",
			ExpectedTokens: []Token{
				{
					Type:    VAL,
					Lexeme:  "val",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "var",
			ExpectedTokens: []Token{
				{
					Type:    VAR,
					Lexeme:  "var",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "fun",
			ExpectedTokens: []Token{
				{
					Type:    FUN,
					Lexeme:  "fun",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "for",
			ExpectedTokens: []Token{
				{
					Type:    FOR,
					Lexeme:  "for",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "if",
			ExpectedTokens: []Token{
				{
					Type:    IF,
					Lexeme:  "if",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "else",
			ExpectedTokens: []Token{
				{
					Type:    ELSE,
					Lexeme:  "else",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "null",
			ExpectedTokens: []Token{
				{
					Type:    NULL,
					Lexeme:  "null",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "True",
			ExpectedTokens: []Token{
				{
					Type:    TRUE,
					Lexeme:  "True",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "False",
			ExpectedTokens: []Token{
				{
					Type:    FALSE,
					Lexeme:  "False",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "print",
			ExpectedTokens: []Token{
				{
					Type:    PRINT,
					Lexeme:  "print",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "return",
			ExpectedTokens: []Token{
				{
					Type:    RETURN,
					Lexeme:  "return",
					Literal: nil,
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "foo",
			ExpectedTokens: []Token{
				{
					Type:    IDENTIFIER,
					Lexeme:  "foo",
					Literal: "foo",
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "\"Hello World!\"",
			ExpectedTokens: []Token{
				{
					Type:    STRING,
					Lexeme:  "\"Hello World!\"",
					Literal: "Hello World!",
					Line:    0,
				},
			},
			ExpectedError: nil,
		},
		{
			Source: "   \nfoo",
			ExpectedTokens: []Token{
				{
					Type:    IDENTIFIER,
					Lexeme:  "foo",
					Literal: "foo",
					Line:    1,
				},
			},
			ExpectedError: nil,
		},
	}

	for _, tc := range tt {
		tc.ExpectedTokens = append(tc.ExpectedTokens, Token{
			Type:    EOF,
			Lexeme:  "",
			Literal: nil,
			Line:    0,
		})
		tokens, errs := Scan(tc.Source)

		assert.Equal(t, len(tc.ExpectedTokens), len(tokens))
		for i, token := range tokens {
			assert.Equal(t, tc.ExpectedTokens[i].Type, token.Type)
			if token.Type != EOF {
				assert.Equal(t, tc.ExpectedTokens[i].Lexeme, token.Lexeme)
				assert.Equal(t, tc.ExpectedTokens[i].Line, token.Line)
			}
		}

		assert.Equal(t, len(tc.ExpectedError), len(errs))
		for i, err := range errs {
			fmt.Println(tc.Source)
			fmt.Println(err)
			assert.Equal(t, tc.ExpectedError[i], err)
		}
	}
}

// TODO: Add test for int, floats, strings
