package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const filename = "words.txt"

var words []string

func main() {
	welcomeMessage()

	// read file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// read words from file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// start new game
	rand.NewSource(time.Now().UnixNano())
	word, lettersGuessed, wrongGuesses := newGame(words)

	// game loop
	runGame(word, lettersGuessed, wrongGuesses)
}

func welcomeMessage() {
	fmt.Println("Welcome to hangman!")
	fmt.Println("You have 6 chances to guess the word.")
	fmt.Println("Good luck!")
}

func generateWord(words []string) string {
	return words[rand.Intn(len(words))]
}

func newGame(words []string) (word string, lettersGuessed []string, wrongGuesses int) {
	word = generateWord(words)
	lettersGuessed = make([]string, 0)
	wrongGuesses = 0

	// print hangman
	printHangman(wrongGuesses)
	for i := 0; i < len(word); i++ {
		fmt.Print("_ ")
	}

	return word, lettersGuessed, wrongGuesses
}

func printHangman(guesses int) {
	if guesses == 0 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else if guesses == 1 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else if guesses == 2 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println(" |   |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else if guesses == 3 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println("/|   |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else if guesses == 4 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println("/|\\  |")
		fmt.Println("     |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else if guesses == 5 {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println("/|\\  |")
		fmt.Println("/    |")
		fmt.Println("     |")
		fmt.Println("=====+")
	} else {
		fmt.Println(" +---+")
		fmt.Println(" |   |")
		fmt.Println(" O   |")
		fmt.Println("/|\\  |")
		fmt.Println("/ \\  |")
		fmt.Println("     |")
		fmt.Println("=====+")
	}
}

func isAlpha(s string) bool {
	if len(s) != 1 {
		return false
	}
	r := rune(s[0])
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func validateGuess(guess string, lettersGuessed []string, word string) bool {
	guess = strings.ToLower(guess)
	if guess == "quit" {
		return true
	} else if len(guess) != 1 {
		switch {
		case len(guess) == 0:
			fmt.Println("Please enter a letter/word.")
			return false
		case len(guess) > 1:
			if guess == word {
				return true
			} else {
				fmt.Printf("Sorry, %s is not the answer.\n", guess)
				return true
			}
		}
	} else if strings.Contains(strings.Join(lettersGuessed, ""), guess) {
		fmt.Println("You have already guessed this letter.")
		return false
	} else if !isAlpha(guess) {
		fmt.Println("Please enter a valid letter.")
		return false
	}
	return true
}

func validateResponse(response string) bool {
	response = strings.ToLower(response)
	if response != "y" && response != "n" {
		fmt.Println("Please enter 'y' or 'n'.")
		return false
	} else {
		return true
	}
}

func runGame(word string, lettersGuessed []string, wrongGuesses int) {
	for wrongGuesses < 6 {
		// get user input
		var guess string
		fmt.Println("\nEnter a letter, word, or type 'quit' to exit: ")
		fmt.Scanln(&guess)

		// validate user input
		if !validateGuess(guess, lettersGuessed, word) {
			continue
		}

		if strings.ToLower(guess) == "quit" {
			fmt.Println("Thanks for playing!")
			return
		}

		// check if letter is in word
		if strings.Contains(word, guess) {
			lettersGuessed = append(lettersGuessed, guess)
		} else {
			wrongGuesses++
		}

		// check if user has won
		if strings.Contains(strings.Join(lettersGuessed, ""), word) {
			fmt.Println("Congratulations! You found the word!")

			var response string
			fmt.Println("Would you like to play again? (y/n): ")
			fmt.Scan(&response)

			for !validateResponse(response) {
				fmt.Println("Answer not recognized. Please enter 'y' or 'n'.")
				fmt.Scan(&response)
			}

			if response == "y" {
				word, lettersGuessed, wrongGuesses := newGame(words)
				runGame(word, lettersGuessed, wrongGuesses)
			} else {
				fmt.Println("Thanks for playing!")
				return
			}
		}

		// print hangman
		printHangman(wrongGuesses)
		for i := 0; i < len(word); i++ {
			if strings.Contains(strings.Join(lettersGuessed, ""), string(word[i])) {
				fmt.Printf("%s ", string(word[i]))
			} else {
				fmt.Print("_ ")
			}
		}
	}
	fmt.Printf("Sorry, you've run out of guesses. The word was: %s\n", word)

	// ask user if they want to play again and validate
	var response string
	fmt.Println("Would you like to play again? (y/n): ")
	fmt.Scan(&response)

	if !validateResponse(response) {
		runGame(word, lettersGuessed, wrongGuesses)
	} else if response == "y" {
		word, lettersGuessed, wrongGuesses := newGame(words)
		runGame(word, lettersGuessed, wrongGuesses)
	} else {
		fmt.Println("Thanks for playing!")
		return
	}
}
