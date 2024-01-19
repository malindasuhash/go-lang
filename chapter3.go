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
	fmt.Println()

	someValue := 10
	double(float64(someValue))
	fmt.Println(someValue)
	fmt.Println("Address of someValue", &someValue, ", type is", reflect.TypeOf(&someValue)) // type = *int

	var myInt int = 42
	var myIntPointer = &myInt
	*myIntPointer = 33 // New value
	fmt.Println(myInt)

	var simpleNmumber int = 10
	var returnedAddOne = addOne(simpleNmumber)
	fmt.Println("Returned value is", *returnedAddOne)
	// addOneAddress := &addTwo <== Why does this not work?
	acceptIntPointer(returnedAddOne)
}

func acceptIntPointer(intPointer *int) {
	fmt.Println("intPointer value", intPointer)
}

func addTwo() {

}

func addOne(number int) *int {
	number += 1
	return &number
}

func double(number float64) { // Pass-by-value
	number *= 2
	fmt.Println("Doubled value", number)

	var holdsPointerValue = &number
	fmt.Println("Pointer type is", reflect.TypeOf(holdsPointerValue))

	var secondaryPointerValue *float64 = holdsPointerValue // Assigning a pointer to another variable
	fmt.Println("Secondary pointer value", secondaryPointerValue)

	fmt.Println("Value at &secondaryPointerValue", *secondaryPointerValue) // Reading pointer value; *

	*secondaryPointerValue = 30 // Update to the variable that pointer points to
	fmt.Println("New value is", number)

	// number => &number (get the pointer) => *(&number) (returns value)
	fmt.Println("Value", *(&number))
}
