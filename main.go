package main

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/filemanager"
	//"github.com/PyaeSoneHan246810/go-practice-project/cmdmanager"
	"github.com/PyaeSoneHan246810/go-practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)
		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result %.0f.json", taxRate*100))
		//cmdManager := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		go priceJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println("Couldn't process job.")
				fmt.Println(err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}
}
