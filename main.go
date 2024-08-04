package main

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/filemanager"
	"github.com/PyaeSoneHan246810/go-practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		ioManager := filemanager.New("prices.txt", fmt.Sprintf("result %.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(ioManager, taxRate)
		priceJob.Process()
	}
}
