package scanner

import (
	"fmt"
	"strconv"
	"unicode"
)

// Scanner scans a source text and turns it into a series of tokens.
type Scanner struct {
	source []rune

	// Tokens is the series of scanned tokens from the source text
	Tokens chan Token
	Errors chan error
	Done   chan bool

	current int
	line    int
}

// Scan will scan a source and return the tokens and errors it found
func Scan(source string) (tokens []Token, err []error) {
	s := NewScanner(source)
	go s.Scan()

Loop:
	for {
		select {
		case t := <-s.Tokens:
			tokens = append(tokens, t)
		case e := <-s.Errors:
			err = append(err, e)
		case <-s.Done:
			break Loop
		}
	}
	return
}

// NewScanner creates a new scanner from a source string
// TODO: Make this take in a Reader instead of a string
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: []rune(source),
		Tokens: make(chan Token),
		Errors: make(chan error),
		Done:   make(chan bool),
	}
}

// Scan starts the scanner. It will only produce results once
func (s *Scanner) Scan() {
	defer close(s.Tokens)
	defer close(s.Errors)
	defer close(s.Done)
	defer func() { s.Done <- true }()

	for !s.atEnd() {
		s.match()
	}

	s.Tokens <- Token{
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    s.line,
	}
}

func (s *Scanner) match() {
	r, done := s.advance()
	if done {
		return
	}
	switch {
	case unicode.IsSpace(r):
		if r == '\n' {
			s.line++
		}
	case r == '(':
		s.singleChar(LEFT_PAREN, r)
	case r == ')':
		s.singleChar(RIGHT_PAREN, r)
	case r == '[':
		s.singleChar(LEFT_SQUARE, r)
	case r == ']':
		s.singleChar(RIGHT_SQUARE, r)
	case r == '{':
		s.singleChar(LEFT_BRACE, r)
	case r == '}':
		s.singleChar(RIGHT_BRACE, r)
	case r == ',':
		s.singleChar(COMMA, r)
	case r == '.':
		s.singleChar(DOT, r)
	case r == '-':
		s.singleChar(MINUS, r)
	case r == '+':
		s.singleChar(PLUS, r)
	case r == '*':
		s.singleChar(STAR, r)
	case r == ':':
		s.singleChar(COLON, r)
	case r == '/':
		next, _ := s.peek()
		switch {
		case next == '/':
			s.lineComment(next)
		case next == '*':
			s.blockComment(next)
		default:
			s.singleChar(SLASH, r)
		}
	case r == '<':
		next, _ := s.peek()
		switch {
		case next == '=':
			s.doubleChar(LESS_EQUAL, r)
		default:
			s.singleChar(LESS, r)
		}
	case r == '>':
		next, _ := s.peek()
		switch {
		case next == '=':
			s.doubleChar(GREATER_EQUAL, r)
		default:
			s.singleChar(GREATER, r)
		}
	case r == '!':
		next, _ := s.peek()
		switch {
		case next == '=':
			s.doubleChar(BANG_EQUAL, r)
		default:
			s.singleChar(BANG, r)
		}
	case r == '=':
		next, _ := s.peek()
		switch {
		case next == '=':
			s.doubleChar(EQUAL_EQUAL, r)
		default:
			s.singleChar(EQUAL, r)
		}
	case r == '&':
		next, _ := s.peek()
		if next == '&' {
			s.doubleChar(AMPER_AMPER, r)
		} else {
			s.Errors <- fmt.Errorf("unrecognized token: found single char &")
		}
	case r == '|':
		next, _ := s.peek()
		if next == '|' {
			s.doubleChar(PIPE_PIPE, r)
		} else {
			s.Errors <- fmt.Errorf("unrecognized token: found single char |")
		}
	case r == '"':
		s.handleString(r)
	case unicode.IsDigit(r):
		s.handleDigits(r)
	case unicode.IsLetter(r):
		s.handleLetters(r)
	default:
		s.Errors <- fmt.Errorf("unknown character %c", r)
		s.advance()
	}
}

func (s *Scanner) blockComment(next rune) {
	foundStar := false
Loop:
	for ; !s.atEnd(); next, _ = s.advance() {

		switch {
		case next == '\n':
			foundStar = false
			s.line++
		case next == '*':
			foundStar = true

		case next == '/':
			if foundStar {
				break Loop
			}
		default:
			foundStar = false
		}
	}
}

func (s *Scanner) lineComment(next rune) {
	for ; next != '\n' && !s.atEnd(); next, _ = s.advance() {
	}
	s.line++
}

func (s *Scanner) singleChar(t TokenType, lit rune) {
	s.Tokens <- Token{
		Type:    t,
		Lexeme:  string(lit),
		Literal: nil,
		Line:    s.line,
	}
}

func (s *Scanner) doubleChar(t TokenType, lit rune) {
	second, _ := s.advance()
	s.Tokens <- Token{
		Type:    t,
		Lexeme:  string([]rune{lit, second}),
		Literal: nil,
		Line:    s.line,
	}
}

func (s *Scanner) atEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() (rune, bool) {
	r, done := s.peek()
	s.current++
	return r, done
}

func (s *Scanner) peek() (rune, bool) {
	if s.atEnd() {
		return 0, true
	}
	return s.source[s.current], false
}

func (s *Scanner) handleDigits(r rune) {
	lex := []rune{r}
	isfloat := false

Loop:
	for !s.atEnd() {
		switch r, _ := s.advance(); {
		case unicode.IsDigit(r):
			lex = append(lex, r)
		case r == '.':
			if isfloat {
				lex = append(lex, r)
				s.Errors <- fmt.Errorf("line: %v, too many '.'s found in float, lex: %v", s.line, string(lex))
				return
			}
			isfloat = true
			lex = append(lex, r)
		default:
			s.current--
			break Loop
		}
	}
	if isfloat {
		lit, err := strconv.ParseFloat(string(lex), 64)
		if err != nil {
			s.Errors <- err
			return
		}
		s.Tokens <- Token{
			Type:    FLOAT,
			Lexeme:  string(lex),
			Literal: lit,
			Line:    s.line,
		}
	} else {
		lit, err := strconv.ParseInt(string(lex), 10, 64)
		if err != nil {
			s.Errors <- err
			return
		}
		s.Tokens <- Token{
			Type:    INT,
			Lexeme:  string(lex),
			Literal: lit,
			Line:    s.line,
		}
	}
}

func (s *Scanner) handleLetters(r rune) {
	lex := []rune{r}
Loop:
	for !s.atEnd() {
		switch r, _ := s.advance(); {
		case unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_':
			lex = append(lex, r)
		default:
			s.current--
			break Loop
		}
	}
	if token, ok := Keywords[string(lex)]; ok {
		s.Tokens <- Token{
			Type:    token,
			Lexeme:  string(lex),
			Literal: nil,
			Line:    s.line,
		}
	} else {
		s.Tokens <- Token{
			Type:    IDENTIFIER,
			Lexeme:  string(lex),
			Literal: string(lex),
			Line:    s.line,
		}
	}
}

func (s *Scanner) handleString(r rune) {
	lex := []rune{r}
	lit := []rune{}
	foundEnd := false
Loop:
	for !s.atEnd() {
		switch r, _ := s.advance(); {
		case r == '"':
			lex = append(lex, r)
			foundEnd = true
			break Loop
		case r == '\n':
			break
		default:
			lex = append(lex, r)
			lit = append(lit, r)
		}
	}
	if foundEnd {
		s.Tokens <- Token{
			Type:    STRING,
			Lexeme:  string(lex),
			Literal: string(lit),
			Line:    s.line,
		}
	} else {
		s.Errors <- fmt.Errorf("closing \" of string not found")
	}
}
