# Go Study Notes

---

## Section: Recursion

Recursion occurs when a function **calls itself**. It's a different way to solve problems that involves breaking a task down into smaller, similar sub-tasks.

### 1. The Iterative Approach (Standard Loop)

```go
func factorial(number int) int {
    result := 1
    for i := 1; i <= number; i++ {
        result = result * i
    }
    return result
}
```

* **Logic**: This uses a `for` loop to multiply numbers from 1 up to `number`.
* **State**: The `result` variable keeps track of the running total.
* **Pros/Cons**: Easy to understand for simple counting, but can be verbose for complex tree-like structures.

### 2. The Recursive Approach

```go
func factor(number int) int {
    if number == 0 {
        return 1
    }
    return number * factor(number - 1)
}
```

* **Base Case (`if number == 0`)**:
  * **Critical**: Every recursive function *must* have a condition to **stop** calling itself.
  * If `number` is 0, we return 1 (since factorial of 0 is 1). This stops the recursion.
* **Recursive Step (`return number * factor(number - 1)`)**:
  * The function calls *itself* with a smaller number (`number - 1`).
  * **Example: `factor(5)`**:
    1. Returns `5 * factor(4)`
    2. ...which returns `4 * factor(3)`
    3. ...which returns `3 * factor(2)`
    4. ...which returns `2 * factor(1)`
    5. ...which returns `1 * factor(0)`
    6. `factor(0)` returns `1`.
    7. The chain resolves upwards: `1 * 1 * 2 * 3 * 4 * 5 = 120`.

### 3. Comparison

* **Recursion** is often cleaner and more mathematical for problems like this.
* **Iteration** is usually more memory-efficient (no stack buildup).

### 4. Code Structure & Imports

* **`package main`**: Defines this as an executable program.
* **`import "fmt"`**:
  * Imports the **Format** package.
  * Used here for `fmt.Println` to print the calculated factorials.
