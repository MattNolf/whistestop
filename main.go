package main

import (
	"fmt"
	"log"

	"whistlestop.com/attire"
	"whistlestop.com/transport"
	"whistlestop.com/weather"
)

func main() {
	fmt.Println("Hello, world")

	weatherClient, err := weather.New("")
	if err != nil {
		log.Fatal(err)
	}

	attire, err := attire.New(weatherClient)
	if err != nil {
		log.Fatal(err)
	}

	transport.Serve(*attire)
}
