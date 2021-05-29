package main

import "fmt"

func main() {
	fmt.Println("Please enter the sum week day!!:  ")
	var weekDay, day string
	_, _ = fmt.Scan(&weekDay)
	switch  {
	case weekDay=="Mo" :
		day="Monday"
		fmt.Println("The day of the week day: ", day)
		fallthrough
	case weekDay=="Tu" :
		day="Tuesday"
		fmt.Println("The day of the week day: ", day)
		fallthrough
	case weekDay=="We" :
		day="Wednesday"
		fmt.Println("The day of the week day: ", day)
		fallthrough
	case weekDay=="Th" :
		day="Thursday"
		fmt.Println("The day of the week day: ", day)
		fallthrough
	case weekDay=="Fr" :
		day="Friday"
		fmt.Println("The day of the week day: ", day)
	default :
		fmt.Println("Leisure day or wrong work Day")
	}
}
