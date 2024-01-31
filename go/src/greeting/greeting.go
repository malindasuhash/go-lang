package greeting

import "fmt"

func Hello() { // Capitalised functions are "exported"
	fmt.Println("Hello!")
}

func Hi() {
	fmt.Println("Hi!")
}
