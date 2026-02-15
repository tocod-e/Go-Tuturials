# Go Application Flow & Error Handling

**Date:** 2026-02-14
**Topic:** Control Flow, Error Handling Patterns, and Interface Usage

## Code Analysis

```go
func main() {
    // 1. Input & Creation (Note)
    title, content := getNoteDate()
    userNote, err := note.New(title, content)
    if err != nil {
        fmt.Println(err)
        return // Stop execution if creation fails
    }

    // 2. Output & Save (Polymorphic)
    err = outputData(userNote)
    if err != nil { return }

    // 3. Input & Creation (Todo)
    // ... similar pattern for Todo ...
}
```

### 1. The "Guard Clause" Pattern (Error Handling)

In Go, control flow usually goes line-by-line. When an error occurs, we handle it immediately using an `if` block.

```go
if err != nil {
    fmt.Println(err)
    return
}
// If we reach here, we know 'err' is nil and it's safe to proceed.
```
*   **Why?** This keeps the "happy path" (the main logic) usage unindented and readable. We "guard" against errors early.

### 2. Factory Pattern usage

`note.New(...)` is acting as a **Factory**.
*   It handles validation logic internally.
*   It ensures we only get a usable object if the inputs are valid.

### 3. Polymorphism in `main`

Notice how `outputData` is called twice:

1.  `outputData(userNote)` -> `userNote` is a `note.Note`.
2.  `outputData(userTodo)` -> `userTodo` is a `todo.Todo`.

The `main` function doesn't need to know *how* to save a note vs a todo. It just asks `outputData` to handle it. This makes `main` very clean and decoupled from the specific implementation details of saving files.

## Challenge: "What's Next?"

What happens if `outputData` fails (returns an error), but we **don't** check it?

```go
outputData(userNote) // No 'err' variable captured
// Program continues...
```

<details>
<summary>Reveal Answer</summary>

**Silent Failure.** 
The program would continue running as if nothing happened. The user would see "Saving succeeded!" (printed inside `outputData` before return) or "Saving failed" (printed inside `saveData`), but if `main` had subsequent logic that *depended* on the save being successful, that subsequent logic would run on invalid state.

**Best Practice:** ALWAYS check errors. `_ = outputData(...)` communicates "I am ignoring this intentionally".
</details>
