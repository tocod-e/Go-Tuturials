# Go JSON & File I/O: Struct Tags Practice

**Date:** 2026-02-14
**Topic:** Struct Tags, JSON Marshalling, and File Permissions

## Code Analysis

```go
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
```

### 1. Struct Tags (`json:"..."`)

*   **Purpose:** Metadata attached to fields.
*   **Behavior:** The `encoding/json` package uses "reflection" to read these tags.
*   **Mapping:**
    *   `Title` (Go) -> `"title"` (JSON Key)
    *   `CreatedAt` (Go) -> `"created_at"` (JSON Key)
*   **Requirement:** Fields **MUST** be exported (Capitalized) for `json.Marshal` to see them. If you made `title` lowercase, the JSON package would ignore it!

### 2. Constructor (`New`)

```go
func New(title, content string) (Note, error)
```

*   **Return Type:** Returns `Note` (Value), not `*Note` (Pointer).
*   **Why?** `Note` is small (2 strings + 1 time). Copying it is cheap. If it had large arrays or many fields, a pointer `*Note` would be better memory-wise.

### 3. Saving to File (`Save`)

```go
func (note Note) Save() error {
    // ...
    json, err := json.Marshal(note)
    return os.WriteFile(fileName, json, 0644)
}
```

*   **Marshal:** Converts the struct to a JSON byte slice `[]byte`.
*   **WriteFile:** Writes data to disk.
*   **Permissions `0644`:** 
    *   **Owner (User):** Read + Write (6)
    *   **Group:** Read (4)
    *   **Others:** Read (4)
    *   Standard permission for readable text files.

## Challenge: "What's Next?"

What happens if you remove the JSON tags?

```go
type Note struct {
    Title   string
    Content string
}
// json.Marshal(Note{Title: "Hi", Content: "Body"})
```

<details>
<summary>Reveal Answer</summary>

**Output:** `{"Title":"Hi","Content":"Body"}`

**Why?**
Without tags, Go defaults to using the **exact field name** as the JSON key. The tags allow you to convert `Title` (Go style) to `title` (JSON/JavaScript style).
</details>
