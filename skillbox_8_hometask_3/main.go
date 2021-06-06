package main

import "fmt"

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, pay := range bills {
		switch pay {
		case 5:
			fives++
		case 10:
			switch {
			case fives > 0:
				fives--
				tens--
			default:
				return false
			}
		case 20:
			switch {
			case tens > 0 && fives > 0:
				tens--
				fives--
			case fives >= 3:
				fives -= 3
			default:
				return false
			}
		}
	}
	return true
}

func main() {
	var bills [] int
	bills = []int {5, 5, 5, 10, 20}
	fmt.Println("Given: ", bills)
	fmt.Println("Change: ", lemonadeChange(bills))

	bills = []int {10, 10}
	fmt.Println("Given: ", bills)
	fmt.Println("Change: ", lemonadeChange(bills))

	bills = []int {5, 5, 10, 10, 20}
	fmt.Println("Given: ", bills)
	fmt.Println("Change: ", lemonadeChange(bills))

	bills = []int {5, 5, 5, 10, 5, 10, 20}
	fmt.Println("Given: ", bills)
	fmt.Println("Change: ", lemonadeChange(bills))
}
