package main

import "fmt"

func main() {
fmt.Println("Please enter the sum of monthly income: ")
var income, tax float64
_, _ = fmt.Scan(&income)
switch {
case income >= 100000:
	tax = income * 30/100
	fmt.Println("Taxes summ is 30% of net income, ", tax)
case income >= 60000:
	tax = income * 20/100
	fmt.Println("Taxes summ is 20% of net income, ", tax)
case income >= 10000:
	tax = income * 10/100
	fmt.Println("Taxes summ is 10% of net income, ", tax)
case income >= 0:
	tax = 0
	fmt.Println("Really>????? Your tax is: ", tax)
default :
	tax = 0
	fmt.Println("Sorry for that bro!!!", tax)
}
}
