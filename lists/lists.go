package lists

import "fmt"

func main() {
	prices := []float64{10.999, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	//prices = append(prices, 5.99)
	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	prices = prices[1:]
	fmt.Println(prices)

	discountPrices := []float64{101.99, 0.99, 20.59}
	prices = append(prices, discountPrices...)

	fmt.Println(prices)
}

// func main() {
// 	var productNames = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)
// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)
// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:3] // 1st value is included, 2nd value is excluded
// 	fmt.Println(featuredPrices)

// 	featuredPrices1 := prices[:3]
// 	fmt.Println(featuredPrices1)

// 	featuredPrices2 := prices[1:]
// 	featuredPrices2[0] = 199.99 // when you create a slice, you dont copy
// 	// the original array, slice has a reference to a part of the array
// 	fmt.Println(featuredPrices2)

// 	// create slices from slices
// 	highlightedPrices := featuredPrices2[:1]
// 	fmt.Println(highlightedPrices)

// 	fmt.Println(prices) //with change in line 20,
// 	// the original array gets modified

// 	fmt.Println(len(featuredPrices2), cap(featuredPrices))

// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// 	//capacity always counts towards the end of the array

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
