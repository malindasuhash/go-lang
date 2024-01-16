package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Printf("Value is %0.2f", 10/3.0) // Formatted %0.2f (2 decimal places)
	//Composition is Formatting verb which is %0.2f
	//and value width which is 0.2 portion

	formattedResult := fmt.Sprintf("Secondary value is %1.3f", 20/3.0)
	fmt.Printf(formattedResult)
	log.Print(reflect.TypeOf(formattedResult))

	fmt.Printf("This is a %d numeric value and my name is %s", 10, "malinda")
	// Onces in %d are called verbs and replacement figures are called values

	fmt.Println()
	fmt.Printf("This is a boolean %t", true)
	fmt.Println()

	fmt.Printf("I need a tab space %v here", "\t") // Literally the value
	fmt.Println()

	fmt.Printf("This is good %#v %#v", "\t", "\n") // Shows the actual value; note #
	fmt.Println()

	// Note the formatting and spacing below, very useful in console apps
	fmt.Printf("%10s # %5s \n", "Name", "Age")
	fmt.Printf("%10s # %5s \n", "--------", "----")
	fmt.Printf("%10s # %5s", "Malinda", "30")
	fmt.Println()
	fmt.Printf("%10s # %5s", "Elvis", "89")
	fmt.Println()
	fmt.Printf("%10s # %5s", "Martin", "9")
	fmt.Println()
	fmt.Printf("%10s # %5.2f", "Martin", 9.784)
}
