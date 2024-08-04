package main

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/filemanager"
	//"github.com/PyaeSoneHan246810/go-practice-project/cmdmanager"
	"github.com/PyaeSoneHan246810/go-practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	for _, taxRate := range taxRates {
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result %.0f.json", taxRate*100))
		//cmdManager := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Couldn't process job.")
			fmt.Println(err)
		}
	}
}
