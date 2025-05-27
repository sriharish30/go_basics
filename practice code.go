package main

import (
	"example/practice/prices"
)

func main() {

	Taxrate := []float64{0, 0.07, 0.1, 0.15}

	for _, rate := range Taxrate {
		JobPrice := prices.NewTaxIncludedPriceJob(rate)
		JobPrice.Process()
	}

}
