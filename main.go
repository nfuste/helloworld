package main

import "fmt"

//"math"
//"math/rand"
//"net/http"
//"time"

/*func sayhello() {
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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Created by Nicky Knights")
}*/

type campaign struct {
	permalink    uint32 // min 0 - max 2147483647
	impressions  uint32
	requests     uint32
	campaignName string
}

func fillRate(x, y uint32) uint32 {
	return x / y * 100
}

func fillRate2(c *campaign) uint32 {
	fmt.Println(float64(c.impressions) / float64(c.requests))
	return c.impressions / c.requests * 100
}

func (c *campaign) fillRate3() uint32 {
	return c.impressions / c.requests * 100
}

func main() {
	/*sayhello()
	saytime()
	saysqrt()
	sayrandomnumber()
	saysum()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe(":8000", nil)*/
	aCampaign := campaign{
		permalink:    12345678,
		impressions:  1234212,
		requests:     2468424,
		campaignName: "Test Campaign"}

	fmt.Println(fillRate(aCampaign.impressions, aCampaign.requests))

	fmt.Println(fillRate2(&aCampaign)) // función - No se relaciona con la estructura, la trata como un parámetro

	fmt.Println(aCampaign.fillRate3()) // método - Se relaciona con la estructura

	fmt.Println("La campaña", aCampaign.campaignName, "tiene un fill rate del", aCampaign.fillRate3())
}
