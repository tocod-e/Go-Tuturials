package recursion

import (
	"fmt"
)

/* Making Sense of Recursion  */

func main() {

	fact := factorial(5)
	fmt.Println(fact)

	factorial := factor(5)
	fmt.Println(factorial)

}

// With Simple func
func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

// With Recursion
func factor(number int) int {
	if number == 0 {
		return 1
	}
	return number * factor(number-1)
}
