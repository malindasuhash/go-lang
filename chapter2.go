package main

import ( // These are library imports
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// var packageScopeVariable = "bla" // Scope: This variable can be accessed by this package (package scoped)

func main() { // Entrypoint note the absence of args, which is retrived via "Flags"

	fmt.Println("Attempting to use the Time package")
	now := time.Now()

	fmt.Println("Time now is ", now.Year())

	toReplace := "G# to replace"

	var simpleString string = "test string"

	replaced := strings.NewReplacer("#", "O").Replace(toReplace)
	fmt.Println("New replaced string", replaced, simpleString)

	// Keyboard input
	fmt.Print("Enter score ?")
	readValue, err := bufio.NewReader(os.Stdin).ReadString('\n') // Blank identifier

	if err != nil { // Nil means NULL here.
		fmt.Println("Error has returned")
		log.Fatalln(err)
	}

	fmt.Println("Value read", readValue)
	input := strings.TrimSpace(readValue) // Removes all whitespace
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatalln("Cannot parse", err)
	}

	var status string
	if grade >= 60 {
		status = "Passing"
	} else {
		status = "Failing"
	}

	fmt.Println("Grade is ", status)

	// Checking OS stats
	var fileName string = "chapter1.go"
	fileCheckResult, err := os.Stat(fileName)

	if err != nil {
		log.Fatalln("There was a problem retriving results", err)
	}

	fmt.Println("File size is", fileCheckResult.Size(), "bytes")

	/*
		var int int = 10 // This essentially "Shadows the int type"
		fmt.Print(int) // Shadowing appears to be silent
		var int2 int = 20 // This breaks because int now refers to a variable
	*/
}
