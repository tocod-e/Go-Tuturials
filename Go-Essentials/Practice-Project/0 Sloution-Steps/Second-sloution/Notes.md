# Go Study Notes

---

## Section: Main Logic and Package Usage

This file serves as the **entry point** for the application. It orchestrates the entire process by using the reusable logic defined in the separate `prices` package.

### 1. Importing Custom Packages

```go
import (
    "example.com/practice/prices"
)
```

* **Custom Import**: We are importing our own local package `prices`.
* **Path**: The import path `example.com/practice/prices` corresponds to the module name defined in `go.mod` plus the subdirectory `prices`.
* **Access**: Once imported, we can access exported (Capitalized) functions and types from that package via `prices.Name`.

### 2. The Loop (Orchestrator)

```go
func main() {
    taxRates := []float64{0, 0.07, 0.1, 0.15}

    for _, taxRateValue := range taxRates {
        // ...
    }
}
```

* **Data Source**: We define the list of `taxRates` we want to process.
* **Iteration**: The loop ensures we run the job logic for *each* tax rate individually.

### 3. Using the Constructor and Methods

```go
priceJob := prices.NewTaxIncludedPriceJob(taxRateValue)
priceJob.Process()
```

* **`prices.NewTaxIncludedPriceJob(...)`**:
  * We call the **Constructor** function from the `prices` package.
  * We pass the current `taxRateValue` (e.g., 0.07).
  * This returns a new `Job` instance (struct pointer) configured with that specific tax rate.
* **`priceJob.Process()`**:
  * We call the **Method** attached to that struct.
  * This triggers the actual calculation and printing logic defined inside `prices.go`.
* **Separation of Concerns**:
  * `main.go` only knows *what* to do (create a job, run it).
  * `prices.go` knows *how* to do it (calculate the math).

### 4. Code Structure

* **`package main`**: Required for the executable entry point.
* **`func main()`**: Where the program execution begins.
