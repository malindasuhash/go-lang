package main

import (
	"errors"
	"fmt"
	"log"
)

var metersPerLiter float64 // Package level variable

func main() {
	metersPerLiter = 10.0
	paint, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
		return // This is not needed because Fatal will break current flow and stop
	}
	fmt.Printf("%.2f liters needed", paint)
}

func paintNeeded(width float64, height float64) (float64, error) {
	var err error = nil // Initialising, not that it matters in this case
	if width < 0 {
		// This is one way of creating formatted errors (via errors package)
		err = errors.New(fmt.Sprintf("Width cannot be negative. It is %5.2f", width))
	}

	if height < 0 {
		// This is another way of creating errors
		err = fmt.Errorf("Height cannot be a negative number. It is %5.2f", height)
	}

	if err != nil {
		return 0, err
	}

	area := width * height
	return area, err // Returning from a function of type float64 and an error

}
