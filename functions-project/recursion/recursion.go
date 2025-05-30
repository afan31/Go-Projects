package recursion

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

}

// factorial of 5 => 5 * 4 * 3 * 2 * 1
func factorial(number int) int {
	// recursive approach - function calls itself
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)

	// result := 1

	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }

	// return result
}
