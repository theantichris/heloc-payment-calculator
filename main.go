package main

import (
	"flag"
	"fmt"
)

// BALANCE X INTEREST RATE = SUM 1
// SUM1 / 365 = SUM 2
// SUM 2 X DAYS OF THE MONTH = PAYMENT

func main() {
	balance := flag.Float64("b", 0, "The owed balance of the HELOC.")
	interestRate := flag.Float64("i", 0, "The HELOC interest rate.")
	daysInMonth := flag.Int("d", 30, "Number of days in the month.")
	flag.Parse()

	fmt.Printf("Days: %.2f\n", *balance)
	fmt.Printf("Interest Rate: %.2f\n", *interestRate)
	fmt.Printf("Days In The Month: %d\n", *daysInMonth)
}

func calculatePayment(balance float64, interestRate float64, daysInMonth int) (payment float64) {
	sum1 := balance * (interestRate / 100)
	sum2 := sum1 / 365

	payment = sum2 * float64(daysInMonth)

	return payment
}
