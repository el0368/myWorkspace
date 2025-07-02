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

	var totalRevenue float64
	var totalExpense float64
	var taxPercent float64

	fmt.Println("Enter your revenue")
	fmt.Scan(&totalRevenue)

	fmt.Println("Enter your expenses")
	fmt.Scan(&totalExpense)

	fmt.Println("Enter your tax rate (%)")
	fmt.Scan(&taxPercent)

	taxRate := taxPercent / 100

	earningsBeforeTax, earningsAfterTax := Calc(totalRevenue, totalExpense, taxRate)
	profitRatio := earningsBeforeTax / earningsAfterTax

	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Profit Ratio: %.2f%%\n", profitRatio)


}


func Calc(totalRevenue, totalExpense, taxRate float64) (float64, float64) {

	earningsBeforeTax := totalRevenue - totalExpense
	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax * taxRate)

	return earningsBeforeTax, earningsAfterTax

}

