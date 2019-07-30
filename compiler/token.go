package compiler

import "fmt"

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%v, %v, %v)", t.Type, t.Lexeme, t.Literal)
}

type TokenType int

const (
	// Literals
	STRING TokenType = iota
	INT
	IDENTIFIER
	FLOAT

	// Single character tokens
	LEFT_PAREN
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	LEFT_SQUARE
	RIGHT_SQUARE
	COMMA
	DOT
	MINUS
	PLUS
	COLON
	SLASH
	STAR

	// 1 - 2 character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL
	PIPE_PIPE
	AMPER_AMPER

	// Keywords
	FUN
	FOR
	IF
	ELSE
	NULL
	TRUE
	FALSE
	PRINT
	RETURN
	VAR
	VAL

	EOF
)

var Keywords = map[string]TokenType{
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"else":   ELSE,
	"null":   NULL,
	"True":   TRUE,
	"False":  FALSE,
	"print":  PRINT,
	"return": RETURN,
	"var":    VAR,
	"val":    VAL,
}
