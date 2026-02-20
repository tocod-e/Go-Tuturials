package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb vedio -LCO")
	PerformGetRequests()
}

func PerformGetRequests() {
	const myurl = "http://localhost:8000/get"
	responce, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()
	fmt.Println("Status code: ", responce.StatusCode)
	fmt.Println("Content length is: ", responce.ContentLength)

	content, _ := io.ReadAll(responce.Body)
	
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte Count is: ",byteCount)

	fmt.Println(responseString.String())
	//fmt.Println(string(content))

}
