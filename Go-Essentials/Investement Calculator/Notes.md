# Go Study Notes

---

## Section: Investment Calculator Modularity - 2026-02-14

The `investment_calculator.go` code demonstrates refactoring a monolithic `main` function into smaller, reusable functions. This approach improves readability and maintainability.

### 1. Function Structuring
The logic is broken down into specific tasks:
- **`getUserInput()`**: Handles all user interaction (scanning inputs). usage of returning multiple values `(float64, float64)` allows fetching both `investmentAmount` and `years` in one call.
- **`calculateFutureValue(...)`**: Encapsulates the compound interest formula `investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)`.
- **`calculateFutureRealValue(...)`**: Adjusts the future value for inflation.
- **`formattedOutput(...)`**: Formats the results into strings without printing them immediately, returning `(string, string)`.

### 2. Variable Scope and Flow
- **`runInvestmentCalculator`**: Acts as the controller, orchestrating the flow:
    1.  Defines constants/variables (`inflationRate`, `expectedReturnRate`).
    2.  Calls `getUserInput` to initialize `investmentAmount` and `years`.
    3.  Passes these values to calculation functions.
    4.  Calls `formattedOutput` to get display strings.
    5.  Prints the final result.

### 3. Key Go Features Used
- **Multiple Return Values**: Seen in `getUserInput` and `formattedOutput`.
- **`math.Pow`**: Used for power calculations (requires `float64`).
- **Formatting**: `fmt.Sprintf` is used to create formatted strings (e.g., `%.1f` for 1 decimal place, `%.2f` for 2) without printing to stdout directly.

---

## Section: Imperative Logic & Basic Syntax - 2026-02-14

The `main` function (lines 8-36) shows the imperative approach to the same problem. This block highlights fundamental syntax and control flow.

### 1. Variables & Constants
- **`const`**: Used for fixed values like `inflationRate` (cannot change).
- **`var`**: Explicit declaration (e.g., `var investmentAmount float64`).
- **Short Declaration (`:=`)**: Infers type automatically (e.g., `expectedReturnRate := 5.5`).

### 2. User Input with Pointers
- **`fmt.Scan(&variable)`**:  Critically, `Scan` requires a **pointer** (`&`) to the variable to modify its value directly in memory. Without `&`, the function would receive a copy and the original variable would remain zero/empty.

### 3. Math & Logic
- **`math.Pow(x, y)`**: Calculates x^y. Requires casting integers to `float64` if mixed types are used (though here constants are float-compatible).
- **Arithmetic**: Standard operators `+`, `/`, `*` work as expected on matching types.

### 4. Output Formatting
- **`fmt.Printf`**: Prints formatted strings to standard output.
- **`fmt.Sprintf`**: Formats a string and **returns** it, rather than printing. Useful for storing or processing the message before display.
- **Formatting Verbs**:
    - `%.1f`: Float with 1 decimal place.
    - `%.2f`: Float with 2 decimal places.
    - `\n`: Newline character.

---

## Section: Function Deep Dive: calculateFutureValue - 2026-02-14

The `calculateFutureValue` function (lines 39-43) highlights the importance of creating pure functions in Go.

### 1. Function Logic
- **Pure Functionality**: This function takes inputs and returns an output without side effects (no printing, no global state changes). This makes it:
    - **Testable**: Easy to write unit tests for.
    - **Reusable**: Can be called from anywhere, not just `main`.
- **Parameter Types**: Takes three `float64` arguments (`investmentAmount`, `years`, `expectedReturnRate`).
- **Return Type**: Explicitly returns a `float64`.

### 2. Implementation details
- **Formula**: Implements the standard compound interest calculation:
  $$FV = P \times (1 + r/100)^t$$
- **Code**:
  ```go
  futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
  return futureValue
  ```
- **Clarity**: Using named parameters makes the `math.Pow` call self-documenting compared to doing the raw calculation inline.

---

## Section: Function Deep Dive: calculateFutureRealValue - 2026-02-14

The `calculateFutureRealValue` function (lines 46-49) adjusts the investment's future value for inflation, giving the "real" purchasing power.

### 1. Concept: Real Value
- **Inflation Adjustment**: Money loses value over time due to inflation. This function calculates what the future amount is worth in *today's* dollars.
- **Formula**:
  $$Real Value = \frac{Future Value}{(1 + Inflation Rate/100)^{years}}$$

### 2. Implementation Logic
- **Modular Design**: It takes `futureValue` as an input (calculated by the previous function), rather than recalculating it. This chaining of functions (`input -> calculateFV -> calculateRealFV`) is a key modular design pattern.
- **Reusing Logic**: It uses the same mathematical structure (`math.Pow`) as the future value calculation but applying division instead of multiplication.

### 3. Code Analysis
```go
func calculateFutureRealValue(futureValue float64, years float64, inflationRate float64) float64 {
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureRealValue
}
```
- **Inputs**: `futureValue` (from previous step), `years`, `inflationRate`.
- **Output**: `futureRealValue` (the adjusted amount).

---

## Section: Function Deep Dive: formattedOutput - 2026-02-14

The `formattedOutput` function (lines 51-54) demonstrates Go's ability to return multiple values, a powerful feature for error handling and data processing.

### 1. Multiple Return Values
- **Signature**: `func ... (...) (string, string)`
- **Mechanism**: Go functions can return any number of results. This function returns two strings: one for the Future Value and one for the Real Value.
- **Usage**: The caller must assign both values (or ignore one using `_`).
  ```go
  formattedFV, formattedFRV := formattedOutput(...)
  ```

### 2. String Formatting
- **`fmt.Sprintf`**: Unlike `Print` or `Println`, `Sprintf` returns the formatted string instead of writing to stdout. This separates the *logic* of formatting from the *action* of printing.
- **Verbs**:
    - `%.1f`: Rounds to 1 decimal place (e.g., `105.4`).
    - `%.2f`: Rounds to 2 decimal places (e.g., `105.43`).

### 3. Separation of Concerns
By returning strings, this function allows the caller to decide *how* to display the data (e.g., print to console, write to a file, send over a network), increasing the code's flexibility.

---

## Section: Function Deep Dive: getUserInput - 2026-02-14

The `getUserInput` function (lines 55-64) manages all console interactions to fetch data from the user.

### 1. Handling Input
- **Pointers `&`**: `fmt.Scan` requires a memory address (pointer) to the variable so it can write the user's input directly into that memory location.
  - `fmt.Scan(&investmentAmount)`
  - `fmt.Scan(&years)`
- **Behavior**: It pauses execution and waits for user input (separated by spaces or newlines).

### 2. Return Values
- **Signature**: `func getUserInput() (float64, float64)`
- **Efficiency**: Instead of using global variables or passing variables in to be modified, this function cleanly returns the two pieces of information needed by the program: `investmentAmount` and `years`.

### 3. Usage Pattern
The caller can assign these directly to new variables:
```go
investmentAmount, years := getUserInput()
```
This keeps the `inputs` isolated from the `calculations`.

---

## Section: Function Deep Dive: runInvestmentCalculator - 2026-02-14

The `runInvestmentCalculator` function (lines 66-74) serves as the **Control Flow Manager** or orchestrator of the entire application logic.

### 1. Orchestration
Instead of having all logic in `main`, this function acts as a high-level summary of what the program *does*:
1.  **Setup**: Defines the configuration constants (`inflationRate`).
2.  **Input**: Asks for data (`getUserInput`).
3.  **Process**: Chains the calculation functions (`calculateFutureValue` -> `calculateFutureRealValue`).
4.  **Output**: Formats and prints results (`formattedOutput`).

### 2. Benefits of this Structure
- **readability**: A developer can read this 8-line function and understand the entire program flow without needing to know *how* the math functions work.
- **Testing**: While `main` is hard to test, this function could potentially be tested if the inputs/outputs were further decoupled (e.g., passing reader/writers).
- **Refactoring**: If we wanted to change the inflation rate source or how we output the data (e.g., to a file), we change it here without touching the calculation logic.

### 3. Scope
Variables defined here (like `investmentAmount`, `futureValue`) are local to this function. They do not pollute the global package scope, which is a best practice.
