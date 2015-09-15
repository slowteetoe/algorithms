package search

import (
	"testing"
)

func TestValueContainedInOddArray(t *testing.T) {
	a := []int{10, 20, 30, 40, 50}
	key := 20
	result := BinarySearch(key, a)
	if result != 1 {
		t.Errorf("Should have found %v in %v", key, a)
	}
}

func TestValueContainedInEvenArray(t *testing.T) {
	a := []int{10, 20, 30, 40, 50, 60}
	key := 20
	result := BinarySearch(key, a)
	if result != 1 {
		t.Errorf("Should have found %v in %v", key, a)
	}
}

func TestValueNotInArray(t *testing.T) {
	a := []int{10, 20, 30, 40, 50, 60}
	key := 25
	result := BinarySearch(key, a)
	if result != -1 {
		t.Errorf("Should not have found %v in %v", key, a)
	}
}
