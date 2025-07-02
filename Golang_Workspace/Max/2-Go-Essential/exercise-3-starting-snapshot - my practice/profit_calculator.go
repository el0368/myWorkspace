package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Goals
// 1) Validate user input
// 	=> Show error message & exit if invalid input is provided
// - No negative numbers
// - Not 0
// 2) Store calculated result into file

const ebtTextFile = "ebt.txt"
// const expTextFile = "expenses.txt"
const profitTextFile = "profit.txt"
const ratioTextFile = "ratio.txt"


func main() {
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	taxRate, err := getUserInput("Tax Rate: ")
	
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	calculateFinancials(revenue, expenses, taxRate)

	// writeToFile(ebt, profit, ratio)

	ReadFromFiles(ebtTextFile, profitTextFile, ratioTextFile)
}

func ReadFromFiles(ebt, profit, ratio string)  {

	ebtData, err := os.ReadFile(ebtTextFile)

	if err != nil {
		fmt.Println("Error reading EBT file:", err)
		os.Exit(1)
	}
	// expData, _ := os.ReadFile(expTextFile)
	profitData, err  := os.ReadFile(profitTextFile)
	if err != nil {
		fmt.Println("Error reading Profit file:", err)
		os.Exit(1)
	}
	ratioData, err := os.ReadFile(ratioTextFile)
	if err != nil {
		fmt.Println("Error reading Ratio file:", err)
		os.Exit(1)
	}

	ebtReader := string(ebtData)
	// expReader := string(expData)
	profitRead := string(profitData)
	ratioReader := string(ratioData)

	ebtReaderOut, err := strconv.ParseFloat(ebtReader, 64)
	if err != nil {
		fmt.Println("Error parsing EBT value:", err)
		os.Exit(1)
	}
	// expReaderOut, _ := strconv.ParseFloat(expReader, 64)
	profitReaderOut, err := strconv.ParseFloat(profitRead, 64)
	if err != nil {
		fmt.Println("Error parsing Profit value:", err)
		os.Exit(1)
	}
	ratioReaderOut, err := strconv.ParseFloat(ratioReader, 64)
	if err != nil {
		fmt.Println("Error parsing Ratio value:", err)
		os.Exit(1)
	}

	fmt.Printf("%.1f\n", ebtReaderOut)
	// fmt.Printf("%.1f\n", expReaderOut)
	fmt.Printf("%.1f\n", profitReaderOut)
	fmt.Printf("%.3f\n", ratioReaderOut)
}

// func writeToFile(ebt, profit, ratio float64) {
// 	ebtFile := fmt.Sprint(ebt)
// 	os.WriteFile("ebt.txt", []byte(ebtFile), 0644)
// 	profitFile := fmt.Sprint(profit)
// 	os.WriteFile("profit.txt", []byte(profitFile), 0644)
// 	ratioFile := fmt.Sprint(ratio)
// 	os.WriteFile("ratioFile.txt", []byte(ratioFile), 0644)

// }

func calculateFinancials(revenue, expenses, taxRate float64)  {
	ebt := revenue - expenses
	ebtFile := fmt.Sprint(ebt)
	os.WriteFile("ebt.txt", []byte(ebtFile), 0644)

	// exp := expenses
	// expFile := fmt.Sprint(exp)
	// os.WriteFile("expenses.txt", []byte(expFile), 0644)

	profit := ebt * (1 - taxRate/100)
	profitFile := fmt.Sprint(profit)
	os.WriteFile("profit.txt", []byte(profitFile), 0644)

	ratio := ebt / profit
	ratioFile := fmt.Sprint(ratio)
	os.WriteFile("ratio.txt", []byte(ratioFile), 0644)
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0{
		return 0, errors.New("value must be a positive number")
	}

	return userInput, nil
}
