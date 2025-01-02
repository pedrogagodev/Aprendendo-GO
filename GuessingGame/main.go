package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Let's start a guessing game! Choose a number for 1 to 100.")
	x := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {

		scanner.Scan()
		chute := scanner.Text()
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("Your choice has to be an entire number. Try again.")
			return
		}

		switch {
		case chuteInt > x:
			fmt.Println("Try a smaller number")
		case chuteInt < x:
			fmt.Println("Try a larger number")
		case chuteInt == x:
			fmt.Printf(
				"Congratulations! You got the number right!\n"+
					"You got it right in %d attempts.\n"+
					"Your attempts were: %v",
				i+1, chutes[:i],
			)
			return
		}

		chutes[i] = chuteInt
	}
	fmt.Printf(
		"Unfortunately you did not hit the number, the number was: %d\n"+
			"You had 10 attempts\n"+
			"Your attempts were: %v",
		x, chutes,
	)
}
