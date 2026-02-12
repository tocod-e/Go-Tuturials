package main

import (
	"fmt"
	"math"
)

func main(){
	/*
	const inflationRate float64 = 2.5
	var investmentAmount float64
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Inter the Years: ")
	fmt.Scan(&years)
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Printf("Future Value of Investment: %.1f $ \nThe Future Real Value %.2f $ \n", futureValue, futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
	fmt.Print(formattedFV , formattedFRV)

	//fmt.Printf(`Future Value of Investment: %.1f $ /\n
	//The Future Real Value: %.2f $`,
	// futureValue, futureRealValue)
	*
	*/
	
	runInvestmentCalculator()
}




func calculateFutureValue(investmentAmount float64, years float64, expectedReturnRate float64) float64 {
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	return futureValue
}

func calculateFutureRealValue(futureValue float64, years float64, inflationRate float64) float64 {
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureRealValue
}

func formattedOutput(futureValue float64, futureRealValue float64) (string, string) {
	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Real Value: %.2f\n", futureRealValue)
	return formattedFV, formattedFRV
}

func getUserInput() (float64, float64) {
	var investmentAmount float64
	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	var years float64
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)
	return investmentAmount, years
}
func runInvestmentCalculator() {
	const inflationRate float64 = 2.5
	var expectedReturnRate float64 = 5.5
	investmentAmount, years := getUserInput()
	futureValue := calculateFutureValue(investmentAmount, years, expectedReturnRate)
	futureRealValue := calculateFutureRealValue(futureValue, years, inflationRate)
	formattedFV, formattedFRV := formattedOutput(futureValue, futureRealValue)
	fmt.Print(formattedFV , formattedFRV)
}

