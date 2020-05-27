package errors

import "fmt"

type GoloxError struct {
	Line    int
	Message string
}

func (g GoloxError) Error() string {
	return fmt.Sprintf("Line %d: %s", g.Line, g.Message)
}
