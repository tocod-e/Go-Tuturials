package main

import (
	"fmt"
)

// Dynamic Array(Dynamic Slices)
func main(){
	prices := []float64{10.99, 8.99}
	fmt.Println(prices)
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	fmt.Println(prices)
	/* prices[2] = 11.99
	fmt.Println(prices) */
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices, prices) // THe Orignal Slice did not changed instead new Slice will be created with contain the orignal and the add Value and the orignal stay unchanged too

	// we can chnage the Orignal by asigning the append to thhe orignal Prices so we avoid to create new slices
	fmt.Println("=============")
	fmt.Println(prices)
	prices = append(prices, 5.99)
	fmt.Println(prices)

	// Remove elemnt from the orignal
	fmt.Println("=========")
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)
}
/* 
// Strict Array and can not chnage or add new values to the Array
func main() {
	var productNames [4]string
	fmt.Println(productNames)
	var products [4]string = [4]string{"A Book"}
	fmt.Println(products)

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	// indexing Array
	fmt.Println(prices[2])
	fmt.Println(prices[0])  // first Index

	// asigen value to an Index in the Array
	products[2] = "A Carpet"
	fmt.Println(products)

	// Selecting parts of Array using slices
	featuredPrices := prices[1:3] // start from index 1 untill index 2 and execluse the last Index 3
	fmt.Println(featuredPrices)

	// More ways of selecting slices
	featuredPrice := prices[:3] // from the begining untill the second the third will be execluded
	fmt.Println(featuredPrice)

	featuredPr := prices[1:] // start from index 1 untill the end
	fmt.Println(featuredPr)
	// In go we can not use negative slices and also can not use a higher index which the orignal array does not have (Out of Pound)

	// Slices on other Slices Or Nested Slices
	highlightedPrices := featuredPrices[:1]
	fmt.Println(highlightedPrices)

	// Editing a value on a Slices edit the Value in the orignal array
	fmt.Println("===========")
	fmt.Println(featuredPrices[0])
	fmt.Println(prices)
	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices[0])
	fmt.Println(prices)

	// MEtadata of Slices
	// len
	fmt.Println(len(featuredPrices))
	fmt.Println(len(prices))

	// length and Capacity
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(prices), cap(prices))

	// Capacity is defrent 

	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices))
	fmt.Println(cap(highlightedPrices)) // it gives 3 becase its based on feuturedprices which in turn in base of array prices so they share the same array under the hood
	// the Capacity only count towards the end of the orignal array not towards the begining of the array

	// REclise highlightedPrices in the same highlightedPrices
	highlightedPrices = highlightedPrices[:3] // in this case the len eqeul the capacity
	fmt.Println(highlightedPrices)
	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

} */