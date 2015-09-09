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
	proot := root(qu.id, p)
	qroot := root(qu.id, q)
	qu.id[proot] = qu.id[qroot]
}

func (qu quickUnion) Connected(p, q int) bool {
	return root(qu.id, p) == root(qu.id, q)
}

func root(arr []int, p int) int {
	parent := arr[p]
	for arr[parent] != parent {
		parent = arr[parent]
	}
	return parent
}

type weightedQuickUnion struct {
	id, sz []int
}

func NewWeightedQuickUnion(n int) *weightedQuickUnion {
	w := new(weightedQuickUnion)
	w.id = make([]int, n)
	w.sz = make([]int, n)
	for i := 0; i < n; i++ {
		w.id[i] = i
		w.sz[i] = 1
	}
	return w
}

func (w weightedQuickUnion) count(p int) int {
	// shouldn't need to ensure this is the root, since it's private
	return w.sz[p]
}

func (w weightedQuickUnion) Union(p, q int) {
	proot := root(w.id, p)
	qroot := root(w.id, q)
	// only difference from non-weighted quick-union is that the smaller tree needs to be appended to the root of the larger tree
	combined := w.sz[qroot] + w.sz[proot]
	if w.count(proot) > w.count(qroot) {
		w.id[proot] = w.id[qroot]
		w.sz[proot] = combined
	} else {
		w.id[qroot] = w.id[proot]
		w.sz[qroot] = combined
	}
}

func (w weightedQuickUnion) Connected(p, q int) bool {
	return root(w.id, p) == root(w.id, q)
}
