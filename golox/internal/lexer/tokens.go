package lexer

import "fmt"

type TokenType int

const (
	// Single character lexemes
	LEFT_PAREN TokenType = iota
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

	// Single/double character lexemes
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
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

	// EOF
	EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int32
}

func (t Token) String() string {
	var typeStr string
	switch t.Type {
	case LEFT_PAREN:
		typeStr = "LEFT_PAREN"
	case RIGHT_PAREN:
		typeStr = "RIGHT_PAREN"
	case LEFT_BRACE:
		typeStr = "LEFT_BRACE"
	case RIGHT_BRACE:
		typeStr = "RIGHT_BRACE"
	case COMMA:
		typeStr = "COMMA"
	case DOT:
		typeStr = "DOT"
	case MINUS:
		typeStr = "MINUS"
	case PLUS:
		typeStr = "PLUS"
	case SEMICOLON:
		typeStr = "SEMICOLON"
	case SLASH:
		typeStr = "SLASH"
	case STAR:
		typeStr = "STAR"
	case IDENTIFIER:
		typeStr = "IDENTIFIER"
	case STRING:
		typeStr = "STRING"
	case NUMBER:
		typeStr = "NUMBER"
	case AND:
		typeStr = "AND"
	case CLASS:
		typeStr = "CLASS"
	case ELSE:
		typeStr = "ELSE"
	case FALSE:
		typeStr = "FALSE"
	case FUN:
		typeStr = "FUN"
	case FOR:
		typeStr = "FOR"
	case IF:
		typeStr = "IF"
	case NIL:
		typeStr = "NIL"
	case OR:
		typeStr = "OR"
	case PRINT:
		typeStr = "PRINT"
	case RETURN:
		typeStr = "RETURN"
	case SUPER:
		typeStr = "SUPER"
	case THIS:
		typeStr = "THIS"
	case TRUE:
		typeStr = "TRUE"
	case VAR:
		typeStr = "VAR"
	case WHILE:
		typeStr = "WHILE"
	case EOF:
		typeStr = "EOF"
	default:
		typeStr = "UNKNOWN"
	}

	return fmt.Sprintf("Token{Type:\"%s\", Line: %d, Lexeme: %s}", typeStr, t.Line, t.Lexeme)
}
