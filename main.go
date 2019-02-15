package main

import (
	"fmt"
	"log"

	"github.com/matthiasbaldi/go-blog-api/communityindex"
)

func main() {
	item, err := communityindex.SearchByName("quarten")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("You selected the community:")
	fmt.Printf("%s: %s -> population: %s", item.Id, item.Name, item.Population.Total)
}
