# Go Study Notes

---

## Section: Variadic Functions

Variadic functions are functions that can accept a **variable number of arguments**. In Go, this is done using the ellipsis `...` syntax before the type name in the function parameter.

### 1. The Standard Approach (Slice Parameter)

```go
func sumup(numbers []int) int {
    sum := 0
    for _, val := range numbers {
        sum += val // sum = sum + val
    }
    return sum
}
```

* **Signature**: `sumup(numbers []int)` expects a **single argument**: a slice of integers.
* **Usage**: You *must* create a slice first (`[]int{1, 2, 3}`) and pass that slice to the function.
* **Internal Logic**: Iterates over the slice to calculate the sum.

### 2. The Variadic Approach

```go
func sume(numbers ...int) int {
    sum := 0
    for _, val := range numbers {
        sum += val
    }
    return sum
}
```

* **Signature**: `sume(numbers ...int)` uses the `...int` syntax.
* **Meaning**: "Accept zero or more arguments of type `int`".
* **Internal Logic**:
  * Inside the function, `numbers` is treated exactly like a **slice** (`[]int`).
  * You can iterate over it using `range`, check `len(numbers)`, or access by index.
* **Benefit**: The internal logic is identical to `sumup`, but the *caller's* experience is different.

### 3. Usability Comparison

```go
// Standard (Slice)
numbers := []int{1, 2, 3, 4, 5}
sum := sumup(numbers)

// Variadic
sumInt := sume(1, 2, 3, 4, 5, 6, -3, -3)
```

* **Slice Version**: Requires explicitly constructing a slice. Good when you already *have* a slice.
* **Variadic Version**:
  * Can pass individual values directly: `sume(1, 2)`.
  * Can pass nothing: `sume()`.
  * **Pro Tip**: You can pass an existing slice to a variadic function by unpacking it: `sume(numbers...)`.

### 4. Code Structure & Imports

* **`package main`**: Defines this as an executable program.
* **`import "fmt"`**:
  * Imports the **Format** package.
  * Used here for `fmt.Println` to print the results.
