package unionfind

import (
	"testing"
)

var rootResults = []struct {
	p, expected int
}{
	{0, 2},
	{1, 1},
	{2, 2},
	{3, 3},
}

func TestQuickUnionRootFunc(t *testing.T) {
	uf := NewQuickUnion(4)
	uf.id = []int{2, 1, 2, 3}

	for _, r := range rootResults {
		actual := uf.root(uf.id, r.p)
		if actual != r.expected {
			t.Errorf("Expected root(%v) to be %v, it was %v\n", r.p, r.expected, actual)
		}
	}
}

func TestQuickUnion(t *testing.T) {
	uf := NewQuickUnion(10)
	setupUnionCalls(uf)
	for _, r := range expectedConnectedResults {
		actual := uf.Connected(r.p, r.q)
		if actual != r.expected {
			t.Errorf("Expected Connected(%v, %v) to be %v, but was %v\n", r.p, r.q, r.expected, actual)
		}
	}
}
