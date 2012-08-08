package main

import (
	"fmt"
	"github.com/injekt/github.go"
	"log"
)

func main() {
	user, err := github.GetUser("injekt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)

	repo, err := github.GetRepo("injekt/slop")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(repo)
}
