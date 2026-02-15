# Go Structs: User Definition Analysis

**Date:** 2026-02-14
**Topic:** Struct Definition, Memory Layout, and Encapsulation

## Code Analysis

```go
type User struct {
	firstName string
	lastName  string
	birthDate string
	age       int
	createdAt time.Time
}
```

### Memory & Layout Table

Structs in Go are essentially a sequence of fields in memory.

| Field | Type | Size (Approx on 64-bit) | Description |
| :--- | :--- | :--- | :--- |
| `firstName` | `string` | 16 bytes | (Ptr + Len) |
| `lastName` | `string` | 16 bytes | (Ptr + Len) |
| `birthDate` | `string` | 16 bytes | (Ptr + Len) |
| `age` | `int` | 8 bytes | Platform dependent (usually 64-bit) |
| `createdAt` | `time.Time` | 24 bytes | Complex struct (wall clock + monotonic) |
| **Total** | | **~80 bytes** | Contiguous block in memory. |

### Visibility (Encapsulation)

*   **Struct Name (`User`):** Capitalized -> **Exported**. Other packages can see the type `User`.
*   **Fields (`firstName`, etc.):** Lowercase -> **Unexported**. Other packages **cannot** access these fields directly (e.g., `u.firstName` will fail outside package `user`).
*   **Why?** This forces external code to use **Constructor Functions** (like `NewUser`) or **Methods** (like `u.FirstName()`) to interact with the data, allowing you to validate input or change implementation later without breaking code.

### Escape Analysis (Stack vs Heap)

*   If you create a user inside a function: `u := User{...}` -> **Stack** (Fast).
*   If you return a pointer to a user: `return &u` -> **Heap** (Shared, Garbage Collected).
*   Go's compiler decides this automatically.

## Challenge: "What's Next?"

Since the fields are private, how would you allow another package to create a **valid** `User`?

**Task:** Write a constructor function `NewUser`.

```go
// defined in package user
func NewUser(fName, lName, bDate string) (*User, error) {
    if fName == "" || lName == "" {
        return nil, errors.New("names are required")
    }
    return &User{
        firstName: fName,
        lastName:  lName,
        birthDate: bDate,
        createdAt: time.Now(),
    }, nil
}
```
