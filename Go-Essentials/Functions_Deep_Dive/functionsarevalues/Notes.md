# Go Study Notes

---

## Section: Functions as Values and Function Types

In Go, functions are **first-class citizens**. This means you can assign them to variables, pass them as arguments to other functions, and return them from functions.

### 1. Function Types as Parameters

```go
func transformNumbers(numbers *[]int, transform func(int) int) []int {
    // ...
}
```

* **`transform func(int) int`**:
  * This is a parameter named `transform`.
  * Its type is a **function** that takes an `int` and returns an `int`.
  * This allows `transformNumbers` to be generic: it doesn't care *how* the number is transformed (doubled, tripled, negated), as long as the function signature matches.

### 2. Passing Functions as Arguments

```go
doubled := transformNumbers(&numbers, double)
tripled := transformNumbers(&numbers, triple)
```

* **`double`** and **`triple`**: These are defined function names. notice we pass them **without parentheses** `()`.
  * `double(...)`: Calls the function.
  * `double`: Refers to the function definition itself (the value).

### 3. Pointers to Slices (A Note)

```go
func transformNumbers(numbers *[]int, ...)
// ...
for _, val := range *numbers { ... }
```

* **`*[]int`**: The function accepts a **pointer** to a slice of integers.
* **`*numbers`**: Inside the function, we must **dereference** the pointer to get the actual slice value to iterate over it.
* **Go Note**: While valid, passing a pointer to a slice is often unnecessary in Go because slices are already lightweight "descriptors" (containing a pointer to the underlying array) passed by value. However, using `*[]int` allows you to modify the *slice header* itself (e.g., resizing/reallocating the slice in the caller), though here it's just used for reading.

### 4. Implementation Details

```go
dNumbers := []int{}
// ...
dNumbers = append(dNumbers, transform(val))
```

* **`transform(val)`**: Here, the function passed in (e.g., `double`) is finally executed on the value `val`.
* **`dNumbers`**: A new slice is created to store the results, leaving the original `numbers` slice unchanged.

### 5. Code Structure & Imports

* **`package main`**: Defines this as an executable program.
* **`import "fmt"`**:
  * Imports the **Format** package.
  * Used here for `fmt.Println` to print the slices to the console (standard output).

---

## Section: Function Type Aliases

Go allows you to define a specific **type** for a function signature. This improves readability and reusability, especially when the function signature is complex or used in multiple places.

### 1. Defining a Function Type

```go
type transformFn func(int) int
```

* **`type`**: Keyword to start a type definition.
* **`transformFn`**: The name of our new custom type.
* **`func(int) int`**: The underlying type. Any function that takes one `int` and returns one `int` matches this type.

### 2. Using the Function Type

Refactoring the `transformNumbers` function signature:

**Before:**

```go
func transformNumbers(numbers *[]int, transform func(int) int) []int
```

**After:**

```go
func transformNumbers(numbers *[]int, transform transformFn) []int
```

* **Cleaner Code**: instead of repeating `func(int) int` everywhere, we just use `transformFn`.
* **Type Safety**: The compiler ensures that only functions matching the `transformFn` signature (like `double` and `triple`) can be passed.

---

## Section: Returning Functions as Values

Just as functions can be passed *into* other functions, they can also be **returned** from functions. This is a core feature of functional programming support in Go.

### 1. Function Returning a Function

```go
func getTransformerFunction(numbers *[]int) transformFn {
    if (*numbers)[0] == 1 {
        return double
    } else {
        return triple
    }
}
```

* **Return Type `transformFn`**:
  * The function signature ends with `transformFn`.
  * This means `getTransformerFunction` **must return a function** that matches the `transformFn` signature (taking an int, returning an int).
* **Logic**:
  * It checks the first element of the slice (`(*numbers)[0]`).
  * If it's `1`, it returns the **function** `double`.
  * Otherwise, it returns the **function** `triple`.
  * Note: We return `double` (the function itself), not `double(5)` (the result of calling it).

### 2. Using Returned Functions

```go
transformFunc1 := getTransformerFunction(&numbers)
transformFunc2 := getTransformerFunction(&moreNumbers)
```

* **`transformFunc1`**:
  * Calls `getTransformerFunction` with `[1, 2, 3, 4]`.
  * Since the first number is `1`, it returns `double`.
  * `transformFunc1` now holds the `double` function.
* **`transformFunc2`**:
  * Calls it with `[5, 1, 4]`.
  * Since the first number is `5`, it returns `triple`.
  * `transformFunc2` now holds the `triple` function.

### 3. Executing the Variable Functions

```go
transformedNumers := transformNumbers(&numbers, transformFunc1)
moreTransformedNumers := transformNumbers(&moreNumbers, transformFunc2)
```

* We treat `transformFunc1` just like any other variable.
* We pass it to `transformNumbers`, which executes it on every element.
* **Result**:
  * `numbers` are doubled (`[2 4 6 8]`).
  * `moreNumbers` are tripled (`[15 3 12]`).
