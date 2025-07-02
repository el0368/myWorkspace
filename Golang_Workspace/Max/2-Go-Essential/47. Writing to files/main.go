package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	
	menuDisplay()

}

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil{
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}



func menuDisplay() {
	fmt.Print(`
	Welcome to Go Bank!
	What do you want to do?
	1. Check balance
	2. Deposit money
	3. Withdraw money
	4. Exit
	`)
	menuChoice()
}

func menuChoice() {
	var myBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("Warning:", err)
		panic("Can't continue, sorry")
	}
	var myDeposit float64
	var myWithdraw float64

	for {
		var choice string
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case "1":
			remainingBalance := myBalance
			fmt.Printf("Balance remaining: %v\n", remainingBalance)
		case "2":
			fmt.Println("Input your deposit")
			fmt.Scan(&myDeposit)
			myBalance += myDeposit
			fmt.Printf("Deposit: %v\n", myDeposit)
			fmt.Printf("Balance remaining: %v\n", myBalance)

			writeBalanceToFile(myBalance)
		case "3":
			fmt.Println("Withdraw your cash")
			fmt.Scan(&myWithdraw)

			
			if myWithdraw > myBalance {
				fmt.Println("Insufficient funds. Withdrawal denied.")
			} else {
				myBalance -= myWithdraw
				fmt.Printf("Withdraw: %v\n", myWithdraw)
				fmt.Printf("Balance remaining: %v\n", myBalance)
			}

			writeBalanceToFile(myBalance)

		case "4":
			fmt.Println("Exiting. Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
