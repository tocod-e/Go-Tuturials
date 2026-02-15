package main

import (
	"fmt"
)
/* Using Variadic functions  */

func main() {
	numbers := []int {1,2,3,4,5}
	sum := sumup(numbers)
	fmt.Println(sum)

	sumInt := sume(1,2,3,4,5,6,-3,-3)
	fmt.Println(sumInt)

	newSum := sume(numbers...)
	fmt.Println(newSum)

}


// The Standard e Approach (Standard Loop)
func sumup(numbers []int) int{
	sum := 0
	for _, val := range numbers{
		sum += val // sum = sum + val
	}
	return sum
}

// the Variadic Approach

func sume(numbers ... int) int{
	sum := 0
	for _, val := range numbers{
		sum += val
	}
	return sum
}