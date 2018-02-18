package main

import (
	"fmt"
	"reflect"
	"runtime"
)


// binsearch, qsort, and btree [insert,querey,delete]
func main() {
	sample_array := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11} // 11 items, missing 8

	run_binsearch(binsearch_inplace, sample_array)
	run_binsearch(binsearch, sample_array)

	run_qsort()
}


func run_qsort() {
	sample_array := []int{6, 10, 11, 16, 2, 4, 8, 32, 1, 0, 3, 7}

	fmt.Printf("Now quicksort. Initial: %v\n", sample_array)
	qsort_bentley(sample_array)

	fmt.Println("Done:\n")
	fmt.Printf("%v\n", sample_array)

}

func run_binsearch(search BinSearch, array []int) {
	fmt.Printf("Now executing: %s\n", getFunctionName(search))
	fmt.Printf("%v\n", array)
	fmt.Println("Searching for 3:")
	fmt.Println(binsearch_inplace(array, 3));
	fmt.Println("Searching for 8:")
	fmt.Println(binsearch_inplace(array, 8));
	fmt.Println("---------------------")
}

// credit StackOverflow
// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}