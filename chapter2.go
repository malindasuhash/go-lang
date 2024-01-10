package main

import ( // These are library imports
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

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
	readValue, _ := bufio.NewReader(os.Stdin).ReadString('\n') // Blank identifier

	fmt.Println("Value read", readValue)
}
