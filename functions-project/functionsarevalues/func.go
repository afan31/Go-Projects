package functionsarevalues

import "fmt"

// store function types as custom types
type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}

	moreNumbers := []int{5, 1, 2}
	// functions are just values in Go, so they can passed in fucntions
	doubled := transformNumbers(&numbers, double) //just use the name
	// of the func, so that it is can be passed as a value,
	// and not executed
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	transformerFn1 := getTransformerFunction(&moreNumbers)
	//transformerFn2 := getTransformerFunction(&moreNumbers)

	transformNumbers := transformNumbers(&numbers, transformerFn1)
	//moreTransformNumbers := transformNumbers(&moreNumbers, transformerFn2)

	fmt.Println(transformNumbers)
	//fmt.Println(moreTransformNumbers)
}

// the value that we recieve for this transform parameter
// should be a function that has one parameter of type int, and
// returns one value of the type int
func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
