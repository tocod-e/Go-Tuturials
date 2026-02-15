package main

import (
	"fmt"
	"example.com/structs/user"
)

func main(){
	userFirstName := getUserDate("Please enter your first name: ")
	userLastName := getUserDate("Please enter your last name: ")
	userBirthDate := getUserDate("Plase enter your birthdate (MM/DD/YYYY): ")


	var userApp *user.User

	userApp, err := user.New(userFirstName, userLastName, userBirthDate)
	if err != nil{
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("test@example.com", "test123")

	admin.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	
	
	userApp.OutputUserDetails()

	userApp.ClearUserName()

	userApp.OutputUserDetails()

}



func getUserDate(promptText string) string{
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)

	return  value
}