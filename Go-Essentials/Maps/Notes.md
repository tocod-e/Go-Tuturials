# Go Slices: Initialization and Append Review

This document reviews three key ways to initialize and work with **Slices** in Go.

## 1. Empty Slice Initialization

```go
userNAmes := []string{}
userNAmes = append(userNAmes, "MAx")
userNAmes = append(userNAmes, "Manuel")
// Result: ["MAx", "Manuel"]
```

*   **Starts as:** `[]` (Length: 0, Capacity: 0)
*   **Behavior:** Go automatically expands the array and adds new elements when you use `append`. Each `append` increases the length (and potentially capacity).

## 2. Slice with Pre-defined Length (The "Gotcha")

```go
users := make([]string, 2)      // Creates ["", ""] (Length 2)
users = append(users, "Max")    // Appends to the END -> ["", "", "Max"]
users = append(users, "Manuel") // Appends to the END -> ["", "", "Max", "Manuel"]
// Result logic:
// users[0] = "Julie" -> Updates the first existing element
// users[1] = "Anna"  -> Updates the second existing element
// Final: ["Julie", "Anna", "Max", "Manuel"]
```

*   **Starts as:** `["", ""]` (Length: 2, filled with zero-values).
*   **Explanation:** `make([]string, 2)` creates a slice that **already has 2 empty strings** in it.
    *   `append` always adds *after* the last element (length).
    *   If you access by index (`users[0] = ...`), you are updating the *existing* pre-allocated zero-value elements.

## 3. Slice with Length & Capacity

```go
userNames := make([]string, 2, 5)
userNames[0] = "Julia"
userNames[1] = "Nic"
userNames = append(userNames, "Max")    // Appends at index 2
userNames = append(userNames, "Manuel") // Appends at index 3
// Result: ["Julia", "Nic", "Max", "Manuel"]
```

*   **Starts as:** `["", ""]` (Length: 2, Capacity: 5)
*   **Explanation:**
    *   **Length (2):** Indices `0` and `1` are created and accessible immediately.
    *   **Capacity (5):** Reserves space in memory for 5 elements total, avoiding immediate re-allocation when appending.
    *   `append` starts adding at the first index *after* the length (index 2).

## Summary Strategy

*   **Fixed Size / Index Access:** If you want a list where you set items by specific index (`list[0]`), use `make` with a **Length**.
*   **Dynamic List:** If you plan to only use `append` to build a list, start with an empty slice (`[]string{}`) or `make` with **Length 0** (e.g., `make([]string, 0, 10)`).

---
## Section: Maps

Maps are Go's built-in associative data type (sometimes called *hashes* or *dicts* in other languages).

### 1. Map Initialization (Literal)

```go
courseRating := map[string]float64{}
// or
coursesRate := map[string]float64{}
```

*   **`map[string]float64`**: This defines the map's type.
    *   `[string]`: The **Key** type. Keys must be comparable (e.g., strings, ints).
    *   `float64`: The **Value** type. Values can be any type.
*   **`{}`**: This initializes an **empty map**.
    *   **zero value**: The zero value of a map is `nil`. A `nil` map behaves like an empty map when reading, but writing to it causes a runtime panic. Using `{}` or `make` initializes it so it's ready to write.

### 2. Map Initialization (make with capacity)

```go
courseRat := make(map[string]float64, 3)
```

*   **`make(Type, size)`**: The `make` function allocates and initializes a hash map.
*   **`3` (Capacity Hint)**:
    *   This argument is optional.
    *   It tells the runtime to allocate enough space for *at least* 3 elements initially.
    *   **Performance**: Providing a size hint reduces the overhead of memory re-allocation as the map grows. If you know roughly how many elements you'll have, use this.

### 3. Adding and Retrieving Elements

```go
courseRating["Go"] = 4.5
```

*   **Syntax**: `map[key] = value` uses the square bracket syntax, just like arrays/slices, but with the key type inside.
*   **Dynamic Growth**: Unlike arrays, maps grow dynamically to accommodate new key-value pairs.

### 4. Printing Maps

```go
fmt.Println(courseRat)
// Output example: map[C Sharp:7.3 Java Script:8.9 Python:9.99]
```

*   **Unordered**: Iteration order over maps is **not specified** and is not guaranteed to be the same from one iteration to the next. Even if it looks sorted in the output, you should never rely on key order.

---
## Section: Type Aliases and Custom Types

Go allows you to define new types based on existing ones. This is powerful for adding behavior (methods) to built-in types like maps or slices.

### 1. Defining a Custom Type

```go
type floatMap map[string]float64
```

*   **`type`**: The keyword used to introduce a new type.
*   **`floatMap`**: The name of the new type.
*   **`map[string]float64`**: The *underlying type*. `floatMap` has all the properties of a map, but it is a distinct type in the type system.
*   **Why do this?** exact type matching is required in Go. But more importantly, **you can attach methods to your own types**, effectively treating them like objects.

### 2. Attaching Methods to Custom Types

```go
func (a floatMap) output() {
    fmt.Println(a)
}
```

*   **Receiver `(a floatMap)`**: This special parameter before the function name links the function `output` to the type `floatMap`.
    *   `a`: The variable name for the instance of the type (like `this` or `self` in other languages, but explicit).
    *   `floatMap`: The type this method belongs to.
*   **Usage**: You can now call `myMap.output()` on any variable of type `floatMap`.

### 3. Using the Custom Type

```go
productsRating := make(floatMap, 3)
productsRating["PC"] = 4.3
// ...
productsRating.output()
```

*   **`make(floatMap, 3)`**: You use the new type name with `make`.
*   **Method Call**: `productsRating.output()` calls the attached function.
*   **Casting**: If you had a regular `map[string]float64`, you would need to cast it to `floatMap` to use the method: `floatMap(regularMap).output()`.
---
## Section: Range Loop on Slices

Go provides a powerful keyword `range` to iterate over collections like slices or maps.

### 1. Basic Range Syntax

```go
for index, value := range userNames {
    fmt.Println(index, value)
}
```

*   **`for`**: This is the only looping construct in Go.
*   **`range userNames`**: The `range` keyword iterates over the `userNames` slice.
    *   In each iteration, it returns **two values**.
*   **`index, value`**:
    *   **`index` (1st return)**: The current position in the slice (starting at 0).
    *   **`value` (2nd return)**: A **copy** of the element at that index.
*   **Definition**: `:=` is used because `index` and `value` are new variables created for the loop scope.

### 2. Ignoring Values

If you only need the value and not the index, you can use the blank identifier `_` to discard the index:

```go
for _, value := range userNames {
    fmt.Println(value)
}
```

Conversely, if you only need the index, you can omit the second variable entirely:

```go
for index := range userNames {
    // ...
}
```
