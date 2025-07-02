package main

import (
	"errors"
	"fmt"
	"os"
)

// Store financial results in a struct for better organization.
type Financials struct {
	EBT    float64
	Profit float64
	Ratio  float64
}

func main() {
	// 1. Get User Input
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 2. Calculate Financials
	financials, err := calculateFinancials(revenue, expenses, taxRate)
	if err != nil {
		fmt.Println("Could not calculate financials:", err)
		return
	}

	// 3. Print and Store Results
	printFinancials(financials)
	err = storeFinancialsToFile(financials)
	if err != nil {
		fmt.Println("Failed to write results to file:", err)
	}
}

// getUserInput now also returns an error for better handling.
func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scanln(&userInput) // Use Scanln to avoid issues with leftover input.

	if userInput < 0 {
		return 0, errors.New("input must be a non-negative number")
	}
	if userInput == 0 && infoText == "Revenue: " {
		return 0, errors.New("revenue cannot be zero")
	}

	return userInput, nil
}

// calculateFinancials now returns the results and an error.
func calculateFinancials(revenue, expenses, taxRate float64) (Financials, error) {
	if revenue <= expenses {
		return Financials{}, errors.New("revenue must be greater than expenses")
	}

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	if profit == 0 {
		return Financials{}, errors.New("profit is zero, cannot calculate ratio")
	}

	ratio := ebt / profit

	return Financials{
		EBT:    ebt,
		Profit: profit,
		Ratio:  ratio,
	}, nil
}

// A dedicated function to print the results.
func printFinancials(f Financials) {
	fmt.Println("\n--- Financial Results ---")
	fmt.Printf("Earnings Before Tax (EBT): $%.2f\n", f.EBT)
	fmt.Printf("Profit: $%.2f\n", f.Profit)
	fmt.Printf("Ratio: %.2f\n", f.Ratio)
}

// A dedicated function to store the results in a single, clean file.
func storeFinancialsToFile(f Financials) error {
	resultsText := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", f.EBT, f.Profit, f.Ratio)
	return os.WriteFile("financials.txt", []byte(resultsText), 0644)
}