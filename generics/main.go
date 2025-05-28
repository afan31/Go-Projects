package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

// any or interface{}
// what we are doing with this generic type placeholder
// is we are telling Go that the concrete type of values that
// we are returning here and receiving will only be known at the
// point of time when add is called (so line 4 in this case)
// here T can be any type as long as it is int, float or a string
func add[T int | float64 | string](a, b T) T {
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)

	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)

	// if aIsFloat && bIsFloat {
	// 	return aFloat + bFloat
	// }

	// aString, aIsString := a.(string)
	// bString, bIsString := b.(string)

	// if aIsString && bIsString {
	// 	return aString + bString
	// }

	// return nil

	return a + b
}
