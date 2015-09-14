package main

import (
	"fmt"
	"github.com/slowteetoe/algorithms/unionfind"
	"log"
	"math/rand"
	"time"
)

type grid struct {
	n        int
	uf       unionfind.UnionFind
	backwash unionfind.UnionFind
	cells    []byte
	top, bot int
}

func (g *grid) Display() {
	for i, val := range g.cells {
		if i%g.n == 0 && i != 0 {
			fmt.Println("")
		}
		if val == 0 {
			fmt.Print("*")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Printf("\nPercolates? : %v\n\n", g.Percolates())
}

func (g *grid) IsOpen(p int) bool {
	return g.cells[p] == 1
}

func (g *grid) IsFull(p int) bool {
	return g.backwash.Connected(g.top, p)
}

func (g *grid) Open(p int) {
	g.cells[p] = 1
	// check to see if any adjacent sites need to be joined

	// see if we need to join to top virtual site
	if p < g.n {
		g.uf.Union(p, g.top)
		g.backwash.Union(p, g.top)
	}

	if p-1 >= 0 && g.IsOpen(p-1) {
		g.uf.Union(p, p-1)
	}
	if p+1 < g.n*g.n && g.IsOpen(p+1) {
		g.uf.Union(p, p+1)
	}
	if p-g.n >= 0 && g.IsOpen(p-g.n) {
		g.uf.Union(p, p-g.n)
	}
	if p+g.n < g.n*g.n && g.IsOpen(p+g.n) {
		g.uf.Union(p, p+g.n)
	}

	// see if we need to join to bot virtual site, backwash does not have a bot virtual site
	if p > len(g.cells)-g.n {
		g.uf.Union(p, g.bot)
	}

}

func (g *grid) Percolates() bool {
	return g.uf.Connected(g.top, g.bot)
}

func NewGrid(initialSize int) grid {
	numCells := initialSize * initialSize
	// we have two "virtual sites" to reduce the number of connected() calls we need to make
	grid := grid{n: initialSize, uf: unionfind.NewWeightedQuickUnion(numCells + 2), backwash: unionfind.NewWeightedQuickUnion(numCells + 1), top: numCells, bot: numCells + 1}
	grid.cells = make([]byte, numCells)

	return grid
}

// Run percolation, returning the percentage of open sites it took to percolate
func MonteCarlo(n int) (grid, int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	g := NewGrid(n)
	cells := r.Perm(n * n)
	for i, cell := range cells {
		g.Open(cell)
		if g.Percolates() {
			log.Printf("Percolated after %v cells were opened.\n", i)
			return g, i
		}
	}
	panic("Never percolated")
}

func main() {
	for _, n := range []int{200, 400, 800, 1600, 3200} {
		start := time.Now()
		_, i := MonteCarlo(n)
		elapsed := time.Since(start)
		percentage := float64(i) / float64(n*n)
		fmt.Printf("(%v) %vx%v grid percolated after %v cells were opened: %f\n", elapsed, n, n, i, percentage)
	}
}
