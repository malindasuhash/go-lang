package main

import (
	"fmt"
)

var metersPerLiter float64 // Package level variable

func main() {
	metersPerLiter = 10.0
	paint := paintNeeded(4.2, 3.0)
	fmt.Printf("%.2f liters needed", paint)
	fmt.Println()
	paint += paintNeeded(5, 8.0)
	fmt.Printf("%.2f liters needed", paint)
	fmt.Println()
}

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	return area // Returning from a function of type float64
}
