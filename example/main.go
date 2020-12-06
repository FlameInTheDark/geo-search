package main

import (
	"log"
	"time"

	gs "github.com/FlameInTheDark/geo-search"
)

func main() {
	client := gs.New("demo", "en", 10, time.Second * 5)

	places, err := client.Search("London")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(places)
}