package main

import (
	"fmt"
	"reflect"
	"runtime"
)


// binsearch, qsort, and btree [insert,querey,delete]
func main() {

	divisorGreater(2)
	divisorGreater(1)
	divisorGreater(8128)
	divisorGreater(0)

	/*
	run_btree()

	run_binsearch(binsearch_inplace)
	run_binsearch(binsearch)

	run_qsort()
	*/
}

func run_btree() {
	tree := BTree{5, &BTree{2, &BTree{1, nil, nil}, &BTree{4, nil, nil}},
		&BTree{7, &BTree{6, nil, nil}, &BTree{9, &BTree{8, nil, nil}, nil}}}

	run_btree_query(tree)

	run_btree_insert(tree)

	run_btree_delete(tree)

}

func run_btree_delete(tree BTree) {
	fmt.Println("---------------------------------------")
	fmt.Println("Starting BTree delete. Deleting 8")

	deleteTree(&tree, 8);

	tree.print()

	fmt.Printf("Does 8 exist now? %b. Should be false\n", query(&tree, 8))

	fmt.Println("Now deleting 2")

	deleteTree(&tree, 2);

	tree.print()

}

func run_btree_query(tree BTree) {
	fmt.Println("-----------------------------")
	fmt.Println("Starting BTree query. 10 should not exist")
	fmt.Println(query(&tree, 10));

	fmt.Println("Starting BTree query. 4 should exist")
	fmt.Println(query(&tree, 4));
}

func run_btree_insert(tree BTree) {
	fmt.Println("-----------------------------")
	fmt.Println("Now executing btree insert")

	fmt.Println("Old tree")
	tree.print()

	insert(&tree, 3);

	fmt.Println("New tree")
	tree.print()
}

func run_qsort() {
	sample_array := []int{6, 10, 11, 16, 2, 4, 8, 32, 1, 0, 3, 7}

	fmt.Printf("Now quicksort. Initial: %v\n", sample_array)
	qsort_bentley(sample_array)

	fmt.Println("Done:\n")
	fmt.Printf("%v\n", sample_array)

}

func run_binsearch(search BinSearch) {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 9, 10, 11} // 11 items, missing 8

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