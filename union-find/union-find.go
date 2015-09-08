package main

import (
	"fmt"
)

type UnionFinder interface {
	Union(p, q int)
	Connected(p, q int) bool
}

type quickFind struct {
	id []int
}

func NewQuickFind(n int) *quickFind {
	q := new(quickFind)
	q.id = make([]int, n)
	for i := 0; i < n; i++ {
		q.id[i] = i
	}
	return q
}

func (qf quickFind) Union(p, q int) {
	idp := qf.id[p]
	idq := qf.id[q]
	for i, val := range qf.id {
		if val == idp {
			qf.id[i] = idq
		}
	}
}

func (qf quickFind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func main() {
	fmt.Println("Todo")
}
