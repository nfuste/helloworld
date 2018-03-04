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

func main() {
	sayhello()
	saytime()
	saysqrt()
	sayrandomnumber()
}
