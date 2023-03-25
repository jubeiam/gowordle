package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

func readLine() string {
	var w1 string
	_, err := fmt.Scanln(&w1)
	if err != nil {
		log.Fatal(err)
	}

	return w1
}

func generateRandomWorld() string {
	rand.Seed(time.Now().Unix())

	words := []string{
		"thumb",
		"strip",
		"cause",
		"fruit",
		"sweep",
		"charm",
		"funny",
		"decay",
		"agile",
		"flood",
		"venus",
		"clerk",
		"bible",
		"outer",
		"evoke",
		"tread",
		"quota",
		"medal",
		"anger",
		"snarl",
		"pause",
		"heavy",
		"guilt",
		"learn",
	}

	ret := words[rand.Intn(len(words))]

	return ret
}

type GuessEnum int

const (
	Invalid GuessEnum = iota
	Contains
	Correct
)

func processGuess(substr string, s string) []GuessEnum {
	ret := []GuessEnum{}
	sMap := strings.Split(s, "")
	if len(substr) != 5 {
		fmt.Printf("Invalid input require 5 characters got %d\n", len(substr))
		for i := 0; i < 5; i++ {
			ret = append(ret, Invalid)
		}

		return ret
	}

	substrMap := strings.Split(substr, "")

	for i, char := range substrMap {

		if char == sMap[i] {
			ret = append(ret, Correct)
			continue
		}

		if strings.Contains(s, char) {
			ret = append(ret, Contains)
			continue
		}

		ret = append(ret, Invalid)

	}

	return ret
}

func printGuess(guess string, ge []GuessEnum) {
	fmt.Println(guess)

	for _, x := range ge {
		switch x {
		case Contains:
			fmt.Print("~")

		case Invalid:
			fmt.Print("*")

		case Correct:
			fmt.Print("^")
		}
	}

	fmt.Println()
}

func main() {
	nTries := 5
	wordToGues := generateRandomWorld()

	// fmt.Printf("%s\n", wordToGues)

	for i := 0; i < nTries; i++ {
		fmt.Printf("Guess [%d/%d]: ", i, nTries)

		guess := readLine()

		printGuess(guess, processGuess(guess, wordToGues))

		if wordToGues == guess {
			fmt.Print("Congratulations You have guessed\n")
			return
		}

	}

	fmt.Print("Sorry, please try again\n")
}
