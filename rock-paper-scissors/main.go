package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// declare variables
	var playerChoice string
	var computerChoice string
	var response string = "y"

	// loop until user wants to quit
	for response == "y" {
		// ask user for input
		fmt.Print("Enter your choice (rock, paper, scissors): ")
		fmt.Scan(&playerChoice)

		// validate
		playerChoice = validateChoice(playerChoice)

		// generate computer choice
		computerChoice = generateComputerChoice()

		var result string = battle(playerChoice, computerChoice)
		fmt.Printf("You chose %s, your opponent chose %s. %s\n", playerChoice, computerChoice, result)

		// ask user if they want to play again
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&response)

		// validate
		response = validateResponse(response)
	}
}

func generateComputerChoice() string {
	// generate random number
	choices := []string{"rock", "paper", "scissors"}

	// generate random choice
	rand.NewSource(time.Now().UnixNano())
	index := rand.Intn(3)
	choice := choices[index]

	return choice
}

func battle(choice1 string, choice2 string) string {
	// determine winner
	var result string

	if choice1 == choice2 {
		result = "It's a tie!"
	} else if choice1 == "rock" && choice2 == "scissors" ||
		choice1 == "paper" && choice2 == "rock" ||
		choice1 == "scissors" && choice2 == "paper" {
		result = "You win!"

	} else {
		result = "Your opponent wins!"
	}

	return result
}

func validateResponse(response string) string {
	// validate response
	response = strings.ToLower(response)

	for response != "y" && response != "n" {
		fmt.Println("Invalid response, try again.")
		fmt.Print("Do you want to play again? (y/n): ")
		fmt.Scan(&response)
	}

	return response
}

func validateChoice(choice string) string {
	// validate choice
	choice = strings.ToLower(choice)

	for choice != "rock" && choice != "paper" && choice != "scissors" {
		fmt.Println("Invalid choice, try again.")
		fmt.Print("Enter your choice (rock, paper, scissors): ")
		fmt.Scan(&choice)
	}

	return choice
}
