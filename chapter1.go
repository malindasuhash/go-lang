package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("hello", "there")
	fmt.Println(math.Floor(9.34))
	fmt.Println(strings.Title("here is the title"))
	fmt.Println(reflect.TypeOf(10))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("10"))
	fmt.Println(reflect.TypeOf('A'))
	fmt.Println(reflect.TypeOf(10.34))

	var number int
	number = 10
	fmt.Println(number)

	var someString1, someString2 string = "malinda", "suhash1"

	fmt.Println(someString1, someString2)

	var number2 int
	fmt.Println((number2))

	var number3 = 9.33
	fmt.Println(number3)
	fmt.Println(reflect.TypeOf(number3))

	var booleanFalse bool
	fmt.Println(booleanFalse)

	tenApples := 10
	fourApples := 4
	thisIsAShortVariableDeclaration := "value to display"
	var apples = "apples."

	fmt.Println("I started with", tenApples, apples)
	fmt.Println("Some jerk ate", fourApples, apples)
	fmt.Println("There are", tenApples-fourApples, "left")

	fmt.Println((thisIsAShortVariableDeclaration))

	var floatVariable float64 = 4.53
	var intVariable int = 5

	fmt.Println(floatVariable * float64(intVariable))
	fmt.Println((floatVariable))

	var price int = 100
	fmt.Println("Price is", price, "dollars")

	var taxRate float64 = 0.08
	var tax = float64(price) * taxRate

	var total = tax + float64(price)
	fmt.Println("Total cost is", total, "dollars")

	availableFunds := 120
	fmt.Println("Available funds", availableFunds, "dollars")
	fmt.Println("Within budget", float64(availableFunds) >= total)

	var r rune = 58
	fmt.Println(r)
}
