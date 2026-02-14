package maps

import "fmt"

func main() {
	// This code snippet is creating a map named `websites` in Go. The map is defined to have keys of type
	// `string` and values of type `string`. It is then initialized with two key-value pairs: "Google"
	// with the corresponding value "https://google.com" and "Amazon Web Services" with the corresponding
	// value "https://aws.com".
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	// The code snippet is demonstrating the use of a map in Go named `websites`. Here's what each line
	// does:
	// `fmt.Println(websites)` is printing the entire map named `websites` in Go. This will display all
	// the key-value pairs present in the map.
	fmt.Println(websites)
// `fmt.Println(websites["Amazon Web Services"])` is accessing the value associated with the key
// "Amazon Web Services" in the `websites` map and then printing that value. In this case, it will
// print the URL "https://aws.com" as it is the value associated with the key "Amazon Web Services" in
// the map.
	fmt.Println(websites["Amazon Web Services"])
	fmt.Println(websites["Google"])

	
	// The code snippet `websites["LinkedIn"] = "https://linkedin.com"` is adding a new key-value pair to
	// the existing map named `websites` in Go. It assigns the key "LinkedIn" with the corresponding value
	// "https://linkedin.com" in the map.
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	fmt.Println(websites["LinkedIn"])

	// The code snippet `delete(websites, "Google")` is deleting the key-value pair with the key "Google"
	// from the map named `websites` in Go. After this deletion, the map will no longer contain the key
	// "Google" and its corresponding value.
	delete(websites, "Google")
	fmt.Println(websites)

	// Maps vs Structs
	/* 
	There are to main defrences
	1) for maps we can use anything as a key and this gives us more flexibility since we are not stuck to just using human readable text as keys
	2) Maps solve a different problem, with strucs we have pre defined data structures, and we can not add new data to it, also we can not delete a key value from structs, also in structs we have restrictions 
	*/

}