// Classic guessing game

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeSeed := time.Now().Unix()
	rand.Seed(timeSeed) // Essentially seeding the randon number generator, deprecated since GO 1.20

	var target int = rand.Intn(100) + 1
	fmt.Println(target)
	// BTW - Go type int32 is not an alias to int; little gotcha!

	fmt.Println("I've chosen a number between 1 and 100")
	fmt.Println("Can you guess it?")

	var guessedCorrectly bool = false

	for guesses := 10; guesses > 0; guesses-- {
		fmt.Println("You have ", guesses, "left")

		reader := bufio.NewReader(os.Stdin)

		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("There was a problem reading the value", err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(userInput))
		if err != nil {
			log.Fatalln("Parsing failed", err)
		}

		if guess == target {
			fmt.Println("Good job! you guessed ", guess, "correctly!")
			guessedCorrectly = true
			break
		}

		if guess < target {
			fmt.Println("OPPS your guess is LESS than my guess")
		} else if guess > target {
			fmt.Println("OPPS your guess is MORE than my guess")
		}

	}

	if !guessedCorrectly {
		fmt.Println("Sorry, you have not guessed", target, "correctly")
	}

}
