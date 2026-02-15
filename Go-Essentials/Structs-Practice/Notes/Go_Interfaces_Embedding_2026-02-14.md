# Go Interfaces: Embedding & Polymorphism

**Date:** 2026-02-14
**Topic:** Interface Definition, Embedding, and Composition

## Code Analysis

```go
type saver interface {
	Save() error
}

type outable interface {
	saver      // Embedded Interface
	Display()
}
```

### 1. Interface Embedding

Just like you can embed structs, you can **embed interfaces**.

*   `saver`: Requires **1 method** (`Save`).
*   `outable`: Requires **2 methods** (`Save` + `Display`).
    *   It "inherits" the requirements of `saver`.
    *   Any type that satisfies `outable` **automatically** satisfies `saver`.

### 2. Polymorphism in Action

```go
func outputData(data outable) error {
	data.Display()    // ok: data is outable
	return saveData(data) // ok: data is outable, which implies it's also a saver!
}

func saveData(data saver) error { ... }
```

*   **Substitutability:** You can pass an `outable` to a function expecting a `saver` (because it has the `Save` method).
*   **Restriction:** You *cannot* pass a mere `saver` to a function expecting an `outable` (it might be missing `Display`).

### 3. Implementation

In your `note` and `todo` packages, your structs implement these methods implicitly.

*   `Note`: Has `Save()` and `Display()` -> Satisfies `outable`.
*   `Todo`: Has `Save()` and `Display()` -> Satisfies `outable`.

This allows `main.go` to treat them identically using the `outputData` function, despite them being completely different structs.

## Challenge: "What's Next?"

Imagine we have a `Log` struct that only has a `Save()` method (no display).

```go
type Log struct { msg string }
func (l Log) Save() error { return nil }

func main() {
    l := Log{"Error"}
    saveData(l)   // Works?
    outputData(l) // Works?
}
```

<details>
<summary>Reveal Answer</summary>

1.  `saveData(l)`: **Works**. `Log` has `Save()`, so it is a `saver`.
2.  `outputData(l)`: **Fails**. `Log` is missing `Display()`, so it is **NOT** `outable`. Go will complain at compile time!
</details>
