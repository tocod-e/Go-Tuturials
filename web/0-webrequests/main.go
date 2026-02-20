package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {
	fmt.Println("Web Requests")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Responce is of type: %T\n", response)
	fmt.Println("Status code:", response.StatusCode)
	defer response.Body.Close()
	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
