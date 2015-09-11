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
	cells    []int8
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

func (g *grid) IsClosed(p int) bool {
	return g.cells[p] == 0
}

func (g *grid) Open(p int) {
	g.cells[p] = 1
	// check to see if any adjacent sites need to be joined
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
}

func (g *grid) Percolates() bool {
	return g.uf.Connected(g.top, g.bot)
}

func NewGrid(initialSize int) grid {
	numCells := initialSize * initialSize
	// we have two "virtual sites" to reduce the number of connected() calls we need to make
	grid := grid{n: initialSize, uf: unionfind.NewWeightedQuickUnion(numCells + 2), top: numCells, bot: numCells + 1}
	grid.cells = make([]int8, numCells)

	// initialize the top and bottom rows with links to both virtual sites
	for i := 0; i < grid.n; i++ {
		grid.uf.Union(i, grid.top)
		grid.uf.Union((numCells-1)-i, grid.bot)
	}
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
	n := 800
	_, i := MonteCarlo(n)
	percentage := float64(i) / float64(n*n)
	fmt.Printf("%vx%v grid percolated after %v cells were opened: %f\n", n, n, i, percentage)
}
