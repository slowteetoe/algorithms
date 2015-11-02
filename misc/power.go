package misc

import ()

// "Round up to the next highest power of 2", from http://graphics.stanford.edu/~seander/bithacks.html
func RoundUpToPowerOf2(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}
