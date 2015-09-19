package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

type stack struct{ vec []string }

func (s stack) Empty() bool      { return len(s.vec) == 0 }
func (s *stack) Push(str string) { s.vec = append(s.vec, str) }
func (s *stack) Pop() string {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}

func toInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Errorf("Unable to handle: %v", s)
	}
	return val
}

func performOp(v1, v2, op string) string {
	p := toInt(v1)
	q := toInt(v2)
	result := 0
	switch op {
	case "+":
		result = p + q
	case "-":
		result = p - q
	case "/":
		result = p / q
	case "*":
		result = p * q
	}
	return strconv.Itoa(result)
}

// use djikstra's two-stack algorithm
func Evaluate(expr string) int {
	values := stack{}
	operators := stack{}
	operands := strings.Split(expr, " ")
	for _, s := range operands {
		switch s {
		default:
			values.Push(s)
		case "(":
			continue
		case "+", "-", "/", "*":
			operators.Push(s)
		case ")":
			values.Push(performOp(values.Pop(), values.Pop(), operators.Pop()))
		}
	}
	result, err := strconv.Atoi(values.Pop())
	if err != nil {
		panic(fmt.Sprintf("Could not process, malformed expression? %v", err))
	}
	return result
}
