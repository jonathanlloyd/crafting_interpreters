package lexer

import (
	"fmt"
	"strconv"
	"strings"

	gerrors "github.com/jonathanlloyd/crafting_interpreters/golox/internal/errors"
)

type loxLexer struct {
	source  string
	start   int
	current int
	line    int
	tokens  []Token
}

func (l *loxLexer) scanToken() error {
	l.start = l.current

	c := l.advance()
	switch c {
	// Single characters
	case "(":
		l.addToken(LEFT_PAREN)
	case ")":
		l.addToken(RIGHT_PAREN)
	case "{":
		l.addToken(LEFT_BRACE)
	case "}":
		l.addToken(RIGHT_BRACE)
	case ",":
		l.addToken(COMMA)
	case ".":
		l.addToken(DOT)
	case "-":
		l.addToken(MINUS)
	case "+":
		l.addToken(PLUS)
	case ";":
		l.addToken(SEMICOLON)
	case "*":
		l.addToken(STAR)
	// Operators
	case "!":
		if l.match("=") {
			l.addToken(BANG_EQUAL)
		} else {
			l.addToken(BANG)
		}
	case "=":
		if l.match("=") {
			l.addToken(EQUAL_EQUAL)
		} else {
			l.addToken(EQUAL)
		}
	case "<":
		if l.match("=") {
			l.addToken(LESS_EQUAL)
		} else {
			l.addToken(LESS)
		}
	case ">":
		if l.match("=") {
			l.addToken(GREATER_EQUAL)
		} else {
			l.addToken(GREATER)
		}
		// Literals
	case "\"": // string
		err := l.scanString()
		if err != nil {
			return err
		}
	// Ignored lexemes
	case "/":
		if l.match("/") { // A comment!
			for l.peek() != "\n" && !l.isAtEnd() {
				l.advance()
			}
		} else {
			l.addToken(SLASH)
		}
	case " ", "\r", "\t": // Ignore whitespace
	case "\n":
		l.line += 1
	// Unexpected character
	default:
		if isDigit(c) {
			err := l.scanNumber()
			if err != nil {
				return err
			}
		} else if isAlpha(c) {
			err := l.scanIdentifierOrKeyword()
			if err != nil {
				return err
			}
		} else {
			return gerrors.GoloxError{
				Line:    l.line,
				Message: fmt.Sprintf("Unexpected character: %q", c),
			}
		}
	}

	return nil
}

func (l *loxLexer) scanString() error {
	for l.peek() != "\"" && !l.isAtEnd() {
		if l.peek() == "\n" {
			l.line += 1
		}
		l.advance()
	}

	if l.isAtEnd() {
		return gerrors.GoloxError{
			Line:    l.line,
			Message: "Unterminated string",
		}
	}

	l.advance()

	stringLiteral := l.source[l.start+1 : l.current-1]
	l.addTokenWithLiteral(STRING, stringLiteral)

	return nil
}

func (l *loxLexer) scanNumber() error {
	for isDigit(l.peek()) {
		l.advance()
	}

	// Consume digits after a dot for fractionals E.g. 1.22
	if l.peek() == "." && isDigit(l.peekNext()) {
		l.advance()
		for isDigit(l.peek()) {
			l.advance()
		}
	}

	numberLiteral, err := strconv.ParseFloat(
		l.source[l.start:l.current],
		64,
	)
	if err != nil {
		panic(err)
	}

	l.addTokenWithLiteral(NUMBER, numberLiteral)

	return nil
}

func (l *loxLexer) scanIdentifierOrKeyword() error {
	for isAlphanumeric(l.peek()) {
		l.advance()
	}

	lexeme := l.source[l.start:l.current]
	tokenType, ok := Keywords[lexeme]
	if ok {
		l.addToken(tokenType)
	} else {
		l.addToken(IDENTIFIER)
	}

	return nil
}

func (l *loxLexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *loxLexer) advance() string {
	l.current += 1
	return l.source[l.current-1 : l.current]
}

func (l *loxLexer) match(next string) bool {
	if l.isAtEnd() {
		return false
	}

	c := l.source[l.current : l.current+1]
	if c != next {
		return false
	}

	l.current += 1
	return true
}

func (l *loxLexer) peek() string {
	if l.isAtEnd() {
		return "\x00"
	} else {
		return l.source[l.current : l.current+1]
	}
}

func (l *loxLexer) peekNext() string {
	if l.current+1 >= len(l.source) {
		return "\x00"
	} else {
		return l.source[l.current+1 : l.current+2]
	}
}

func (l *loxLexer) addToken(tokenType TokenType) {
	l.addTokenWithLiteral(tokenType, nil)
}

func (l *loxLexer) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	l.tokens = append(l.tokens, Token{
		Type:    tokenType,
		Lexeme:  l.source[l.start:l.current],
		Literal: literal,
		Line:    l.line,
	})
}

func isDigit(c string) bool {
	return len(c) == 1 &&
		strings.Compare("0", c) < 1 &&
		strings.Compare(c, "9") < 1
}

func isAlpha(c string) bool {
	return len(c) == 1 &&
		((strings.Compare("a", c) < 1 && strings.Compare(c, "z") < 1) ||
			(strings.Compare("A", c) < 1 && strings.Compare(c, "Z") < 1) ||
			c == "_")
}

func isAlphanumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func Scan(source string) (tokens []Token, err error) {
	l := loxLexer{source: source}
	for !l.isAtEnd() {
		if err := l.scanToken(); err != nil {
			return nil, err
		}
	}
	return l.tokens, nil
}
