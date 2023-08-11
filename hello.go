package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	now := time.Now()
	fmt.Printf("Hello World from GoLang \nThis code was created in: \n\t05/08/2023 \n\nToday is: %s\n\n", now)
	fmt.Println("\"" + quote.Go() + "\"")
	fmt.Println("\"" + quote.Glass() + "\"")
	fmt.Println("\"" + quote.Hello() + "\"")
	fmt.Println("\"" + quote.Opt() + "\"")
}
