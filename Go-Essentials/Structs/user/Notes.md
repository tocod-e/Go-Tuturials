# Go Study Notes

---

## Section: Structs and Embedding

Structs are the building blocks of custom data types in Go. They allow you to group related fields together. Go also supports **embedding**, which works like composition (and sometimes inheritance).

### 1. Defining a Struct

```go
// type User struct {
    firstName string
    lastName  string
    birthDate string
    age       int
    createdAt time.Time
}
```

- **Definition**: Uses the `type [Name] struct` syntax.
- **Fields**: Contains named fields (properties) with specific types.
- **Visibility**:
  - Fields like `firstName` (lowercase) are **unexported** (private to the `user` package).
  - If you wanted to access them from `main.go`, they would need to be `FirstName` (uppercase).

### 2. Constructor Pattern

```go
func New(firstName, lastName, birthDate string) (*User, error) {
    if firstName == "" || ... {
        return nil, errors.New("...")
    }
    return &User{
        firstName: firstName,
        // ...
        createdAt: time.Now(),
    }, nil
}
```

- **Purpose**: Enforces validation (checking for empty strings) and sets default values (`createdAt`).
- **Return**: Returns a **pointer** to the struct (`*User`) to avoid copying large data structures and allow modification.

### 3. Methods (Value vs Pointer Receivers)

**Value Receiver (Read-Onlyish)**:

```go
func (u User) OutputUserDetails() { ... }
```

- Receives a **copy** of the user. Modifying `u` inside here won't affect the original.

**Pointer Receiver (Modifiable)**:

```go
func (u *User) ClearUserName() {
    u.firstName = ""
    u.lastName = ""
}
```

- Receives a **pointer** to the user. Modifying `u` **changes the original struct**.

### 4. Struct Embedding (Inheritance-like)

```go
type Admin struct {
    email    string
    password string
    User
}
```

- **Embedding**: By listing `User` without a name inside `Admin`, `Admin` automatically "inherits" the fields and methods of `User`.
- **Usage**: You can call `admin.OutputUserDetails()` directly, even though that method belongs to `User`.

### 5. Code Structure & Imports

- **`package user`**: Defines this code as part of the `user` package (reusable library).
- **`import`**:
  - `errors`: For creating error messages.
  - `fmt`: For printing.
  - `time`: For timestamps.
