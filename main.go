package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sayhello() {
	fmt.Println("Hello World!")
}

func saytime() {
	t := time.Now()
	fmt.Println("The time is", t.Format("3:04PM"))
}

func saysqrt() {
	z := 4.0
	fmt.Println("The square root of 4 is", math.Sqrt(z))
}

func sayrandomnumber() {
	fmt.Println("A number from 1 to 100:", rand.Intn(100))
}

func add(x float64, y float64) float64 {
	return x + y
}

func saysum() {
	num1 := 365.0
	num2 := 213.0
	fmt.Println("La suma de", num1, "y", num2, "da", add(num1, num2))
}

func main() {
	sayhello()
	saytime()
	saysqrt()
	sayrandomnumber()
	saysum()
}
