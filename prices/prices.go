package prices

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/conversion"
	"github.com/PyaeSoneHan246810/go-practice-project/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := filemanager.OpenFile("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines, err := filemanager.ReadLines(file)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	file, err := filemanager.CreateFile(fmt.Sprintf("result %.0f.json", job.TaxRate*100))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = filemanager.WriteJson(file, job)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
}
