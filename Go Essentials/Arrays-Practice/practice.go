package main

import "fmt"

// Bonus: Create a "Product" struct with title, id, price and create a dynamic list of products (at least 2 products). Then add a third product to the existing list of products.

type Product struct{
		id int
		title string
		price float64
		
	}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have, Output (print) that array in the command line.
	myarr := [3] string {"swimming", "biking", "coding"}
	fmt.Println("myarr: ",myarr)
	// The first element (standalone)

	firstElm := myarr[0]
	fmt.Println("First Elm: ",firstElm)
	// The second and third

	secondThird := myarr[1:]
	fmt.Println("Second and Third: ",secondThird, len(secondThird), cap(secondThird))

	// Create a Slice based on the first element that contains the first and second elements
	// first methode
	slice1_a := myarr[:2]
	fmt.Println("Slice1_a: ",slice1_a, len(slice1_a), cap(slice1_a))
	// second MEthode
	slice1_b := myarr[0:2]
	fmt.Println("slice1_b: ", slice1_b, len(slice1_b), cap(slice1_b))

	// RE-slice the slice from 3 and change it to contain second and last element of the orignal array
	slice2 := slice1_a[1:3]
	fmt.Println("slice2: ",slice2,len(slice2), cap(slice2)) 


	// Create a "dynamic array" that contains your course goals (at least 2 goals)
	fmt.Println("=============================")
	dynArr := []string {"LearnGo", "Theses"}
	fmt.Println("Orignal Dynamic Array: ", dynArr)
	


	// set the second goal to defferent one AND then add a third goal to that existing dynamic array 
	dynArr[1] = "Scripting"
	fmt.Println("Update dynamic: ", dynArr)

	dynArr = append(dynArr, "Coding")
	fmt.Println("Updated Dynamic Array: ", dynArr)
	
	
	// Bonus
	products := []Product{
		{
			id : 1,
			title: "PC",
			price: 329.99,
		},
		{
			id: 2,
			title: "Station",
			price: 449.99,
		},
	}
	fmt.Println("=============================")

	fmt.Println("Struct Products: ", products)
	newProduct := Product{
		id: 3,
		title: "NoteBook",
		price: 399.99,
	}
	products = append(products, newProduct)
	fmt.Println("Struct Products: ", products)
	fmt.Println("=============================")


	// Other Practice
	prices := []float64{10.99,8.99}
	fmt.Println(prices)
	fmt.Println(prices[0:1])
	prices[1] = 99.99
	fmt.Println(prices)

	prices = append(prices, 5.99, 12.99, 29.99, 100.10)
	fmt.Println(prices)
	prices = prices[1:]
	fmt.Println(prices)
	
	fmt.Println("=============================")
	discountPrices := []float64{101.99, 80.99, 20.99}
	fmt.Println(discountPrices)
	fmt.Println(prices)
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

}