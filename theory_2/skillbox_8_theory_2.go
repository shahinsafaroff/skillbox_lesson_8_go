package main

import "fmt"

func main() {
	fmt.Println("Please enter the sum of monthly income: ")
	var income, tax float64
	_, _ = fmt.Scan(&income)
	switch {
	case income >= 100000:
		tax = (income - 100000) * 30/100
		fmt.Println("Taxes summ is 30% of net income for below 100K, ", tax)
		income = 100000
		fallthrough
	case income >= 60000:
		tax = (income - 60000) * 20/100
		fmt.Println("Taxes summ is 20% of net income, between 100K and 60K", tax)
		income = 60000
		fallthrough
	case income >= 10000:
		tax = (income - 10000) * 10/100
		fmt.Println("Taxes summ is 10% of net income, between 60k and 10k", tax)
		income = 10000
		fallthrough
	case income >= 0:
		tax = 0
		fmt.Println("Really>????? Your tax is: ", tax)
	default :
		tax = 0
		fmt.Println("Sorry for that bro!!!", tax)
	}
}
