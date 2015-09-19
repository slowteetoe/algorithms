package calculator

import (
	"testing"
)

func TestValidExpression(t *testing.T) {
	expr := "( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )"
	expected := 101
	result := Evaluate(expr)
	if result != expected {
		t.Errorf("Evaluating %v resulted in %v, should have been %v", expr, result, expected)
	}
}
