package main

import (
	"fmt"
)

func main() {

	// Profit calculator
	// ask for revenue, expenses & tax rate

    // Earning before tax
    // Calculate ratio ebt/profit
    // Earning after tax

	// var totalRevenue float64
	// var totalExpense float64
	// var taxPercent float64

	totalRevenue := userInput("Total Revenue: ")
	totalExpense := userInput("Total Expense: ")
	taxPercent := userInput("Tax Percent: ")

	taxRateDecimal := taxPercent / 100

	earningsBeforeTax, earningsAfterTax := calc(totalRevenue, totalExpense, taxRateDecimal)
	textDisplay(earningsBeforeTax, earningsAfterTax)


}

func userInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		panic(err) 
	}
	return userInput
}

func calc(totalRevenue, totalExpense, taxRateDecimal float64) (float64, float64) {
	earningsBeforeTax := totalRevenue - totalExpense
	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax * taxRateDecimal)

	return earningsBeforeTax, earningsAfterTax

}


func textDisplay(earningsBeforeTax, earningsAfterTax float64) {
	var profitRatio float64
	if earningsAfterTax != 0 {
		profitRatio = earningsBeforeTax / earningsAfterTax

	}
	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Profit Ratio: %.2f%%\n", profitRatio)

}
