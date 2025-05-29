package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}
	//sum := sumup(numbers)
	sum := sumup(1, 10, 15, -5, 10)
	anotherSum := sumup(1, numbers...) // now called with 4 params,
	// starting value, and 3 values being pulled out from the slice
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// variadic functions : you can write functions with any amount of parameters

// func sumup(numbers []int) int {
// 	sum := 0

// 	for _, val := range numbers {
// 		sum += val // sum = sum + val
// 	}

//		return sum
//	}
//
// call with as many params as you want, as long as every param is int
func sumup(startingValue int /*optional*/, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val // sum = sum + val
	}

	return sum
}
