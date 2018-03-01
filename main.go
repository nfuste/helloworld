package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Hello World", "The time is", t.Format("3:04PM"))

}
