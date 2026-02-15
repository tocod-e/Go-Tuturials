# Go Slices Practice: Append & Re-slicing

**Date:** 2026-02-14
**Topic:** Slices, Memory Allocation, and Append Mechanics

## Code Analysis

We are analyzing the behavior of slice `prices` as it undergoes modification, appending, and re-slicing.

```go
prices := []float64{10.99, 8.99}
prices[1] = 99.99
prices = append(prices, 5.99, 12.99, 29.99, 100.10)
prices = prices[1:]
prices = append(prices, discountPrices...)
```

### Memory & Capacity Table

| Line | Operation | Slice State (Len, Cap) | Underlying Array | Implementation Note |
| :--- | :--- | :--- | :--- | :--- |
| 7 | `:= {10.99, 8.99}` | **Len: 2, Cap: 2** | `[10.99, 8.99]` | Initial allocation. |
| 10 | `prices[1] = 99.99` | **Len: 2, Cap: 2** | `[10.99, 99.99]` | **In-place mutation**. Accessing by index changes the existing array. |
| 13 | `append(..., 4 items)` | **Len: 6, Cap: 6** (or 8) | `[10.99, 99.99, 5.99, ...]` | **Reallocation**. Current Cap (2) is insufficient. Go creates a **NEW** larger array, copies data, and appends new items. The old array is abandoned. |
| 15 | `prices[1:]` | **Len: 5, Cap: 5** (or 7) | `[99.99, 5.99, ...]` | **Re-slicing**. No new array. The slice header's **pointer shifts** forward by 1 element. Capacity decreases by 1. |
| 22 | `append(..., 3 items)` | **Len: 8, Cap: 8+** | `[..., 101.99, ...]` | **Potential Reallocation**. If the capacity from line 13/15 wasn't enough to hold 3 more, a new array is allocated again. |

### Escape Analysis (Stack vs Heap)

In Go, variables live on the **Stack** (fast, temporary) or **Heap** (slower, persistent, garbage collected).

*   `prices` (Slice Header) & Backing Array: **Escapes to Heap**.
    *   **Reason:** You pass `prices` to `fmt.Println()`.
    *   `fmt.Println` accepts arguments of type `interface{}` (any). When a value is passed to an interface, it often escapes because the compiler cannot determine at compile time how the function will use the memory (dynamic dispatch).
*   `discountPrices`: **Escapes to Heap** (Same reason: passed to `Println`).

## Key Takeaways

1.  **Append = Potential Copy**: If `len + new_items > cap`, Go forces a move to a new, larger memory address.
2.  **Slicing = Window Shift**: `prices[1:]` does *not* copy data. It just moves the "window" (pointer) forward.
3.  **Capacity**: Tracking capacity helps you predict when these expensive allocations happen.

## Challenge: "What's Next?"

Test your understanding with this variation. What does `final` look like?

```go
func challenge() {
    // 1. Create slice with Cap 4
    nums := make([]int, 2, 4) // [0, 0]
    nums[0] = 1
    nums[1] = 2
    
    // 2. Take a "view"
    view := nums[0:2] // [1, 2]
    
    // 3. Append to original 'nums' (Still fits in Cap 4!)
    nums = append(nums, 3) 
    
    // 4. Modify 'view'
    view[0] = 999 
    
    fmt.Println(nums) // ???
}
```

<details>
<summary>Reveal Answer</summary>

**Output:** `[999 2 3]`

**Why?**
Since `append(nums, 3)` fit within the capacity of `nums` (2 -> 3, Cap 4), it **did not reallocate**. `nums` and `view` still share **the same underlying array**. Modifying `view[0]` changed the memory that `nums[0]` also points to!
</details>
