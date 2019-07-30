package compiler

import (
	"fmt"
)

type parser struct {
	source  []Token
	current int

	ast    expr
	Errors []error
}

func Parse() (Expr, errs []error) {

}

type expr interface {
	String() string
}

type binary struct {
	left  expr
	op    Token
	right expr
}

type grouping struct {
	expr expr
}

type literal struct {
	val interface{}
}

type unary struct {
	op    Token
	right expr
}

func (b binary) String() string {
	return "( " + b.op.Lexeme + " " + b.left.String() + " " + b.right.String() + " )"
}

func (g grouping) String() string {
	return "( " + g.expr.String() + " )"
}

func (u unary) String() string {
	return "( " + u.op.Lexeme + " " + u.right.String() + " )"
}

func (l literal) String() string {
	return fmt.Sprintf("%v", l.val)
}
