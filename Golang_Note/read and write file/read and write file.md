Of course. Here is a complete Go program that simulates a basic ATM. It uses a single file named `balance.txt` to store and manage the account balance.

This example combines file reading and writing to perform the essential ATM functions.

---

### ATM Simulation Program


```Go
package main

import (
	"fmt"
	"os"
	"strconv"
)

// The name of our "database" file
const balanceFile = "balance.txt"

// readBalance reads the current balance from the file.
func readBalance() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	// If the file doesn't exist, we assume a starting balance of 0.
	if os.IsNotExist(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err // For other errors, return the error.
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

// writeBalance overwrites the file with the new balance.
func writeBalance(balance float64) error {
	// Format the balance to a string with 2 decimal places.
	data := []byte(fmt.Sprintf("%.2f", balance))
	// Write the data to the file, overwriting it.
	return os.WriteFile(balanceFile, data, 0644)
}

func main() {
	fmt.Println("Welcome to the Go ATM!")

	for {
		fmt.Println("\n--------------------")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Exit")
		fmt.Print("Please choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// --- Check Balance (Read) ---
			balance, err := readBalance()
			if err != nil {
				fmt.Println("Error checking balance:", err)
				continue
			}
			fmt.Printf("Your current balance is: $%.2f\n", balance)

		case 2:
			// --- Deposit (Read -> Calculate -> Write) ---
			fmt.Print("Enter amount to deposit: ")
			var amount float64
			fmt.Scanln(&amount)

			if amount <= 0 {
				fmt.Println("Invalid deposit amount.")
				continue
			}

			balance, err := readBalance()
			if err != nil {
				fmt.Println("Error reading balance for deposit:", err)
				continue
			}

			newBalance := balance + amount
			err = writeBalance(newBalance)
			if err != nil {
				fmt.Println("Error updating balance:", err)
				continue
			}
			fmt.Printf("Deposit successful. New balance: $%.2f\n", newBalance)

		case 3:
			// --- Withdraw (Read -> Validate -> Calculate -> Write) ---
			fmt.Print("Enter amount to withdraw: ")
			var amount float64
			fmt.Scanln(&amount)

			if amount <= 0 {
				fmt.Println("Invalid withdrawal amount.")
				continue
			}

			balance, err := readBalance()
			if err != nil {
				fmt.Println("Error reading balance for withdrawal:", err)
				continue
			}

			if amount > balance {
				fmt.Printf("Insufficient funds. You only have $%.2f\n", balance)
				continue
			}

			newBalance := balance - amount
			err = writeBalance(newBalance)
			if err != nil {
				fmt.Println("Error updating balance:", err)
				continue
			}
			fmt.Printf("Withdrawal successful. New balance: $%.2f\n", newBalance)

		case 4:
			// --- Exit ---
			fmt.Println("Thank you for using the Go ATM. Goodbye!")
			return // Exits the main function and the program

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
```

---

### How It Works

1. **The "Database" (`balance.txt`)**
    
    - This program relies on a single text file, `balance.txt`, to store the current balance. If the file contains `150.75`, that's the account balance.
        
2. **Reading from the File (`readBalance`)**
    
    - Before any operation, the program calls `readBalance()`.
        
    - This function uses **`os.ReadFile()`** to get the contents.
        
    - It handles a special case: if the file doesn't exist yet (e.g., the first time you run it), it assumes a balance of `$0`.
        
    - The text from the file is converted into a number (`float64`) using **`strconv.ParseFloat()`** so we can do math with it.
        
3. **Writing to the File (`writeBalance`)**
    
    - After a deposit or withdrawal, the program calls `writeBalance()`.
        
    - This function takes the new balance, formats it into a string with two decimal places (like money), and converts it to bytes.
        
    - It then uses **`os.WriteFile()`** to completely **overwrite** `balance.txt` with the new balance. This ensures the file is always up-to-date.
        
4. **Main Loop**
    
    - The `main` function runs in an infinite `for` loop, constantly showing the menu.
        
    - Based on the user's choice, it calls the read/write functions to perform the correct actions.
        
    - Choosing `4` executes a `return`, which exits the `main` function and gracefully stops the program.