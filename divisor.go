package main

import "fmt"

func isDivisor(numerator int, denominator int) bool {
	return numerator % denominator == 0
}

func sumPotentialDivisor(num int, divisor int, c chan int) {
	if isDivisor(num, divisor) {
		c <- divisor
	} else {
		c <- 0
	}
}

func divisorSumConcurrent(num int) int {

	sum := 0

	c := make(chan int)

	for i := 1; i < num; i++ {
		go sumPotentialDivisor(num, i, c)
		sum = sum +<- c
	}

	return sum
}

func divisorSum(num int) int {
	sum := 0

	for i := 1; i < num; i++ {
	//	sumPotentialDivisor(num, i, &sum)
	}

	return sum
}

func divisorGreater(num int) {
	less := "deficient"
	greater := "abundant"
	eq := "equals"

	//divisorSum := divisorSum(num)
	divisorSum := divisorSumConcurrent(num)

	fmt.Printf("Is %d a perfect number?\n", num)
	if divisorSum < num {
		fmt.Println(less)
	} else if divisorSum > num {
		fmt.Println(greater)
	} else {
		fmt.Println(eq)
	}
}