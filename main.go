package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func sayhello() {
	fmt.Println("Hello World!") // Saludo inicial
}

func saytime() {
	t := time.Now()
	fmt.Println("The time is", t.Format("3:04PM")) // Hora actual
}

func saysqrt() {
	z := 4.0
	fmt.Println("The square root of 4 is", math.Sqrt(z)) // Raiz cuadrada
}

func sayrandomnumber() {
	fmt.Println("A number from 1 to 100:", rand.Intn(100)) // Aquí me falta saber cómo hacer que cada vez salga un número diferente
}

func add(x, y float64) float64 {
	return x + y // Hace la suma entre dos valores
}

func saysum() {
	num1, num2 := 365.0, 213.0
	fmt.Println("La suma de", num1, "y", num2, "da", add(num1, num2)) // Te dice los dos valores a sumar y los suma
}

func main() {
	sayhello()
	saytime()
	saysqrt()
	sayrandomnumber()
	saysum()
}
