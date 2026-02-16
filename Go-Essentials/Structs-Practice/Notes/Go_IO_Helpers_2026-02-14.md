# Go Input/Output: Helpers & bufio

**Date:** 2026-02-14
**Topic:** Reading User Input, Buffer Handling, and String Manipulation

## Code Analysis

### 1. The `getUserInput` Helper

```go
func getUserInput(prompt string) string {
    fmt.Printf("%v ", prompt)
    reader := bufio.NewReader(os.Stdin)
    text, err := reader.ReadString('\n')
    // ... logic to handle empty reads ...
    return strings.TrimSpace(text)
}
```

#### Why `bufio.NewReader`?
*   **Standard `fmt.Scan`**: Stops reading at the first space. Bad for sentences (like "Buy Milk").
*   **`bufio` (Buffered I/O)**: Allows reading until a specific character (like `\n` - the "Enter" key). This captures the full sentence.

#### Why `strings.TrimSpace`?
*   When you type "Hello" and hit Enter, the program actually receives `"Hello\n"` (on Linux/Mac) or `"Hello\r\n"` (on Windows).
*   **TrimSpace** removes these invisible "whitespace" characters from the start and end, giving you clean data (`"Hello"`).

#### The "Double Read" Logic
```go
if strings.TrimSpace(text) == "" && err == nil {
    text, err = reader.ReadString('\n')
}
```
*   **Problem:** Sometimes, previous input operations leave a "phantom" newline character in the input buffer. The logic sees this as an empty line immediately.
*   **Fix:** If we read an empty line when we expected text, we try reading **one more time** to consume the actual input.

### 2. Polymorphic Helpers (`outputData`)

```go
func outputData(data outable) error {
    data.Display()         // 1. Show it
    return saveData(data)  // 2. Save it
}
```
*   **DRY (Don't Repeat Yourself):** Instead of writing "Display then Save" logic twice (once for Note, once for Todo), we write it once.
*   **Flexibility:** If we add a `DiaryEntry` type later, we don't need to change `outputData`. As long as `DiaryEntry` satisfies `outable`, it just works.

## Mechanics of `saveData`

```go
func saveData(data saver) error { ... }
```
It takes a `saver` interface. It doesn't care *what* the data is, only that it has a `.Save()` method. This is the power of Go interfaces: **Decoupling** "what needs to be done" from "how it is done".

## Challenge: "What's Next?"

How could you modify `getUserInput` to **validate** that the user didn't just type spaces?

<details>
<summary>Reveal Answer</summary>

Currently, `strings.TrimSpace` handles this implicitly!
If the user types "   ", `TrimSpace` turns it into `""`.
You check this in `New()`: `if title == "" ... return error`.

So the validation happens in the **Factory**, not the Input helper. This is good design (separation of concerns).
</details>
