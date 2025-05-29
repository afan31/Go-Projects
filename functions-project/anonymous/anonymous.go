package anonymous

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor // if you use a variable from an outer scope,
		//such anonymous functions are called closures
	}
}
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// when a function wants another function as a param value or a function
// returns a function, you can save time by using anonymous functions
// which is a feature that allows you to define a function just in time when you need it
// instead in advance

//closures also use anonymous functions, but use a specific aspect of anonymous functions
//
