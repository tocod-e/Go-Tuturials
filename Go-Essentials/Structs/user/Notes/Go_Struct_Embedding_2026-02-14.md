# Go Struct Embedding (Composition)

**Date:** 2026-02-14
**Topic:** Struct Embedding, Composition vs Inheritance, Promoted Fields

## Code Analysis

```go
type Admin struct {
	email    string
	password string
	User     // Embedded Struct (Anonymous Field)
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "ADMIN",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}
```

### 1. Struct Embedding (Not Inheritance!)

Go does **not** have classes or inheritance (like Java/Python `extends`). Instead, it uses **Composition**.

*   **Syntax:** putting `User` inside `Admin` without a field name makes it an **Embedded Field**.
*   **Result:** All fields and methods of `User` are **promoted** to `Admin`.
    *   You can call `admin.firstName` directly (if exported).
    *   You can call `admin.OutputUserDetails()` directly!

### 2. Initialization (Composite Literal)

When creating the struct, you must be explicit about the inner struct:

```go
Admin{
    User: User{ ... }, // You must initialize the embedded struct explicitly
}
```

### 3. Memory Layout

The `Admin` struct is just a slightly larger block of memory.

| Field | Type | Function |
| :--- | :--- | :--- |
| `email` | `string` | Admin specific |
| `password` | `string` | Admin specific |
| `User` | `struct` | **The entire User struct is embedded here directly.** |

### 4. Shadowing

If `Admin` also had a method named `OutputUserDetails()`, it would **shadow** (hide) the one from `User`. 
*   `admin.OutputUserDetails()` -> Calls Admin's version.
*   `admin.User.OutputUserDetails()` -> Calls User's version explicitly.

## Challenge: "What's Next?"

Since `Admin` has an embedded `User`, and `User` has a pointer receiver method `ClearUserName()`, does `Admin` automatically get that method too?

```go
func main() {
    a := NewAdmin("admin@test.com", "1234")
    a.ClearUserName() // Does this work?
    a.OutputUserDetails()
}
```

<details>
<summary>Reveal Answer</summary>

**YES!** 

Method promotion works for pointer receivers too. calling `a.ClearUserName()` is automatically translated by Go to `a.User.ClearUserName()`. The admin's inner user name will be cleared.
</details>
