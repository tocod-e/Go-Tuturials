package main

import (
	"fmt"

	//"example.com/practice/cmdmanger"
	"example.com/practice/filemanager"
	"example.com/practice/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRateValue := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRateValue*100))
		//cmdm := cmdmanger.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRateValue)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not Process the Job.")
			fmt.Println(err)
		}
	}

}
