# Go Study Notes

## Section: Bank Application - 2026-02-14

### Overview
This program implements a simple persistent banking system. It allows users to check their balance, deposit money, and withdraw money. The balance is persisted to a file named `balance.txt`, ensuring data survives between program runs.

### Line-by-Line Explanation

#### Package and Imports
```go
package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
```
- `package main`: Declares this file as part of the `main` package, making it an executable program.
- `import (...)`: imports necessary standard library packages.
    - `"errors"`: Used to create simple error messages (e.g., `errors.New(...)`).
    - `"fmt"`: **Format** package. Used for formatted I/O like printing to console (`Println`, `Printf`) and scanning input (`Scan`).
    - `"os"`: **Operating System** package. Used for file operations like `ReadFile` and `WriteFile`.
    - `"strconv"`: **String Conversion** package. Used to convert strings to numbers (`ParseFloat`) and vice versa.

#### Constants
```go
const balanceFilePath = "balance.txt"
```
- `const`: Declares a constant value.
- `balanceFilePath`: Stores the filename "balance.txt". Using a constant prevents typos and makes it easy to change the filename in one place.

#### The `main` Function
```go
func main() {
	helloClient()
	for { ... }
}
```
- Entry point of the program.
- `helloClient()`: Prints the welcome message.
- `for { ... }`: An infinite loop. This keeps the program running until the user explicitly chooses to exit (option 4).

**Inside the Loop:**
```go
		choice, err := printMenu()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
```
- `printMenu()`: Displays options and gets user input. Returns the user's `choice` (int) and potential `error`.
- `if err != nil`: Standard Go error handling. If an error occurred (e.g., invalid input), print it and `continue` to the next iteration of the loop, skipping the rest of the current iteration.

**Exit Condition:**
```go
		if choice == 4 {
			fmt.Println("Thank you for using the Bank. Goodbye!")
			break
		}
```
- `break`: Terminates the infinite `for` loop, causing the program to exit `main` and terminate.

**Switch Statement:**
```go
		switch choice {
		case 1:
			err := checkBalance()
            ...
		case 2:
			err := depositMoney()
            ...
		case 3:
			err := withdrawMoney()
            ...
		default:
			fmt.Println("Invalid choice...") 
		}
```
- `switch choice`: Route execution based on the integer value of `choice`.
- Each case calls a specific function (`checkBalance`, `depositMoney`, `withdrawMoney`).
- Error handling inside cases: If a function returns an error, a user-friendly message is printed to valid "System error".

#### File Operations

**`readBalanceFromFile`**
```go
func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFilePath)
    ...
```
- **Goal**: Reads the balance from disk.
- `os.ReadFile(balanceFilePath)`: Reads the entire file into a byte slice (`[]byte`).
- **Error Handling**:
    - `if os.IsNotExist(err)`: If the file doesn't exist (first run), return `0` balance and `nil` error (no error, just new account).
    - Other errors: Return `0` and the error.
- **Empty File Check**: `if len(data) == 0`. Returns 0 if file is empty.
- **Conversion**:
    - `string(data)`: Converts bytes to string.
    - `strconv.ParseFloat(balanceText, 64)`: Parses the string into a `float64`.
    - Returns the parsed balance.

**`writeBalanceToFile`**
```go
func writeBalanceToFile(balance float64) (error) {
	balanceText := fmt.Sprint(balance)
   	err :=	os.WriteFile(balanceFilePath,[]byte(balanceText), 0644)
    ...
}
```
- **Goal**: Saves the new balance to disk.
- `fmt.Sprint(balance)`: Converts the `float64` balance back to a `string`.
- `[]byte(balanceText)`: Casts the string to a byte slice for writing.
- `os.WriteFile(...)`: Writes data to file.
    - `0644`: File permission mode (Owner: Read/Write, Group: Read, Others: Read).
- Returns `nil` on success.

#### Core Functionality

**`checkBalance`**
- Calls `readBalanceFromFile`.
- Prints the verified balance.

**`depositMoney`**
```go
func depositMoney() (error) {
	var depositMoney float64
	fmt.Println("Enter the amount to deposit:")
	fmt.Scan(&depositMoney)
    ...
```
- `fmt.Scan(&depositMoney)`: Reads user input into the `depositMoney` variable. Note the `&` (address-of operator), passing a pointer so `Scan` can modify the variable directly.
- **Validation**: Checks `if depositMoney <= 0`. Returns error if invalid.
- **Logic**:
    1. Read current balance (`readBalanceFromFile`).
    2. Add deposit: `balance += depositMoney`.
    3. Save: `writeBalanceToFile(balance)`.
    4. Print confirmation.

**`withdrawMoney`**
- Similar flow to `depositMoney`.
- **Validation**: Checks `if withdrawMoney > balance`. Prevents overdrafts.
- **Logic**: Subtracts amount and saves.

#### Utility Functions

**`getUserChoice`**
```go
func getUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}
```
- dedicated helper to scan a single integer.

**`helloClient`**
- Simple void function to print the header.

**`printMenu`**
```go
func printMenu() (int , error){
    ...
	if choice < 1 || choice > 4 || choice != int(choice) {
		fmt.Println("Invalid choice...")
		return printMenu()
	} ...
}
```
- Displays visual menu.
- **Recursion**: If input is invalid, it calls `return printMenu()` (calls itself again). *Note: While recursive retry is one way, a loop usually preferred in production to avoid stack overflow, but manageable here.*
- `choice != int(choice)` check is technically redundant for an `int` type variable `choice`, but shows intent to validate integer input.

