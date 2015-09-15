package search

import ()

// Yes, it is better to use sort.Search() - see http://golang.org/pkg/sort/#Search for how it was implemented
// For fun, read about how Java's was broken until 2006: http://googleresearch.blogspot.com/2006/06/extra-extra-read-all-about-it-nearly.html
// Return the index of i in the array, if it contains i. -1 otherwise
func BinarySearch(i int, data []int) int {
	lo := 0
	hi := len(data) - 1
	for lo <= hi {
		mid := lo + (hi-lo)/2 // (low+hi) / 2 can overflow
		if i < data[mid] {
			hi = mid - 1
		} else if i > data[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
