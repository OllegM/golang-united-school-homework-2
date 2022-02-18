package main

import (
	"fmt"
	"gounited/lecture2/square"
)

func main() {
	fmt.Printf("Circle square = %.2f\n", square.CalcSquare(10, square.SidesCircle))
	fmt.Printf("Square square = %.2f\n", square.CalcSquare(10, square.SidesSquare))
	fmt.Printf("Triangle square = %.2f\n", square.CalcSquare(10, square.SidesTriangle))
}
