package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101)
	fmt.Println("Enter the number between 1 and 100")
	outerLoop:
	for {
		var answer int
		innerLoop:
		for {
			_, _ = fmt.Scan(&answer)
			switch {
			case answer < 1 || answer > 100:
				fmt.Println("The number should be between 1 and 100")
			default:
				break innerLoop
			}
		}
		switch {
		case answer == n:
			fmt.Println("Woow!!! You've found the number")
			break outerLoop
		case answer < n:
			fmt.Println("Searched number is higher")
		default:
			fmt.Println("Searched number is lower")
		}
	}
}
