package main

import "fmt"

func main() {
	result1 := add(5, 10)  // type of result is interface
	fmt.Println("Result of adding integers:", result1)

	result2 := genericAdd(5, 10)  // type of result is int
	fmt.Println("Result of adding integers with generics:", result2)
	
}


// The function `add` takes two parameters of any type and returns their sum if they are both integers,
// strings, or floats and they are of the same type. If the types do not match or are not supported, it returns nil. The function uses type assertions to determine the types of the parameters at runtime.
// it return a type of interface, which can hold any type of value. The caller needs to perform a type assertion to retrieve the actual value from the interface. This approach allows for flexibility but requires careful handling to avoid runtime errors due to type mismatches.
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

// The function `genericAdd` can add two values of type `int`, `float64`, or `string` and return the
// result. the reurun value can be of type `int`, `float64`, or `string` depending on the types of the input parameters. The function uses Go's generics feature to allow for type flexibility while ensuring type safety at compile time. The caller does not need to perform any type assertions, as the function will return the correct type based on the input parameters.
func genericAdd[T int|float64|string](a, b T) T {
	return a + b
}
