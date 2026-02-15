package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()
	result := make(map[string]string)
	for _, priceVal := range job.InputPrices {
		taxIncludedPrice := priceVal * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}
	fmt.Println(result)
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Reading the file content failed.")
		fmt.Println(err)
		file.Close()
		return
	}

	prices := make([]float64, len(lines))

	for lineIndex, lineVal := range lines{
		floatPrice, err := strconv.ParseFloat(lineVal, 64)
		if err != nil {
			fmt.Println("Converting price to float faild.")
			fmt.Println(err)
			return
		}
		prices[lineIndex] = floatPrice
	}
	job.InputPrices = prices
}
