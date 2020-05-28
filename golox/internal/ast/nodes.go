package ast

import "github.com/jonathanlloyd/crafting_interpreters/golox/internal/lexer"

type Node interface {
	Apply(operation Operation) (interface{}, error)
}

type Operation interface {
	ApplyToBinaryNode(n BinaryNode) (interface{}, error)
	ApplyToGroupingNode(n GroupingNode) (interface{}, error)
	ApplyToLiteralNode(n LiteralNode) (interface{}, error)
	ApplyToUnaryNode(n UnaryNode) (interface{}, error)
}

type BinaryNode struct {
	Left     Node
	Operator lexer.Token
	Right    Node
}

func (b BinaryNode) Apply(operation Operation) (interface{}, error) {
	return operation.ApplyToBinaryNode(b)
}

type GroupingNode struct {
	Expression Node
}

func (g GroupingNode) Apply(operation Operation) (interface{}, error) {
	return operation.ApplyToGroupingNode(g)
}

type LiteralNode struct {
	Value interface{}
}

func (l LiteralNode) Apply(operation Operation) (interface{}, error) {
	return operation.ApplyToLiteralNode(l)
}

type UnaryNode struct {
	Operator lexer.Token
	Right    Node
}

func (u UnaryNode) Apply(operation Operation) (interface{}, error) {
	return operation.ApplyToUnaryNode(u)
}
