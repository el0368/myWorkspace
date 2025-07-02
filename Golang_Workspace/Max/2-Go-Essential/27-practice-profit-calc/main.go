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

	var totalRevenue float32
	var totalExpenses float32
	var taxPercent float32

	fmt.Println("Enter your revenue")
	fmt.Scan(&totalRevenue)

	fmt.Println("Enter your expenses")
	fmt.Scan(&totalExpenses)

	fmt.Println("Enter your tax rate (%)")
	fmt.Scan(&taxPercent)

	taxRate := taxPercent / 100

	earningsBeforeTax := totalRevenue - totalExpenses
	earningsAfterTax := earningsBeforeTax - (earningsBeforeTax * taxRate)

	profitRatio := (totalRevenue / earningsAfterTax) * 100

	fmt.Printf("Earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earnings after tax: %.2f\n", earningsAfterTax)
	fmt.Printf("Profit Ratio: %.2f%%\n", profitRatio)


}