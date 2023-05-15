package main

import (
	"flag"
	"fmt"
)

func main() {
	balance := flag.Float64("b", 0, "The owed balance of the HELOC.")
	interestRate := flag.Float64("i", 0, "The HELOC interest rate.")
	daysInMonth := flag.Int("d", 30, "Number of days in the month.")
	flag.Parse()

	fmt.Printf("Balance: $%.2f\n", *balance)
	fmt.Printf("Interest Rate: %.2f%%\n", *interestRate)
	fmt.Printf("Days In The Month: %d\n", *daysInMonth)

	payment := calculateMonthlyPayment(*balance, *interestRate, *daysInMonth)

	fmt.Printf("\nYour monthly payment will be $%.2f", payment)
}

func calculateMonthlyPayment(balance float64, interestRate float64, daysInMonth int) (monthlyPayment float64) {
	dailyPayment := (balance * (interestRate / 100)) / 365
	monthlyPayment = dailyPayment * float64(daysInMonth)

	return monthlyPayment
}
