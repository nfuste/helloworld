package main

import (
	"fmt"
	"math"
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
	fmt.Println("The square root of 4 is", math.Sqrt(4))
}

func main() {
	sayhello()
	saytime()
	saysqrt()

}
