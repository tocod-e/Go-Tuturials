package functionsarevalues

import (
	"fmt"
)

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Orignal Numbers: ",numbers)

	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled Numbers: ", doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled Numbers: ",tripled)


	// Functions as return Value calls

	moreNumbers := []int{5,1,4}
	transformFunc1 := getTransformerFunction(&numbers)
	transformFunc2 := getTransformerFunction(&moreNumbers)
	//fmt.Println("Double Func",double)
	//fmt.Println("Triple Func",triple)

	fmt.Println("Transformer Numbers: ", transformFunc1)
	fmt.Println("Transformer Numbers: ", transformFunc2)

	transformedNumers := transformNumbers(&numbers, transformFunc1)
	moreTransformedNumers := transformNumbers(&moreNumbers, transformFunc2)

	fmt.Println("transformedNumers: ", transformedNumers)
	fmt.Println("moreTransformedNumers", moreTransformedNumers)

}

// Functions as Values and Function Types

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
	//	dNumbers = append(dNumbers, val * 2)
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

func double(number int) int{
	return  number * 2
}

func triple(number int) int{
	return number * 3
}


// Returning Functions as a Values

func getTransformerFunction(numbers *[]int) transformFn{
	if (*numbers)[0] == 1{
		return double
	}else{
		return triple
	}
	
}