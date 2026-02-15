package anonymous

import (
	"fmt"
)

/*
  Introducing Anonymous Functions
*/

func main() {
	numbers := []int{1, 2, 3, 4}
	fmt.Println("Orignal Numbers: ", numbers)
	// Anonymous Function as a parameter
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println("transformed: ", transformed)

	fmt.Println("============")
	fmt.Println("Orignal Numbers: ", numbers)

	double := creatTransformer(2)
	triple := creatTransformer(3)
	//fmt.Println(double)

	doubled := transformNumbers(&numbers, double)
	fmt.Println("Doubled: ", doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println("Tripled: ", tripled)
}

// Functions as Values and Function Types

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

// Understanding Closurs and Factory functions
func creatTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
