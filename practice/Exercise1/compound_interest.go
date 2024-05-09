package main

import (
	"fmt"
	"math"
)

// CompoundInterest calculates compound interest per annum
func CompoundInterest(principal float64, rate float64, years int) float64 {
	// Calculate compound interest using the formula: A = P * (1 + r/n)^(nt)
	// A = final amount, P = principal amount, r = annual interest rate (in decimal), n = number of times interest applied per time period, t = time period in years
	// For simplicity, we assume interest is compounded annually (n = 1)
	// Formula simplifies to A = P * (1 + r)^t
	// rate = rate / 100
	// principalforYearOne := principal * 12
	// AmountAfterInterest := principalforYearOne * rate
	rate = rate / 100

	amount := principal * math.Pow(1+rate, float64(years))
	return amount
}

func main() {
	var principal, rate float64
	var years int

	// Prompt for principal amount
	fmt.Print("Enter the principal amount: ")
	fmt.Scanln(&principal)

	// Prompt for annual interest rate
	fmt.Print("Enter the annual interest rate (in decimal): ")
	fmt.Scanln(&rate)

	// Prompt for number of years
	fmt.Print("Enter the number of years: ")
	fmt.Scanln(&years)

	// Calculate compound interest
	amount := CompoundInterest(principal, rate, years)

	// Display the result
	fmt.Printf("Compound interest after %d years: %.2f\n", years, amount)
}
