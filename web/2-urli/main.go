package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	reslut, _ := url.Parse(myurl)

	fmt.Println(reslut.Scheme)
	fmt.Println(reslut.Host)
	fmt.Println(reslut.Path)
	fmt.Println(reslut.Port())
	fmt.Println(reslut.RawQuery)

	qparams := reslut.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	

	for key, val := range qparams {
		fmt.Println("Param is: ", key, " Valu is: ", val)
	}
	fmt.Println(qparams["coursename"])

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
