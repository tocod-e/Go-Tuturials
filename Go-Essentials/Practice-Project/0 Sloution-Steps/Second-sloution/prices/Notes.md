# Go Study Notes

---

## Section: Structs, Constructors, and Methods

This file demonstrates how to organize logic around data using Structs. It introduces a `Job` struct that holds data and a method to process that data.

### 1. Defining the Custom Type (Struct)

```go
type TaxIncludedPriceJob struct {
    TaxRate           float64
    InputPrices       []float64
    TaxIncludedPrices map[string]float64
}
```

* **Purpose**: This struct groups all the data needed for one specific calculation job.
* **Fields**:
  * `TaxRate`: The specific tax percentage for this job (e.g., 0.07).
  * `InputPrices`: The list of original prices to process.
  * `TaxIncludedPrices`: A map to store the calculated results.

### 2. The Constructor Function

```go
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
    return &TaxIncludedPriceJob{
        InputPrices: []float64{10, 20, 30},
        TaxRate:     taxRate,
    }
}
```

* **Naming Convention**: Functions starting with `New...` are typically **Constructors**. They prepare and return a new instance of a struct.
* **Arguments**: Takes `taxRate` as input because it changes per job.
* **Hardcoded Values**: Sets `InputPrices` to specific values (`10, 20, 30`) for this example.
* **Return Value**: Returns `*TaxIncludedPriceJob` (a **pointer**). This is efficient because it passes the *address* of the data, not a copy of the entire struct.

### 3. Adding Methods to the Struct

```go
func (job TaxIncludedPriceJob) Process() {
    result := make(map[string]float64)
    for _, priceVal := range job.InputPrices {
        result[fmt.Sprintf("%.2f", priceVal)] = priceVal * (1 + job.TaxRate)
    }
    fmt.Println(result)
}
```

* **Receiver `(job TaxIncludedPriceJob)`**: This attaches the `Process` function to the struct. It means you can call it like `myJob.Process()`.
* **String Formatting**:
  * `fmt.Sprintf("%.2f", priceVal)`: Converts the float price (e.g., `10`) into a formatted string with 2 decimal places (e.g., `"10.00"`).
  * This string is used as the **Key** for the map.
* **Logic**: Iterates through the input prices, calculates the tax, and stores the result in the map.

### 4. Code Structure & Imports

* **`package prices`**: Defines this code as part of the reusable `prices` package.
* **`import "fmt"`**: used for string formatting (`Sprintf`) and printing (`Println`).
