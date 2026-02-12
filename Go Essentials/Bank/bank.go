package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)


const balanceFilePath = "balance.txt"
func main(){
	helloClient()
	for {
		choice, err := printMenu()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		if choice == 4 {
			fmt.Println("Thank you for using the Bank. Goodbye!")
			break
		}
		switch choice {
		case 1:
			err := checkBalance()
			if err != nil {
				fmt.Println("System error: We could not retrieve your balance at this time. Please try again later.")
			}
		case 2:
			err := depositMoney()
			if err != nil {
				fmt.Println("System error: We could not process your deposit at this time. Please try again later.")
			}
		case 3:
			err := withdrawMoney()
			if err != nil {
				fmt.Println("System error: We could not process your withdrawal at this time. Please try again later.")
			}
		default:
			fmt.Println("Invalid choice. Please choose a number between 1 and 4.") 
		}	
	}
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, nil
		}
		return 0, err
	}

	if len(data) == 0 {
		return 0, nil
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 0,errors.New("Error by parsing balance")
	}	
	return balance , nil
}

func writeBalanceToFile(balance float64) (error) {
	balanceText := fmt.Sprint(balance)
   	err :=	os.WriteFile(balanceFilePath,[]byte(balanceText), 0644)
	if err != nil {
		fmt.Println("Error by writing balance to file:", err)
	}
	fmt.Println("Balance updated successfully.")
	return nil
}


func checkBalance() (error) {
	balance, err := readBalanceFromFile()
	if err != nil {
		return errors.New("Can not read Balance")
	}
	fmt.Println("Your current balance is:", balance)
	return nil
}

func depositMoney() (error) {
	var depositMoney float64
	fmt.Println("Enter the amount to deposit:")
	fmt.Scan(&depositMoney)
	
	if depositMoney <= 0 {
		return errors.New("Deposit amount must be greater than zero.")
		
	}else {
		balance, err := readBalanceFromFile()
		if err != nil {
			return errors.New("Can not read Balance")
		}
		balance += depositMoney
		err = writeBalanceToFile(balance)
		if err != nil {
			return errors.New("Can not write to balance")
		}

		fmt.Printf("You deposit: %v  Your new balnce is: %v \n", depositMoney, balance)
		return nil
	}
}

func withdrawMoney() error{
	balance, err := readBalanceFromFile()
	if err != nil {
		return errors.New("Can not read Balance")
	}	
	var withdrawMoney float64
	fmt.Println("Enter the amount to withdraw:")
	fmt.Scan(&withdrawMoney)
	if withdrawMoney > balance {
		return errors.New("Insufficient funds, cannot withdraw more than the current balance. Your current balance is: " + fmt.Sprint(balance))
	} else {
		fmt.Println("You Withdraw: ", withdrawMoney)
		balance -= withdrawMoney
		err = writeBalanceToFile(balance)
		if err != nil {
			return errors.New("Can not write to balance")
		}
		fmt.Printf("You withdaw: %v  Your new balance is: %v \n", withdrawMoney, balance)
	}
	return nil
}
 
func getUserChoice() int {
	var choice int
	fmt.Scan(&choice)
	return choice
}

func helloClient(){
	fmt.Println("====================")
	fmt.Println("Welcome to the Bank!")
}
func printMenu() (int , error){
	fmt.Println("====================")
	fmt.Println("What do you want to do? Chose a number to perform the operation:")
	fmt.Println("1. Check Balance?")
	fmt.Println("2. Deposit Money?")
	fmt.Println("3. Withdraw Money?")
	fmt.Println("4. Exit?")
	fmt.Println("====================")

	var choice int
	choice = getUserChoice()
	if choice < 1 || choice > 4 || choice != int(choice) {
		fmt.Println("Invalid choice. Please choose a number between 1 and 4.")
		return printMenu()
	}else {
		fmt.Println("You chose option: ", choice)
		return choice, nil
	}
	
}