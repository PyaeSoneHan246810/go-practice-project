package prices

import (
	"fmt"

	"github.com/PyaeSoneHan246810/go-practice-project/conversion"
	"github.com/PyaeSoneHan246810/go-practice-project/filemanager"
)

type TaxIncludedPriceJob struct {
	IOManager         *filemanager.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func NewTaxIncludedPriceJob(ioManager *filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: ioManager,
		TaxRate:   taxRate,
	}
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := job.IOManager.OpenFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	lines, err := job.IOManager.ReadLines(file)
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

	file, err := job.IOManager.CreateFile()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = job.IOManager.WriteJson(file, job)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
}
