package main

import "fmt"

func main() {
	fmt.Println("Please enter the month of year: ")
	var month, season string
	_, _ = fmt.Scan(&month)
	switch month {
	case "december", "january", "february":
		season = "winter"
		fmt.Println("Season is : ", season)
	case "march", "april", "may":
		season = "spring"
		fmt.Println("Season is : ", season)
	case "june", "july", "august":
		season = "summer"
		fmt.Println("Season is : ", season)
	case "september", "october", "november":
		season = "autumn"
		fmt.Println("Season is : ", season)
	default :
		fmt.Println("Sorry are you from Mars???!!!")
	}
}
