package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	//1)
	hobbies := [3]string{"Sports", "Cooking", "Reading"}

	//2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4)
	fmt.Println(cap(hobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5)
	courseGoals := []string{"Learn Go!", "Learn all the basics!"}
	fmt.Println(courseGoals)

	//6)
	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basics!")
	fmt.Println(courseGoals)

	//7)
	products := []Product{
		{
			"first-product",
			"A first product",
			12.99,
		},
		{
			"second-product",
			"A second product",
			129.99,
		},
	}
	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A third product",
		199.99,
	}

	products = append(products, newProduct)

	fmt.Println(products)
}
