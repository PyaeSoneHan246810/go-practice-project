package main

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
		fmt.Println(priceJob.TaxIncludedPrices)
	}
}
