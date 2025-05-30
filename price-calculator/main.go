package main

import (
	"fmt"
	"price-calculator/cmdmanager"
	"price-calculator/prices"
)

func main() {
	//var prices []float64 = []float64{10, 20, 30} //two ways to initialize
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	//result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		//fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}
	}
}
