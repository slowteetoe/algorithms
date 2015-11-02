package misc

import (
	"testing"
)

var data = []struct {
	input, expected uint64
}{
	{1, 1},
	{2, 2},
	{3, 4},
	{62, 64},
	{1020, 1024},
        {2047, 2048},
	{2048, 2048},
	{2049, 4096},
}

func TestInputs(t *testing.T) {
	for _, testcase := range data {
		actual := RoundUpToPowerOf2(testcase.input)
		if actual != testcase.expected {
			t.Errorf("Expected %v for an input of %v, got %v", testcase.expected, testcase.input, actual)
		}
	}
}
