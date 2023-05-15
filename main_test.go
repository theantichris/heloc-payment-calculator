package main

import "testing"

func Test_CalculatePayment(t *testing.T) {
	balance := 80000.00
	interestRate := 10.00
	daysInMonth := 30

	sum1 := balance * (interestRate / 100)
	sum2 := sum1 / 365
	expected := sum2 * float64(daysInMonth)

	actual := calculateMonthlyPayment(balance, interestRate, daysInMonth)

	if expected != actual {
		t.Errorf("Expected %.2f got %.2f", expected, actual)
	}
}
