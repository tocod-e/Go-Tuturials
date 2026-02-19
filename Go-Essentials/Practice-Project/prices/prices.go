package prices

import (
	"fmt"
	"example.com/practice/conversion"
	"example.com/practice/iomanager"
)

type TaxIncludedPriceJob struct {
	IOManger          iomanager.IOManger `json:"-"`
	TaxRate           float64            `json:"tax_rate"`
	InputPrices       []float64          `json:"input_prices"`
	TaxIncludedPrices map[string]string  `json:"tax_included_prices"`
}

func NewTaxIncludedPriceJob(iom iomanager.IOManger, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManger:    iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() error{
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)
	for _, priceVal := range job.InputPrices {
		taxIncludedPrice := priceVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	job.TaxIncludedPrices = result

	return 	job.IOManger.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error{
	lines, err := job.IOManger.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		return err
	}
	job.InputPrices = prices
	return nil

}
