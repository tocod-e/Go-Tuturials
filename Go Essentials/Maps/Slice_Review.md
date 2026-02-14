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
