package lists

import "fmt"

func main() {
	// Other Practice
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	fmt.Println(prices[0:1])
	prices[1] = 99.99
	fmt.Println(prices)

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)

	fmt.Println("=============================")
	discountPrices := []float64{101.99, 80.99, 20.99}
	fmt.Println(discountPrices)
	fmt.Println(prices)
	prices = append(prices, discountPrices...)
	fmt.Println(prices)
}