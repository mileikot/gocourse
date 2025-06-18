package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func guessingGame() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	//Generate rundom number between 1 and 100
	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing game")
	fmt.Println("I have choosen a number between 1 and 100")
	fmt.Println("Can you guess the number?")

	var guess int
	for {
		fmt.Println("Enter your guess")
		fmt.Scanln(&guess)

		// Check if the guess is correct

		if guess == target {
			fmt.Println("Congratulations! You guess the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}
