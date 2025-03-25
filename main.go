package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Guess - игра, в которой пользователю необходимо угадать число , заданное компьютером
	randomNumber := rand.Intn(100) + 1
	fmt.Println("Game - guess the number")
	fmt.Println("Enter your number between 1 to 100")
	var userEnter int
	for i := 9; i >= 0; i-- {
		scan, err := fmt.Scan(&userEnter)
		if err != nil {
			return
		}
		if scan < randomNumber {
			fmt.Println("Your guess is too low")
			fmt.Println(scan)
		} else if scan > randomNumber {
			fmt.Println("Your guess is too high")
		} else if scan == randomNumber {
			fmt.Println("Your guess is guess!!")
			break
		}
		if i == 0 {
			fmt.Println("Sorry. You didn't guess my number. It was", randomNumber)
			break
		}
		fmt.Println("You have only", i, "tries to guess")
	}
}
