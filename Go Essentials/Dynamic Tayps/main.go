package main

import "fmt"

func main() {
	// Integers
	result1 := add(5, 10) // type of result is interface
	fmt.Println("Result of adding integers:", result1)

	result2 := genericAdd(5, 10) // type of result is int
	fmt.Println("Result of adding integers with generics:", result2)

	// Floats
	result3 := add(5.5, 10.5)
	fmt.Println("Result of adding floats:", result3)

	result4 := genericAdd(5.5, 10.5)
	fmt.Println("Result of adding floats with generics:", result4)

	// Strings
	result5 := add("Hello, ", "World!")
	fmt.Println("Result of adding strings:", result5)

	result6 := genericAdd("Hello, ", "World!")
	fmt.Println("Result of adding strings with generics:", result6)

	// Mixed Types (Runtime failure vs Compile-time error)
	result7 := add(5, "World") // Returns nil because types don't match
	fmt.Println("Result of adding mixed types (int + string):", result7)

	// result8 := genericAdd(5, "World") // Uncommenting this causes a compile error

}

func add(a, b interface{}) interface{} {
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}
	aStr, aIsStr := a.(string)
	bStr, bIsStr := b.(string)
	if aIsStr && bIsStr {
		return aStr + bStr
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFloat := b.(float64)
	if aIsFloat && bIsFloat {
		return aFloat + bFloat
	}
	return nil

}

func genericAdd[T int | float64 | string](a, b T) T {
	return a + b
}
