package ast

import "fmt"

var OperationPrettyPrint = printAst{}

type printAst struct{}

func (pa printAst) ApplyToBinaryNode(n BinaryNode) (interface{}, error) {
	return pa.paren(n.Operator.Lexeme, n.Left, n.Right), nil
}

func (pa printAst) ApplyToGroupingNode(n GroupingNode) (interface{}, error) {
	return pa.paren("group", n.Expression), nil
}

func (pa printAst) ApplyToLiteralNode(n LiteralNode) (interface{}, error) {
	return fmt.Sprintf("%+v", n.Value), nil
}

func (pa printAst) ApplyToUnaryNode(n UnaryNode) (interface{}, error) {
	return pa.paren(n.Operator.Lexeme, n.Right), nil
}

func (pa printAst) paren(name string, nodes ...Node) string {
	output := "("
	output += name
	for _, node := range nodes {
		out, _ := node.Apply(pa)
		output += " " + out.(string)
	}
	output += ")"
	return output
}
