package main

import (
	"testing"
)

// TODO think about return here - it's order dependent, so if we switch impl, does it fail?  Maybe switch to asserting via Connected()?
func prepareTest(uf UnionFinder) []int {
	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)
	uf.Union(8, 9)
	uf.Union(5, 0)
	uf.Union(7, 2)
	return []int{0, 1, 1, 8, 8, 0, 0, 1, 8, 8}
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
	expected := prepareTest(uf)
	assertEqual(t, uf.id, expected)
}

func TestQuickFindUsingTable(t *testing.T) {
	uf := NewQuickFind(10)
	prepareTest(uf)
	for _, r := range expectedConnectedResults {
		actual := uf.Connected(r.p, r.q)
		if actual != r.expected {
			t.Errorf("Expected Connected(%v, %v) to be %v, but was %v\n", r.p, r.q, r.expected, actual)
		}
	}
}
