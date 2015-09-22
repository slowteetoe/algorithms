package sort

import (
)

func ShellSort(arr []string) []string {
	n := len(arr)
	h := chooseH(n)
	for h >= 1 {
		// use insertion sort
		for i := h; i < n; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				tmp := arr[j]
				arr[j] = arr[j-h]
				arr[j-h] = tmp
			}
		}
		h = h / 3
	}
	return arr
}

func chooseH(n int) int {
	h := 1
	for h < n / 3 {
		h = 3*h + 1
	}
	return h
}
