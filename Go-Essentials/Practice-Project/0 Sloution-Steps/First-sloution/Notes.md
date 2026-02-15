# Go Study Notes

---

## Section: Practice Project Logic (Maps & Slices)

This code calculates tax-inclusive prices for a list of products across multiple tax rates. It uses nested loops and complex data structures (a map of slices).

### 1. Data Initialization

```go
prices := []float64{10, 20, 30}
taxRates := []float64{0, 0.07, 0.1, 0.15}
result := make(map[float64][]float64)
```

* **`result`**: A map where:
  * **Key (`float64`)**: The tax rate (e.g., `0.07`).
  * **Value (`[]float64`)**: A slice of prices calculated with that tax rate.

### 2. The Outer Loop (Tax Rates)

```go
for _, taxRate := range taxRates {
    taxIncludedPrices := make([]float64, len(prices))
    // ...
    result[taxRate] = taxIncludedPrices
}
```

* We iterate through each `taxRate`.
* **Important**: Inside this loop, we create a **new, empty slice** `taxIncludedPrices` with a length equal to the number of original prices (`len(prices)`).
  * This slice will hold the calculated prices *specifically for this tax rate*.

### 3. The Inner Loop (Prices)

```go
for priceIndex, price := range prices {
    taxIncludedPrices[priceIndex] = price * (1 + taxRate)
}
```

* We iterate through the original `prices` slice.
* **Calculation**: `price * (1 + taxRate)` adds the tax to the base price.
* **Storage**: We use `priceIndex` to store the result in the corresponding position of our new `taxIncludedPrices` slice.
* **Note on Shadowing**: In your code, you used `for priceIndex, prices := range prices`. The inner variable `prices` (the values 10, 20, 30) **shadows** the outer slice variable `prices`. While valid, it's often clearer to name the inner variable singular (e.g., `price`).

### 4. Code Structure & Imports

* **`package main`**: Defines this as an executable program.
* **`import "fmt"`**:
  * Imports the **Format** package.
  * Used here for `fmt.Println` to print the final map.
