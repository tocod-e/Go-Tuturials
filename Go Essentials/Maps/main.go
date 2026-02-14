package main

import "fmt"

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

}