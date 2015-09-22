package sort

import (
	"reflect"
	"strings"
	"testing"
)

func TestShellSort(t *testing.T) {
	input := strings.Split("shellsortexample", "")
	actual := ShellSort(input)
	expected := strings.Split("aeeehlllmoprsstx", "")
	if reflect.DeepEqual(expected, actual) != true {
		t.Errorf("Expected %v, was %v", expected, actual)
	}
}

var hSizeTests = []struct {
	n, expected int
}{
	{1, 1},
	{2, 1},
	{3, 1},
	{120, 40},
	{200, 121},
	{1000, 364},
}

func TestSizingH(t *testing.T) {
	for _, h := range hSizeTests {
		actual := chooseH(h.n)
		if h.expected != actual {
			t.Errorf("[input=%v] Expected %v, was %v\n", h.n, h.expected, actual)
		}
	}
}
