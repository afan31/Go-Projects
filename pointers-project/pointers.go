package main

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age // get the address of the value

	// *agePointer to get the value stored in the pointer,
	// deferencing the pointer

	fmt.Println("Age: ", *agePointer)

	getAdultYears(agePointer)
	fmt.Println("Adult Years: ", age)
}

func getAdultYears(age *int) {
	//return *age - 18
	*age = *age - 18 //de-deference the pointer
}
