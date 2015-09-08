package main

import ()

type UnionFind interface {
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

type quickUnion struct {
	id []int
}

func NewQuickUnion(n int) *quickUnion {
	u := new(quickUnion)
	u.id = make([]int, n)
	for i := 0; i < n; i++ {
		u.id[i] = i
	}
	return u
}

func (qu quickUnion) Union(p, q int) {
	proot := qu.root(qu.id, p)
	qroot := qu.root(qu.id, q)
	qu.id[proot] = qu.id[qroot]
}

func (qu quickUnion) Connected(p, q int) bool {
	return qu.root(qu.id, p) == qu.root(qu.id, q)
}

func (qu quickUnion) root(arr []int, p int) int {
	parent := arr[p]
	for arr[parent] != parent {
		parent = arr[parent]
	}
	return parent
}
