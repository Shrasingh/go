package main

import (
	"fmt"
	"math"

	"math/rand"
	"time"
)

func main() {
	fmt.Println("Maths in Go")
	fmt.Println("Square root of 16 is", math.Sqrt(16))
	fmt.Println("Power of 3^4 is", math.Pow(3, 4))
	fmt.Println("Value of Pi is", math.Pi)
	fmt.Println("Value of e is", math.E)
	fmt.Println("Sin value of 90 is", math.Sin(90))
	fmt.Println("Cos value of 0 is", math.Cos(0))
	fmt.Println("Log value of e is", math.Log(math.E))

	var one int = 2
	var two float64 = 3.5
	fmt.Println("Max of 2 and 3.5 is", math.Max(float64(one), two))

	// random number
	fmt.Println("Random number between 0 and 1 is", math.Round(0.6))
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	fmt.Println("Random number between 0 and 100 is", rand.Intn(100))
	fmt.Println("Random float number between 0.0 and 1.0 is", rand.Float64())
	fmt.Println("Random float number between 0.0 and 10.0 is", rand.Float64()*10)

	// Generate a random number in a specific range, e.g., between 50 and 100
	min := 50
	max := 100
	randomInRange := rand.Intn(max-min+1) + min
	fmt.Println("Random number between 50 and 100 is", randomInRange)

}
