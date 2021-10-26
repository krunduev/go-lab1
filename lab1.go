package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Enter a: ")

	var a float64

	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Error in reading a.")
		return
	}

	z1 := math.Sin(4*a) * math.Cos(2*a) / (1+math.Cos(4*a)) / (1+math.Cos(2*a))
	fmt.Println("z1 = ", z1)

	b := 3 * math.Pi / 2 - a
	z2 := math.Cos(b) / math.Sin(b)
	fmt.Println("z2 = ", z2)
	fmt.Println("Press enter to continue...")
	fmt.Scanln()
}
