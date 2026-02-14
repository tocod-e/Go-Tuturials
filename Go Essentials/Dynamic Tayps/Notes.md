# Go Study Notes

## Section: Dynamic Types vs Generics - 2026-02-14

### Overview
This code demonstrates two approaches to writing polymorphic functions in Go:
1. **Dynamic Typing** using the empty interface `interface{}`.
2. **Static Typing** using Go 1.18+ **Generics**.

### Line-by-Line Explanation

#### Package and Imports
```go
package main

import "fmt"
```
- `package main`: Defines this as an executable program rather than a library. The `main` function is the entry point.
- `import "fmt"`: Imports the **Format** package, which handles input/output formatting. We use it here specifically for printing to the console (`Println`).

#### The `main` Function
```go
func main() { ... }
```
The `main` function orchestrates the demonstration. It calls both `add` (dynamic) and `genericAdd` (static) with different data types to show behavior.

**1. Integers**
```go
	// Integers
	result1 := add(5, 10) // type of result is interface
	fmt.Println("Result of adding integers:", result1)

	result2 := genericAdd(5, 10) // type of result is int
	fmt.Println("Result of adding integers with generics:", result2)
```
- `add(5, 10)`: Calls the interface-based function. The integers `5` and `10` are implicitly converted to `interface{}`. The return type is also `interface{}`, meaning the compiler doesn't know it's an integer at this point.
- `genericAdd(5, 10)`: Calls the generic function. The compiler infers `T` is `int`. The return type is strictly `int`.

**2. Floats**
```go
	// Floats
	result3 := add(5.5, 10.5)
	...
	result4 := genericAdd(5.5, 10.5)
```
- Similar to integers, but passes `float64` values. Use of `float64` is standard in Go for floating-point numbers.

**3. Mixed Types (The Danger Zone)**
```go
	// Mixed Types (Runtime failure vs Compile-time error)
	result7 := add(5, "World") // Returns nil because types don't match
	fmt.Println("Result of adding mixed types (int + string):", result7)

	// result8 := genericAdd(5, "World") // Uncommenting this causes a compile error
```
- `add(5, "World")`: The compiler allows this because both `int` and `string` satisfy `interface{}`. Use of `interface{}` sacrifices type safety for flexibility. The function returns `nil` at **runtime** because it cannot add an integer and a string.
- `genericAdd(5, "World")`: The compiler **forbids** this. It tries to match `T` to both `int` and `string` simultaneously, which is impossible. This prevents bugs before the code even runs.

#### The `add` Function (Dynamic / Interface{})
```go
func add(a, b interface{}) interface{} {
    ...
}
```
- **Signature**: Accepts any type (`interface{}`) and returns any type (`interface{}`).
- **Internal Logic**: It must manually check what types were passed.

**Type Assertions**
```go
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)

	if aIsInt && bIsInt {
		return aInt + bInt
	}
```
- `a.(int)`: This is a **Type Assertion**. It checks "Is the value inside interface `a` an `int`?".
- Returns two values:
  1. `aInt`: The underlying integer value (if successful, else 0).
  2. `aIsInt`: A boolean (`true` if successful, `false` otherwise).
- Checks are repeated for `string` and `float64`.
- **Inefficiency**: This requires runtime checking and allocation/conversion overhead.

#### The `genericAdd` Function (Static / Generics)
```go
func genericAdd[T int|float64|string](a, b T) T {
	return a + b
}
```
- **Constraint `[T int|float64|string]`**: Defines a type parameter `T`. `T` is restricted to be *only* `int`, `float64`, or `string`. This is a **Union Element**.
- **Signature `(a, b T) T`**:
  - Arguments `a` and `b` must both be type `T` (e.g., both `int` or both `string`).
  - Returns a value of exact type `T`.
- **Performance**: The compiler generates optimized code for the specific type used. No runtime type assertions are needed.

### Summary Table

| Feature | `interface{}` (`add`) | Generics (`genericAdd`) |
| :--- | :--- | :--- |
| **Type Safety** | Low (Runtime Checks) | High (Compile-Time Checks) |
| **Performance** | Slower (Overhead) | Faster (Optimized Check) |
| **Flexibility** | Extreme (Any code) | Constrained (Specific Set) |
| **Readability** | Verbose (Assertions) | Clean (Direct logic) |

