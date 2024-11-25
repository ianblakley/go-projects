package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// colors
const colorYellow = "\033[33m"
const colorReset = "\033[0m"

// word bank
var wordBank, err = readLines("words.txt")

func main() {
	// check for error reading file
	if err != nil {
		fmt.Println("Error reading file.")
		return
	}

	// seed random number generator
	rand.NewSource(time.Now().UnixNano())

	// welcome message & letters
	initGame()
	centerLetter := generateCenterLetter()
	otherLetters := generateOtherLetters(centerLetter)

	// run the game
	points, guessedWords := runGame(centerLetter, otherLetters)

	// print final points
	fmt.Printf("\nYou guessed %d words for a total of %d points!\n", len(guessedWords), points)
	var response string
	fmt.Println("Play again? (y/n): ")
	fmt.Scanln(&response)

	if strings.ToLower(response) == "y" {
		main()
	} else {
		fmt.Println("Thanks for playing!")
	}
}

func readLines(path string) ([]string, error) {
	// open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func initGame() {
	// print welcome message
	fmt.Println("Welcome to Spelling Bee!")
	var response string
	fmt.Println("Would you like to hear the rules? (y/n): ")
	fmt.Scanln(&response)

	if response == "y" {
		fmt.Println("\nHere are the rules:")
		fmt.Println(`1. The goal is to make as many words as possible using the center letter and any combination of the other letters.`)
		fmt.Println("2. Each word must be at least 4 letters long.")
		fmt.Println("3. Each word must include the center letter.")
		fmt.Println(`4. 1 point will be awarded for a four letter word, and the length of the word will be the amount of points awarded for longer words.`)
	}
	fmt.Println("Let's begin!")
}

func generateCenterLetter() rune {
	// generate random center letter
	centerLetter := rune(rand.Intn(26) + 'a')

	return centerLetter
}

func generateOtherLetters(centerLetter rune) []rune {
	// initialize list of other letters
	var otherLetters = []rune{}

	// generate random other letters without the center letter
	rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		letter := rune(rand.Intn(26) + 'a')
		if letter != centerLetter && !strings.ContainsRune(string(otherLetters), letter) {
			otherLetters = append(otherLetters, letter)
		}
	}
	return otherLetters
}

func isWordInLetters(guess string, centerLetter rune, otherLetters []rune) bool {
	// check if all letters in guess are in the list of letters
	letters := append(append(otherLetters[:3], centerLetter), otherLetters[3:]...)
	for _, letter := range guess {
		if !strings.ContainsRune(string(letters), letter) {
			return false
		}
	}
	return true
}

func printLetters(centerLetter rune, otherLetters []rune) {
	// print each letter with center letter in the middle colored yellow
	for _, letter := range otherLetters[:3] {
		fmt.Printf("%c  ", letter)
	}
	fmt.Printf("%s%c%s", colorYellow, centerLetter, colorReset)
	for _, letter := range otherLetters[3:] {
		fmt.Printf("  %c", letter)
	}
	fmt.Println()
}

func validateGuess(guess string, centerLetter rune, otherLetters []rune) bool {
	guess = strings.ToLower(guess)

	// check if user is quitting
	if guess == "q" {
		return false
	}

	if len(guess) < 4 {
		fmt.Println("Word must be at least 4 letters long.")
		return false
	} else if !strings.ContainsRune(guess, centerLetter) {
		fmt.Println("Word must include the center letter.")
		return false
	} else if !isWordInLetters(guess, centerLetter, otherLetters) {
		// check if all letters in the word are in the list of letters
		fmt.Println("Word must be made up of the given letters.")
		return false
	} else { // check if word in word bank
		for _, word := range wordBank {
			if word == guess {
				return true
			}
		}
		fmt.Println("Word not recognized.")
		return false
	}
}

func calculatePoints(guess string) int {
	if len(guess) == 4 {
		return 1
	} else {
		return len(guess)
	}
}

func runGame(centerLetter rune, otherLetters []rune) (int, []string) {
	points := 0
	guessedWords := []string{}

	play := true
	for play {
		// print letters, points, and valid words
		printLetters(centerLetter, otherLetters)
		fmt.Printf("Points: %d\n", points)
		fmt.Println("Enter 'q' to quit.")

		// get user input
		var guess string
		fmt.Println("Enter a word:")
		fmt.Scanln(&guess)

		// make sure guess is one word
		for strings.Contains(guess, " ") {
			fmt.Println("Guess must be one word.")
			var guess string
			fmt.Println("Enter a word:")
			fmt.Scanln(&guess)
		}

		// check if user wants to quit
		if strings.ToLower(guess) == "q" {
			play = false
		}

		// validate guess and add points
		if validateGuess(guess, centerLetter, otherLetters) {
			fmt.Printf("%s is worth %d points!", strings.ToTitle(guess), calculatePoints(guess))
			points += calculatePoints(guess)
			guessedWords = append(guessedWords, guess)
		}
	}
	return points, guessedWords
}
