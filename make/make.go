package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//userNames := []string{} // if you dont assign a length,
	// its not an array but a slice, and thus you can
	// dynamically add elements

	userNames := make([]string, 2, 5) // go allocates enough memory space for

	userNames[0] = "Julie"

	// a string array that takes 5 items, and it wll create an array with 2 memory slots
	// pre-allocating memory for slots make it efficient
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	//courseRatings := map[string]float64{}

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 5.1

	//fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
