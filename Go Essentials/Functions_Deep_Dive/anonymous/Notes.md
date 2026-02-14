# Go Study Notes

---

## Section: Anonymous Functions

Go allows you to define functions **without a name**. These are called **anonymous functions** (or *function literals*). They are useful when you want to define a small piece of logic right where it's used.

### 1. Structure of an Anonymous Function

```go
func(number int) int {
    return number * 2
}
```

* **`func`**: Keyword to start the function definition.
* **`(number int)`**: The parameters (just like a named function).
* **`int`**: The return type.
* **`{ ... }`**: The function body.
* **No Name**: Notice there is no name between `func` and `(number int)`.

### 2. Passing as an Argument

```go
transformed := transformNumbers(&numbers, func(number int) int {
    return number * 2
})
```

* Instead of defining `func double(...)` separately and passing `double` by name, we define the **entire function inline**.
* **Why?**
  * **Conciseness**: If `double` is only used here, defining it inline saves space and keeps related logic together.
  * **Flexibility**: You can write any logic you need on the fly without cluttering the global scope with helper functions.

### 3. Comparison

**Named Function Approach:**

```go
func double(n int) int { return n * 2 }
// ...
transformNumbers(&nums, double)
```

**Anonymous Function Approach:**

```go
transformNumbers(&nums, func(n int) int { return n * 2 })
```

Both achieve the same result. Use anonymous functions for quick, one-off operations.

### 4. Code Structure & Imports

* **`package main`**: Defines this as an executable program.
* **`import "fmt"`**:
  * Imports the **Format** package.
  * Used here for `fmt.Println` to print the output.

---

## Section: Closures and Factory Functions

Go supports **Closures**: anonymous functions that reference variables from outside their own body. A powerful use case for this is creating **Factory Functions**—functions that build and return other functions.

### 1. The Factory Function Protocol

```go
func creatTransformer(factor int) func(int) int {
    return func(number int) int {
        return number * factor
    }
}
```

* **Goal**: This function *creates* a new multiplier function.
* **`factor int`**: This variable is passed in *once* when creating the function.
* **The Closure**:
  * The returned anonymous function `func(number int) int` **closes over** the `factor` variable.
  * Even after `creatTransformer` finishes executing, the returned function **remembers** the value of `factor`.

### 2. Creating Custom Functions

```go
double := creatTransformer(2)
triple := creatTransformer(3)
```

* **`double`**:
  * We call `creatTransformer(2)`.
  * Returns a function where `factor` is locked in as `2`.
  * Effect: `func(n) { return n * 2 }`.
* **`triple`**:
  * We call `creatTransformer(3)`.
  * Returns a function where `factor` is locked in as `3`.
  * Effect: `func(n) { return n * 3 }`.

### 3. Usage

```go
doubled := transformNumbers(&numbers, double)
tripled := transformNumbers(&numbers, triple)
```

* We can now use these generated functions just like any standard function.
* **Benefit**: We avoided writing separate `func double()` and `func triple()` functions manually. We have a **dynamic** way to create any multiplier we need (e.g., `timesTen := creatTransformer(10)`).
