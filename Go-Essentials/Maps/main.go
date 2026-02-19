package main

import "fmt"

// type aliases
type floatMap map[string]float64
func (a floatMap) output()  {
	fmt.Println(a)
}


func main() {

	userNAmes := []string{}
	userNAmes = append(userNAmes, "MAx")
	userNAmes = append(userNAmes, "Manuel")
	fmt.Println(userNAmes)
	fmt.Println("===========")

	users := make([]string, 2)
	users = append(users, "Max")
	users = append(users, "Manuel")
	fmt.Println(users)
	users[0] = "Julie"
	users[1] = "Anna"
	fmt.Println(users)
	fmt.Println("===========")

	userNames := make([]string, 2, 5)
	userNames[0] = "Julia"
	userNames[1] = "Nic"
	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	fmt.Println(userNames)

	fmt.Println("===========")

	courseRating := map[string]float64{}
	courseRating["Go"] = 4.5
	courseRating["Java"] = 7.3

	fmt.Println(courseRating)

	coursesRate := map[string]float64{}
	coursesRate["Go"] = 4.5
	coursesRate["Java"] = 7.3

	fmt.Println(coursesRate)

	courseRat := make(map[string]float64, 3)
	courseRat["Python"] = 9.99
	courseRat["C Sharp"] = 7.3
	courseRat["Java Script"] = 8.9

	fmt.Println(courseRat)


	// Using the type Aliases
	productsRating := make(floatMap, 3)
	productsRating["PC"] = 4.3
	productsRating["Notebook"] = 7.9
	productsRating["Mac"] = 9.7

	productsRating.output()


	// for in arrays

	for index, value := range userNames{
		fmt.Println(index, value)
	}

}
