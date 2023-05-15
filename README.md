# heloc-payment-calculator

A CLI to calculate the monthly payment of a HELOC.

## Usage

`heloc -b 80000 -i 10 -d 30`

``` text
Usage of ./heloc:
  -b float
    The owed balance of the HELOC.
  -d int
    Number of days in the month. (default 30)
  -i float
    The HELOC interest rate.
```

## Formula

``` text
BALANCE * INTEREST RATE = SUM1
SUM1 / 365 = SUM2
SUM2 * DAYS OF THE MONTH = PAYMENT
```
