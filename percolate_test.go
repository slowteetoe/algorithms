package main

import (
	"testing"
)

func TestPercolation(t *testing.T) {
	g := NewGrid(4)
	g.Display()
	g.Open(9)
	g.Display()
	g.Open(8)
	g.Display()
	g.Open(5)
	g.Display()
	g.Open(11)
	g.Display()
	g.Open(14)
	g.Display()
	g.Open(1)
	g.Display()
	g.Open(15)
	g.Display()
	g.Open(10)
	g.Display()
}
