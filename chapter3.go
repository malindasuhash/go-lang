package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Printf("Value is %0.2f", 10/3.0) // Formatted %0.2f (2 decimal places)

	formattedResult := fmt.Sprintf("Secondary value is %1.3f", 20/3.0)
	fmt.Printf(formattedResult)
	log.Print(reflect.TypeOf(formattedResult))
}
