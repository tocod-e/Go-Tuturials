# Go Study Notes: Detailed Analysis of `prices.go`

---

## 1. Package and Imports

```go
package prices

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)
```

* **`package prices`**: Declares that this file belongs to the `prices` package.
* **Imports**:
  * `bufio`: Provides buffered I/O, specifically the `Scanner`, for efficient reading of data.
  * `fmt`: The "Format" package, used for printing (`Println`) and formatting strings (`Sprintf`).
  * `os`: Provides an interface to correct Operating System functionality, like opening files (`os.Open`).
  * `strconv`: "String Converter", used to convert strings (like text from a file) into numbers (floats).

---

## 2. Struct Definition

```go
type TaxIncludedPriceJob struct {
    TaxRate           float64
    InputPrices       []float64
    TaxIncludedPrices map[string]float64
}
```

* **`type ... struct`**: Defines a new custom data structure.
* **Fields**:
  * `TaxRate` (`float64`): Stores the tax percentage (e.g., 0.07).
  * `InputPrices` (`[]float64`): A slice (dynamic list) to hold the prices read from the file.
  * `TaxIncludedPrices`: A map to store the results.
    * **Key** (`string`): The formatted original price.
    * **Value** (`float64`): The final price with tax.

---

## 3. Constructor Function

```go
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
    return &TaxIncludedPriceJob{
        TaxRate:     taxRate,
        InputPrices: []float64{10, 20, 30}, // Initial dummy data (overwritten later)
    }
}
```

* **Purpose**: Creates and initializes a new `TaxIncludedPriceJob`.
* **Parameter**: Accepts `taxRate` so each job can have a different rate.
* **Return**: Returns a pointer (`*`) to the struct for efficiency.
* **`&TaxIncludedPriceJob{...}`**: Creates the struct and gets its memory address.

---

## 4. Method: `Process` (The Main Logic)

```go
func (job *TaxIncludedPriceJob) Process() {
    job.LoadData()
    // ...
```

* **Receiver**: `(job *TaxIncludedPriceJob)` attaches this function to the struct pointer.
* **`job.LoadData()`**: Calls the helper method to read the file and fill `job.InputPrices`.

```go
    result := make(map[string]string)
    for _, priceVal := range job.InputPrices {
        taxIncludedPrice := priceVal * (1 + job.TaxRate)
        result[fmt.Sprintf("%.2f", priceVal)] = fmt.Sprintf("%.2f", taxIncludedPrice)
    }
    fmt.Println(result)
}
```

* **`make(map[string]string)`**: Creates an empty map to store strings (formatted prices).
* **Loop**: Iterates through every price in `InputPrices`.
* **Calculation**: `priceVal * (1 + job.TaxRate)` adds the tax.
* **String Formatting**:
  * `fmt.Sprintf("%.2f", ...)`: Formats the number to strictly **2 decimal places** (e.g., "10.00").
  * We use this formatted string as both the **Key** and the **Value** in the map for clean output.
* **`fmt.Println(result)`**: Prints the final map to the console.

---

## 5. Method: `LoadData` (File Reading)

This method is responsible for reading `prices.txt` and populating `job.InputPrices`.

### A. Opening the File

```go
func (job *TaxIncludedPriceJob) LoadData() {
    file, err := os.Open("prices.txt")
    if err != nil {
        fmt.Println("Could not open file!")
        fmt.Println(err)
        return
    }
```

* **`os.Open`**: Tries to open the file.
* **Error Check**: If `err` is not `nil` (meaning something went wrong), it prints the error and **stops** execution (`return`).

### B. Scanning Lines

```go
    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
```

* **`bufio.NewScanner`**: Creates a tool to read the file.
* **`scanner.Scan()` loop**: Runs as long as there is a new line to read.
* **`scanner.Text()`**: Gets the text of the current line.
* **`append`**: Adds that line to our `lines` slice.

### C. Check for Scanning Errors

```go
    err = scanner.Err()
    if err != nil {
        // ... error handling and file.Close()
        return
    }
```

* Checks if the scanner stopped due to an error (rather than just reaching the end of the file).

### D. Converting Strings to Floats

```go
    prices := make([]float64, len(lines))

    for lineIndex, lineVal := range lines {
        floatPrice, err := strconv.ParseFloat(lineVal, 64)
        if err != nil {
            // ... error handling
            return
        }
        prices[lineIndex] = floatPrice
    }
    job.InputPrices = prices
}
```

* **`make([]float64, len(lines))`**: Creates a slice of floats with the *exact* size needed (same as the number of lines).
* **`strconv.ParseFloat(lineVal, 64)`**: Converts the text "10.99" into the number `10.99`.
  * If the file contains text like "abc", this will return an error.
* **Assignment**: `prices[lineIndex] = floatPrice` puts the number into the correct slot.
* **Final Step**: `job.InputPrices = prices` updates the struct with the data we just read and converted.
