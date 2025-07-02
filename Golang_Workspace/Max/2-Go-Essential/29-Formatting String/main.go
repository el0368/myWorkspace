package main

import (
	"fmt"
	"math"
)

func main() {

	var initialInvestment float64
	var interestRate float64
	var years float64

	fmt.Println("Enter the initial investment amount:")
	fmt.Scan(&initialInvestment)

	fmt.Println("Enter the annual interest rate (e.g., 5 for 5%):")
	fmt.Scan(&interestRate)

	fmt.Println("Enter the number of years for the deposit:")
	fmt.Scan(&years)

	rateDecimal := interestRate / 100.0
	finalAmount := initialInvestment * math.Pow((1+rateDecimal), years)
	totalInterest := finalAmount - initialInvestment

	fmt.Println("\n--- Investment Projection ---")
	fmt.Printf("Initial Investment: %.2f\n", initialInvestment)
	fmt.Printf("Annual Interest Rate: %.2f%%\n", interestRate)
	fmt.Printf("Investment Period: %.0f years\n", years)
	fmt.Printf("----------------------------------\n")
	fmt.Printf("Total Interest Earned: %.2f\n", totalInterest)
	fmt.Printf("Future Value: %.2f\n", finalAmount)
	fmt.Printf(`

	Once upon a time, an investor decided to grow their savings. 
	By depositing an initial amount and letting it earn interest over several years, they watched their money multiply. 
	With patience and the power of compounding, their investment blossomed into a rewarding future.

	`)
}