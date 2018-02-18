package main

import "fmt"

// quicksort is a difficult beast
// there's that bizarre algorithm, jon bentley


// idea: choose a pivot and move everything to be above
// and below the pivot

// naive extra memory implementation
func qsort_naive(array []int) {
	temp := make([]int, len(array))
	copy(temp, array)
	// randomly choose the middle element as the pivot
	pivotIndex := int(len(array) / 2)
	pivot := array[pivotIndex]

	for _, i := range temp {
		l := 0
		if i > pivot {
			array[l]  = i
			l++
		}
	}

	// wait, I don't even know how this works
	// don't know where the pivot is going to end up in the list
	// I guess that's what bentley's sort is so cool being in-place
}


func qsort_bentley(array []int) {
	_qsort_bentley(array, 0, len(array) - 1)
}

func _qsort_bentley(array []int, l int, r int) {
	pivot := &array[l]

	if l >= r {
		return
	}

	m := l // choose an element to be pivot
	for i := l+1; i <= r ; i++ {
/*
		if array[i] < pivot {
			m++
			array[i], array[m] = array[m], array[i]
		}
*/
// try out some pointer arithmetic
		element := &array[i]
		if *element < *pivot {
			m++
			*element, array[m] = array[m], *element
		}
	}

	array[m], *pivot = *pivot, array[m]

	fmt.Printf("Current: %v\n", array)

	_qsort_bentley(array, l, m-1)
	_qsort_bentley(array, m+1, r)
}
