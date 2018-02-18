package main

import "fmt"

type BinSearch func([]int, int) int

// binary search algorithm on a sorted array
// returns


// does the array have multiple of the same elements?
// It should return the first one? or any of them?
// https://www.geeksforgeeks.org/the-ubiquitous-binary-search-set-1/

// initial implementation - recursive
func binsearch(array []int, target int) int {
	length := len(array)

	// return -1 if the element cannot be found
	if length == 0 {
		return -1
	}

	// take the middle int of the array and compare it
	medianIndex := int(length / 2) // assuming no overflow values since length would already overflow
	median := array[medianIndex]

	if median == target {
		return medianIndex
	}

	// return -1 if the element cannot be found
	if length == 1 {
		return -1
	}

	if median > target {
		return binsearch(array[:medianIndex], target)
	} else {
		// don't need to search the median index again
		return binsearch(array[medianIndex+1:], target)
	}
}


// in-place implementation
func binsearch_inplace(array []int, target int) int {
	r := len(array) - 1
	l := 0

	for r > l {
		medianIndex := l + int((r - l) / 2)
		median := array[medianIndex]
		fmt.Printf("Index: %d, Value: %d\n", medianIndex, median)
		if median == target {
			return medianIndex
		} else if median < target {
			l = medianIndex + 1
		} else if median > target {
			r = medianIndex - 1
		}
	}

	return -1
}