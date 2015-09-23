package convexhull

import ()

type Point2D struct {
	x, y double
}

// this is useful, given points a,b,c is a->b->c a counterclockwise turn?
// i.e. is c to the left of ray a->b
// return -1 clockwise, 1 ccw or 0 for collinear
func CounterClockWise(a, b, c Point2D) int {
	area := (b.x-a.x)*(c.y-a.y) - (b.y-a.y)*(c.x-a.x)
	if area < 0 {
		return -1
	} else if area > 0 {
		return 1
	} else {
		return 0
	}
}
