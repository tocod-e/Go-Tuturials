# Go Methods & Pointers: Constructors and Receivers

**Date:** 2026-02-14
**Topic:** Constructor Pattern, Value vs Pointer Receivers

## Code Analysis

```go
// 1. Constructor Pattern
func New(...) (*User, error) { ... }

// 2. Value Receiver (Copy)
func (u User) OutputUserDetails() { ... }

// 3. Pointer Receiver (Execute on Original)
func (u *User) ClearUserName() { ... }

// 4. Standalone Function
func OutputUserDetails(u *User) { ... }
```

### 1. The Constructor Pattern (`New`)

*   **Goal:** Enforce validation (inputs required) and set private fields (`createdAt`).
*   **Returns `*User`:** It returns a **pointer**. This is efficient because it avoids copying the struct when returning it. It also allows returning `nil` on error.
*   **Encapsulation:** This is the *only* way for external packages to get a valid `User` since fields are unexported.

### 2. Methods: Value vs. Pointer Receivers

This is a critical distinction in Go.

| Receiver Type | Syntax | Behavior | Memory Cost | usage |
| :--- | :--- | :--- | :--- | :--- |
| **Value Receiver** | `(u User)` | **COPIES** the entire struct. | High (Sizes > small). Copies all 80+ bytes. | Read-only operations on small structs. Logic that shouldn't mutate state. |
| **Pointer Receiver** | `(u *User)` | **POINTS** to the original struct. | Low (8 bytes). Copies only the memory address. | **Mutating state** (like `ClearUserName`). Efficiency for large structs. |

#### Scenario 1: `OutputUserDetails` (Value Receiver)
```go
u := User{...}
u.OutputUserDetails() 
// Go forces a COPY of 'u'. 
// If specific logic inside modified 'u', the original 'u' would NOT change.
```

#### Scenario 2: `ClearUserName` (Pointer Receiver)
```go
u.ClearUserName()
// Go passes the address of 'u'.
// The method sets fields to "". The original 'u' IS changed.
```

### 3. Methods vs. Standalone Functions

*   **Function:** `OutputUserDetails(&u)` - You **must** pass the exact type (`*User`). You cannot pass a `User` value if it expects a pointer.
*   **Method:** `u.OutputUserDetails()` - Go is smart! 
    *   If you have a `User` value and call a `*User` method, Go automatically does `(&u).Method()`.
    *   If you have a `*User` pointer and call a `User` (value) method, Go automatically dereferences `(*u).Method()`.

## Challenge: "What's Next?"

What happens if you change `ClearUserName` to use a **Value Receiver**?

```go
// CHANGED to Value Receiver
func (u User) ClearUserName() {
    u.firstName = ""
    u.lastName = ""
}

func main() {
    u, _ := New("Max", "Power", "1980")
    u.ClearUserName()
    u.OutputUserDetails() // What does this print?
}
```

<details>
<summary>Reveal Answer</summary>

**Output:** `Max Power 1980` (The name is NOT cleared!)

**Why?**
The method received a **copy** of the user. It successfully cleared the name on the **copy**, but that copy was destroyed immediately after the function finished. The original `u` remained untouched.
</details>
