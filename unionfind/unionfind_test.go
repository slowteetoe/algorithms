package unionfind

import (
	"fmt"
	"math"
	"testing"
)

func setupUnionCalls(uf UnionFind) {
	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)
	uf.Union(8, 9)
	uf.Union(5, 0)
	uf.Union(7, 2)
}

var expectedConnectedResults = []struct {
	p, q     int
	expected bool
}{
	{0, 1, false},
	{0, 5, true},
	{1, 6, false},
	{6, 1, false},
	{8, 9, true},
	{4, 9, true},
}

func assertEqual(t *testing.T, actual, expected []int) {
	for i, _ := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("expected %v but got %v", expected, actual)
		}
	}
}

func TestOneStep(t *testing.T) {
	uf := NewQuickFind(10)
	uf.Union(0, 1)
	assertEqual(t, uf.id, []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestQuickFind(t *testing.T) {
	uf := NewQuickFind(10)
	setupUnionCalls(uf)
	for _, r := range expectedConnectedResults {
		actual := uf.Connected(r.p, r.q)
		if actual != r.expected {
			t.Errorf("Expected Connected(%v, %v) to be %v, but was %v\n", r.p, r.q, r.expected, actual)
		}
	}
}

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
		actual := root(uf.id, r.p)
		if actual != r.expected {
			t.Errorf("Expected root(%v) to be %v, it was %v\n", r.p, r.expected, actual)
		}
	}
}

func TestUnionFind(t *testing.T) {
	uf := NewQuickUnion(10)
	setupUnionCalls(uf)
	for _, r := range expectedConnectedResults {
		actual := uf.Connected(r.p, r.q)
		if actual != r.expected {
			t.Errorf("Expected Connected(%v, %v) to be %v, but was %v\n", r.p, r.q, r.expected, actual)
		}
	}
}

func TestWeightedUnionFind(t *testing.T) {
	w := NewWeightedQuickUnion(10)
	setupUnionCalls(w)
	for _, r := range expectedConnectedResults {
		actual := w.Connected(r.p, r.q)
		if actual != r.expected {
			t.Errorf("Expected Connected(%v, %v) to be %v, but was %v\n", r.p, r.q, r.expected, actual)
		}
	}
	fmt.Printf("id:%v\nsz:%v\n", w.id, w.sz)
}

func TestTreeSizeWeightedQU(t *testing.T){
	n := 8
	w := NewWeightedQuickUnion(n)
	for i := 1; i < n; i++ {
		w.Union(i,0)
	}
	log2n := int(math.Log2(float64(n)))
	for _, val := range w.sz {
		if val > log2n {
			t.Errorf("No depth should be greater than log2(%v), which is %v.  This one was: %v", n, log2n, val)
		}
	}
	// fmt.Printf("%v", w.sz)
}
